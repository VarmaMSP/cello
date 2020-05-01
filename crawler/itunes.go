package crawler

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
)

const (
	ITUNES_SEED_URL   = "https://podcasts.apple.com/us/genre/podcasts/id26"
	ITUNES_LOOKUP_URL = "https://itunes.apple.com/lookup?id="
)

// A primitive implementation of Mecator crawler to crawl Itunes
// http://www.cs.cornell.edu/courses/cs685/2002fa/mercator.pdf
//
// 1 - Crawl itunes website and find itunes ids for podcasts.
// 2 - Use Itunes lookup API to fetch feed urls for above ids.
// 3 - Push this data into rabbitmq queue to process later.
type ItunesCrawler struct {
	// store
	store store.Store
	// logger
	log zerolog.Logger
	// message queue
	importPodcastP messagequeue.Producer
	// url frontier
	urlF *Frontier
	// itunes Id frontier
	itunesIdF *Frontier
	// chan to hold pages till processed
	pageQ chan io.ReadCloser
	// http client
	httpClient *http.Client
	// rate limiter
	rateLimiter chan struct{}
}

func NewItunesCrawler(s store.Store, b messagequeue.Broker, log zerolog.Logger, config *model.Config) (*ItunesCrawler, error) {
	importPodcastP, err := b.NewProducer(
		messagequeue.EXCHANGE_PHENOPOD_DIRECT,
		messagequeue.ROUTING_KEY_IMPORT_PODCAST,
		config.Queues.ImportPodcast.DeliveryMode,
	)
	if err != nil {
		return nil, err
	}

	crawler := &ItunesCrawler{
		store:          s,
		log:            log.With().Str("ctx", "crawler").Logger(),
		urlF:           NewFrontier(10000),
		itunesIdF:      NewFrontier(10000),
		pageQ:          make(chan io.ReadCloser, 10),
		importPodcastP: importPodcastP,
		httpClient: &http.Client{
			Timeout: 40 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        20,
				MaxIdleConnsPerHost: 20,
			},
		},
		rateLimiter: make(chan struct{}, 10),
	}

	for off, lim := 0, 10000; ; off += lim {
		feeds, err := crawler.store.Feed().GetBySourcePaginated("ITUNES_SCRAPER", off, lim)
		if err != nil {
			return nil, err
		}
		for _, feed := range feeds {
			crawler.itunesIdF.Ignore(feed.SourceId)
		}
		if len(feeds) < lim {
			break
		}
	}

	go crawler.pollAndFetchPages()
	go crawler.pollAndProcessPages()
	go crawler.pollAndSaveFeedDetails()

	return crawler, nil
}

func (c *ItunesCrawler) Call() {
	c.urlF.Clear()
	c.urlF.I <- ITUNES_SEED_URL
}

func (c *ItunesCrawler) pollAndFetchPages() {
	for {
		c.rateLimiter <- struct{}{}
		go func(url string) {
			defer func() { <-c.rateLimiter }()

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := c.httpClient.Do(req)
			if err != nil {
				return
			}

			if resp.StatusCode == 200 {
				c.pageQ <- resp.Body
			}
		}(<-c.urlF.O)
	}
}

func (c *ItunesCrawler) pollAndProcessPages() {
	for {
		go func(page io.ReadCloser) {
			doc, err := goquery.NewDocumentFromReader(page)
			page.Close()
			if err != nil {
				return
			}

			sel := doc.Find("a")
			for i := range sel.Nodes {
				link, exist := sel.Eq(i).Attr("href")
				if !exist {
					continue
				}
				if ok, itunesId := isItunesPodcastPage(link); ok {
					c.itunesIdF.I <- itunesId
					continue
				}
				if ok, link := isItunesGenrePage(link); ok {
					c.urlF.I <- link
					continue
				}
			}
		}(<-c.pageQ)
	}
}

func (c *ItunesCrawler) pollAndSaveFeedDetails() {
	timeout := time.NewTimer(time.Minute)
	batchSize := 190

	for {
		var batch []string
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(30*time.Second); i < batchSize; i++ {
			select {
			case itunesId := <-c.itunesIdF.O:
				batch = append(batch, itunesId)
				if len(batch) == batchSize && !timeout.Stop() {
					<-timeout.C
				}
			case <-timeout.C:
				break BATCH_LOOP
			}
		}
		if len(batch) == 0 {
			continue
		}

		results, err := itunesLookup(batch, c.httpClient)
		if err != nil {
			continue
		}

		for _, result := range results {
			if result.Kind == "" || result.FeedUrl == "" {
				continue
			}

			feed := &model.Feed{
				Source:   "ITUNES_SCRAPER",
				SourceId: strconv.Itoa(result.Id),
				Url:      result.FeedUrl,
			}
			if err := c.store.Feed().Save(feed); err != nil {
				c.log.Error().Msg(err.Error())
				continue
			}

			if err := c.importPodcastP.Publish(feed); err != nil {
				c.log.Error().Msg(err.Error())
				continue
			}
		}
	}
}
