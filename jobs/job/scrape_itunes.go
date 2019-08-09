package job

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

type ScrapeItunesJob struct {
	// input channel
	I chan interface{}
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

func NewScrapeItunesJob(store store.Store, producer *rabbitmq.Producer, workerLimit int) model.Job {
	return &ScrapeItunesJob{
		I:         make(chan interface{}),
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
}

func (job *ScrapeItunesJob) Start() *model.AppError {
	for off, lim := 0, 10000; ; off += lim {
		itunesIds, err := job.store.ItunesMeta().GetItunesIdList(off, lim)
		if err != nil {
			return err
		}
		if len(itunesIds) == 0 {
			break
		}
		for _, itunesId := range itunesIds {
			job.itunesIdF.Ignore(itunesId)
		}
	}

	go job.pollInput()
	go job.pollAndFetchPages()
	go job.pollAndProcessPages()
	go job.pollAndSaveItunesMeta()

	return nil
}

func (job *ScrapeItunesJob) Stop() *model.AppError {
	return nil
}

func (job *ScrapeItunesJob) InputChan() chan interface{} {
	return job.I
}

func (job *ScrapeItunesJob) pollInput() {
	for {
		<-job.I
		job.urlF.Clear()
		job.urlF.I <- ITUNES_SEED_URL
	}
}

func (job *ScrapeItunesJob) pollAndFetchPages() {
	semaphore := make(chan int, job.workerLimit)

	for {
		semaphore <- 0
		go func(url string) {
			defer func() { <-semaphore }()

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := job.httpClient.Do(req)
			if err != nil {
				fmt.Printf("GET %s: %s\n\n", url, err.Error())
				return
			}

			if resp.StatusCode == 200 {
				job.pageQ <- resp.Body
			}
		}(<-job.urlF.O)
	}
}

func (job *ScrapeItunesJob) pollAndProcessPages() {
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
					job.itunesIdF.I <- itunesId
					continue
				}
				if ok, link := isGenrePage(link); ok {
					job.urlF.I <- link
					continue
				}
			}
		}(<-job.pageQ)
	}
}

func (job *ScrapeItunesJob) pollAndSaveItunesMeta() {
	timeout := time.NewTimer(time.Minute)
	batchSize := 190

	for {
		var batch []string
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(30*time.Second); i < batchSize; i++ {
			select {
			case itunesId := <-job.itunesIdF.O:
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

		results, err := itunesLookup(batch, job.httpClient)
		if err != nil {
			continue
		}

		for _, result := range results {
			if result.Kind == "" || result.FeedUrl == "" {
				continue
			}

			if err := job.store.ItunesMeta().Save(&model.ItunesMeta{
				ItunesId:  strconv.Itoa(result.Id),
				FeedUrl:   result.FeedUrl,
				AddedToDb: model.StatusPending,
			}); err != nil {
				continue
			}

			job.producer.D <- map[string]string{
				"source":   "itunes",
				"id":       strconv.Itoa(result.Id),
				"feed_url": result.FeedUrl,
			}
		}
	}
}
