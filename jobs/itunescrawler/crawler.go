package itunescrawler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/varmamsp/cello/services/rabbitmq"

	"github.com/PuerkitoBio/goquery"
	"github.com/varmamsp/cello/model"
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
	// url frontier
	urlF *Frontier
	// itunes Id frontier
	itunesIdF *Frontier
	// chan to hold pages till processed
	pageQ chan io.ReadCloser
	// store
	store store.Store
	// rabbitmq producer
	producer *rabbitmq.Producer
	// http client
	httpClient *http.Client
	// worker limit
	workerLimit int
}

func New(store store.Store, producer *rabbitmq.Producer, workerLimit int) (*ItunesCrawler, error) {
	c := &ItunesCrawler{
		urlF:      NewFrontier(10000),
		itunesIdF: NewFrontier(10000),
		pageQ:     make(chan io.ReadCloser, workerLimit),
		store:     store,
		producer:  producer,
		httpClient: &http.Client{
			Timeout: 40 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        2 * workerLimit,
				MaxIdleConnsPerHost: 2 * workerLimit,
			},
		},
		workerLimit: workerLimit,
	}

	for off, lim := 0, 10000; ; off += lim {
		itunesIds, err := c.store.ItunesMeta().GetItunesIdList(off, lim)
		if err != nil {
			return nil, err
		}
		if len(itunesIds) == 0 {
			break
		}

		for _, itunesId := range itunesIds {
			c.itunesIdF.Ignore(itunesId)
		}
	}

	go c.pollAndFetchPages()
	go c.pollAndProcessPages()
	go c.pollAndSaveNewPodcasts()

	return c, nil
}

func (c *ItunesCrawler) Run() {
	c.urlF.Clear()
	c.urlF.I <- ITUNES_SEED_URL
}

func (c *ItunesCrawler) pollAndFetchPages() {
	semaphore := make(chan int, c.workerLimit)

	for {
		semaphore <- 0

		go func(url string) {
			defer func() { <-semaphore }()

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := c.httpClient.Do(req)
			if err != nil {
				fmt.Printf("GET %s: %s\n\n", url, err.Error())
				return
			}

			if resp.StatusCode == 200 {
				c.pageQ <- resp.Body
			}
		}(<-c.urlF.S)
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
				if ok, itunesId := isPodcastPage(link); ok {
					c.itunesIdF.I <- itunesId
					continue
				}
				if ok, link := isGenrePage(link); ok {
					c.urlF.I <- link
					continue
				}
			}
		}(<-c.pageQ)
	}
}

func (c *ItunesCrawler) pollAndSaveNewPodcasts() {
	timeout := time.NewTimer(time.Minute)
	batchSize := 190

	for {
		var batch []string
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(30*time.Second); i < batchSize; i++ {
			select {
			case itunesId := <-c.itunesIdF.S:
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

			if err := c.store.ItunesMeta().Save(&model.ItunesMeta{
				ItunesId:  strconv.Itoa(result.Id),
				FeedUrl:   result.FeedUrl,
				AddedToDb: model.StatusPending,
			}); err != nil {
				continue
			}

			c.producer.D <- map[string]string{
				"source":   "itunes",
				"id":       strconv.Itoa(result.Id),
				"feed_url": result.FeedUrl,
			}
		}
	}
}
