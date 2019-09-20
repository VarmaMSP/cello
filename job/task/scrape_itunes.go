package task

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
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

type ScrapeItunes struct {
	*app.App
	// url frontier
	urlF *Frontier
	// itunes Id frontier
	itunesIdF *Frontier
	// chan to hold pages till processed
	pageQ chan io.ReadCloser
	// rabbitmq producer
	importPodcastP *rabbitmq.Producer
	// http client
	httpClient *http.Client
	// rate limiter
	rateLimiter chan struct{}
}

func NewScrapeItunes(app *app.App, config *model.Config) (*ScrapeItunes, error) {
	importPodcastP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_IMPORT_PODCAST,
		DeliveryMode: config.Queues.ImportPodcast.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}

	workerLimit := 10

	scrapeItunes := &ScrapeItunes{
		App:            app,
		urlF:           NewFrontier(10000),
		itunesIdF:      NewFrontier(10000),
		pageQ:          make(chan io.ReadCloser, workerLimit),
		importPodcastP: importPodcastP,
		httpClient: &http.Client{
			Timeout: 40 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: workerLimit,
			},
		},
		rateLimiter: make(chan struct{}, workerLimit),
	}

	for off, lim := 0, 10000; ; off += lim {
		feeds, err := scrapeItunes.Store.Feed().GetAllBySource("ITUNES_SCRAPER", off, lim)
		if err != nil {
			return nil, err
		}
		for _, feed := range feeds {
			scrapeItunes.itunesIdF.Ignore(feed.SourceId)
		}
		if len(feeds) < lim {
			break
		}
	}

	go scrapeItunes.pollAndFetchPages()
	go scrapeItunes.pollAndProcessPages()
	go scrapeItunes.pollAndSaveFeedDetails()

	return scrapeItunes, nil
}

func (s *ScrapeItunes) Call() {
	s.Log.Info().Msg("Scrape Itunes started")
	s.urlF.Clear()
	s.urlF.I <- ITUNES_SEED_URL
}

func (s *ScrapeItunes) pollAndFetchPages() {
	for {
		s.rateLimiter <- struct{}{}
		go func(url string) {
			defer func() { <-s.rateLimiter }()

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := s.httpClient.Do(req)
			if err != nil {
				fmt.Printf("GET %s: %s\n\n", url, err.Error())
				return
			}

			if resp.StatusCode == 200 {
				s.pageQ <- resp.Body
			}
		}(<-s.urlF.O)
	}
}

func (s *ScrapeItunes) pollAndProcessPages() {
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
				if ok, itunesId := isPodcastPage(link); ok {
					s.itunesIdF.I <- itunesId
					continue
				}
				if ok, link := isGenrePage(link); ok {
					s.urlF.I <- link
					continue
				}
			}
		}(<-s.pageQ)
	}
}

func (s *ScrapeItunes) pollAndSaveFeedDetails() {
	timeout := time.NewTimer(time.Minute)
	batchSize := 190

	for {
		var batch []string
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(30*time.Second); i < batchSize; i++ {
			select {
			case itunesId := <-s.itunesIdF.O:
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

		results, err := itunesLookup(batch, s.httpClient)
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
			if err := s.Store.Feed().Save(feed); err != nil {
				continue
			}

			s.importPodcastP.D <- feed
		}
	}
}