package crawler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-collections/go-datastructures/set"
)

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
	// Queues new podcast ids
	podcastIdQ chan string
	// Set of visited urls
	urlSet *set.Set
	// Set of podcast ids that are crawled
	podcastSet *set.Set
	// http client
	httpClient *http.Client
}

func NewCrawler() *Crawler {
	c := &Crawler{
		urlQ:       make(chan string, 10000),
		respQ:      make(chan *http.Response),
		podcastIdQ: make(chan string, 1000),
		urlSet:     set.New(ITUNES_SEED_URL),
		podcastSet: set.New(),
		httpClient: &http.Client{Timeout: time.Second * 50},
	}
	c.bootstrap()
	return c
}

func (c *Crawler) Start() {
	// Seed urlQ with root url to start crawling
	c.urlQ <- ITUNES_SEED_URL

	var cx chan int
	<-cx
}

func (c *Crawler) bootstrap() {
	go c.fetchDocuments()
	go c.processResponses()
}

func (c *Crawler) fetchDocuments() {
	// A counting semaphore to keep at most n processes running in parallel
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

func (c *Crawler) processResponses() {
	for {
		go func(resp *http.Response) {
			if resp == nil {
				return
			}
			fmt.Printf("PROCESSING %s\n", resp.Request.URL.String())
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return
			}

			sel := doc.Find("a")
			for i := range sel.Nodes {
				if link, exist := sel.Eq(i).Attr("href"); exist {
					if ok, podcastId := isPodcastPage(link); ok {
						if !c.podcastSet.Exists(podcastId) {
							c.podcastSet.Add(podcastId)
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
		}(<-c.respQ)
	}
}
