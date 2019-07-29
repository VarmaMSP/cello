package itunescrawler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-collections/go-datastructures/set"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

// TODO: Implement Logging

const (
	ITUNES_SEED_URL              = "https://podcasts.apple.com/us/genre/podcasts/id26"
	ITUNES_LOOKUP_URL            = "https://itunes.apple.com/lookup?id="
	PARALLEL_HTTP_REQUESTS_LIMIT = 5
)

// A primitive implementation of Mecator crawler to crawl Itunes
// http://www.cs.cornell.edu/courses/cs685/2002fa/mercator.pdf

type ItunesCrawler struct {
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

func New(store store.Store) *ItunesCrawler {
	c := &ItunesCrawler{
		urlQ:        make(chan string, 10000),
		respQ:       make(chan *http.Response),
		podcastIdsQ: make(chan []string, 100),
		urlSet:      set.New(),
		podcastSet:  set.New(),
		httpClient:  &http.Client{Timeout: time.Second * 20},
		store:       store,
	}

	c.loadPodcastSet()
	go c.pollAndFetchPages()
	go c.pollAndProcessPages()
	go c.pollAndSaveNewPodcasts()

	return c
}

func (c *ItunesCrawler) Run() {
	// Clear visited url set and seed urlQ with root url
	c.urlSet.Clear()
	c.urlSet.Add(ITUNES_SEED_URL)
	c.urlQ <- ITUNES_SEED_URL
}

func (c *ItunesCrawler) loadPodcastSet() {
	afterId := ""

	for {
		res := <-c.store.PodcastItunes().GetItunesIdsAfter(afterId, 50000)
		ids := res.Data.([]string)
		if len(ids) > 0 {
			for _, id := range ids {
				c.podcastSet.Add(id)
			}
			afterId = ids[len(ids)-1]
		} else {
			return
		}
	}
}

func (c *ItunesCrawler) pollAndFetchPages() {
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

func (c *ItunesCrawler) pollAndProcessPages() {
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
				link, exist := sel.Eq(i).Attr("href")
				if !exist {
					continue
				}

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

			// Queue batches of size 200
			i := 0
			l := len(newPodcastIds)
			for i < l {
				if i+200 < l {
					c.podcastIdsQ <- newPodcastIds[i : i+200]
				} else {
					c.podcastIdsQ <- newPodcastIds[i:l]
				}
				i = i + 200
			}
		}(<-c.respQ)
	}
}

func (c *ItunesCrawler) pollAndSaveNewPodcasts() {
	for {
		podcastIds := <-c.podcastIdsQ
		req, _ := http.NewRequest("GET", ITUNES_LOOKUP_URL+strings.Join(podcastIds, ","), nil)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			continue
		}
		lookupResp := &ItunesLookupResp{}
		err = json.NewDecoder(resp.Body).Decode(lookupResp)
		if err != nil {
			continue
		}

		podcastItunesMeta := make([]*model.PodcastItunes, len(lookupResp.Results))
		for i, result := range lookupResp.Results {
			podcastItunesMeta[i] = &model.PodcastItunes{
				ItunesId:   strconv.Itoa(result.Id),
				FeedUrl:    result.FeedUrl,
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
