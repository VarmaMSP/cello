package crawler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-collections/go-datastructures/set"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

// TODO: Implement Logging

const (
	ITUNES_SEED_URL              = "https://podcasts.apple.com/us/genre/podcasts/id26"
	PARALLEL_HTTP_REQUESTS_LIMIT = 5
)

// A primitive implementation of Mecator crawler to crawl Itunes
// http://www.cs.cornell.edu/courses/cs685/2002fa/mercator.pdf

type Crawler struct {
	// Url frontier
	urlQ chan string
	// Queues http responses
	respQ chan *http.Response
	// Queues batches of new podcast ids
	podcastIdsQ chan []string
	// Set of visited urls
	urlSet *set.Set
	// Set of podcast ids that are crawled
	podcastSet *set.Set
	// http client
	httpClient *http.Client
	// store
	store store.Store
}

func NewCrawler(store store.Store) *Crawler {
	c := &Crawler{
		urlQ:        make(chan string, 10000),
		respQ:       make(chan *http.Response),
		podcastIdsQ: make(chan []string, 100),
		urlSet:      set.New(),
		podcastSet:  set.New(),
		httpClient:  &http.Client{Timeout: time.Second * 20},
		store:       store,
	}

	go c.pollAndFetchPages()
	go c.pollAndProcessPages()
	go c.pollAndSaveNewPodcasts()

	return c
}

func (c *Crawler) Start() {
	// Clear visited url set and seed urlQ with root url
	c.urlSet.Clear()
	c.urlSet.Add(ITUNES_SEED_URL)
	c.urlQ <- ITUNES_SEED_URL
}

func (c *Crawler) pollAndFetchPages() {
	// A counting semaphore to keep at most n processes are running in parallel
	semaphore := make(chan int, PARALLEL_HTTP_REQUESTS_LIMIT)

	for {
		go func(url string) {
			semaphore <- 0
			defer func() { <-semaphore }()

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := c.httpClient.Do(req)
			if err != nil {
				fmt.Printf("REQ FAILED: GET %s - %s\n\n", url, err.Error())
			}
			c.respQ <- resp
		}(<-c.urlQ)
	}
}

func (c *Crawler) pollAndProcessPages() {
	for {
		go func(resp *http.Response) {
			if resp == nil {
				return
			}
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return
			}

			var newPodcastIds []string
			sel := doc.Find("a")
			for i := range sel.Nodes {
				if link, exist := sel.Eq(i).Attr("href"); exist {
					if ok, podcastId := isPodcastPage(link); ok {
						if !c.podcastSet.Exists(podcastId) {
							c.podcastSet.Add(podcastId)
							newPodcastIds = append(newPodcastIds, podcastId)
						}
						continue
					}

					if ok, nLink := isGenrePage(link); ok {
						if !c.urlSet.Exists(nLink) {
							c.urlSet.Add(nLink)
							c.urlQ <- nLink
						}
						continue
					}
				}
			}
			c.podcastIdsQ <- newPodcastIds
		}(<-c.respQ)
	}
}

func (c *Crawler) pollAndSaveNewPodcasts() {
	for {
		podcastIds := <-c.podcastIdsQ
		podcastItunesMeta := make([]*model.PodcastItunes, len(podcastIds))
		for i := 0; i < len(podcastIds); i++ {
			podcastItunesMeta[i] = &model.PodcastItunes{
				ItunesId:   podcastIds[i],
				ScrappedAt: time.Now().UTC().Format(model.MYSQL_DATETIME),
				AddedAt:    time.Now().UTC().Format(model.MYSQL_DATETIME),
			}
		}

		res := <-c.store.PodcastItunes().SaveAll(podcastItunesMeta)
		if res.Err != nil {
			fmt.Println(res.Err.Error())
		}
	}
}
