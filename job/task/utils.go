package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	h "github.com/go-http-utils/headers"
	"github.com/golang-collections/go-datastructures/set"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/gofeed/rss"
)

var (
	itunesLookupUrl = "https://itunes.apple.com/lookup?id="

	regexpItunesGenrePageUrl      = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/genre\/.*`)
	regexpItunesPodcastPageUrl    = regexp.MustCompile(`https?:\/\/(?:\bpodcasts\b|\bitunes\b).apple.com\/[a-z]+\/podcast(?:\/.+)?\/id([0-9]+).*`)
	regexpChartablePodcastPageUrl = regexp.MustCompile(`https?:\/\/chartable.com\/podcasts\/.+`)
)

type Frontier struct {
	// Input channel
	I chan string
	// Output channel
	O chan string
	// Set of strings processed till now
	set *set.Set
}

func NewFrontier(size int) *Frontier {
	frontier := &Frontier{
		I:   make(chan string, 1000),
		O:   make(chan string, size),
		set: set.New(),
	}

	go func() {
		for i := range frontier.I {
			if !frontier.set.Exists(i) {
				frontier.set.Add(i)
				frontier.O <- i
			}
		}
	}()

	return frontier
}

// Ignore input
func (f *Frontier) Ignore(s string) {
	f.set.Add(s)
}

// Clear all values received till now
func (f *Frontier) Clear() {
	f.set.Clear()
}

// Check if given link points to itunes podcast page
// and return podcast id if true
func isItunesPodcastPage(url string) (bool, string) {
	url = model.RemoveQueryFromUrl(url)
	if regexpItunesPodcastPageUrl.MatchString(url) {
		res := regexpItunesPodcastPageUrl.FindStringSubmatch(url)
		return true, res[1]
	}
	return false, ""
}

// Check if given link points to itunes genre page
// and return url with fragment removed
func isItunesGenrePage(url string) (bool, string) {
	url = model.RemoveFragmentFromUrl(url)
	if regexpItunesGenrePageUrl.MatchString(url) {
		return true, url
	}
	return false, ""
}

// Check if given link points to chartable podcast page
// and return url with query removed if true
func isChartablePodcastPage(url string) (bool, string) {
	url = model.RemoveQueryFromUrl(url)
	if regexpChartablePodcastPageUrl.MatchString(url) {
		return true, url
	}
	return false, ""
}

type ItunesLookupResp struct {
	Count   int                  `json:"resultCount"`
	Results []ItunesLookupResult `json:"results"`
}

type ItunesLookupResult struct {
	Kind    string `json:"kind"`
	Id      int    `json:"collectionId"`
	FeedUrl string `json:"feedUrl"`
}

// Fetch podcast details from itunes lookup API
func itunesLookup(podcastIds []string, httpClient *http.Client) ([]ItunesLookupResult, error) {
	url := itunesLookupUrl + strings.Join(podcastIds, ",")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	lookupResp := &ItunesLookupResp{}
	if err := json.NewDecoder(resp.Body).Decode(lookupResp); err != nil {
		return nil, err
	}

	return lookupResp.Results, nil
}

func fetchAndParseHtml(url string, retryIfThrottled bool) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 503 && retryIfThrottled {
		<-(time.NewTimer(2 * time.Minute)).C
		return fetchAndParseHtml(url, false)
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// TODO: REMOVE THIS
// Fetch RSS feed
func fetchRssFeed(feedUrl string, headers map[string]string, httpClient *http.Client) (*rss.Feed, map[string]string, *model.AppError) {
	appE := (&model.AppError{}).Id("fetch_rss_feed")

	// request
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return nil, nil, appE.
			DetailedError(err.Error()).
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			Retry()
	}
	if v, ok := headers[h.ETag]; ok {
		req.Header.Add(h.IfNoneMatch, v)
	}
	if v, ok := headers[h.LastModified]; ok {
		req.Header.Add(h.IfModifiedSince, v)
	}
	req.Header.Add(h.CacheControl, "no-cache")

	// make request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, appE.
			DetailedError(err.Error()).
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			Retry()
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, map[string]string{
			h.ETag:         resp.Header.Get(h.ETag),
			h.LastModified: resp.Header.Get(h.LastModified),
		}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, appE.
			Comment(model.COMMENT_INVALID_STATUS_CODE).
			DetailedError(fmt.Sprintf("invalid status code: %d", resp.StatusCode))
	}

	// Relax retrictions on content type
	// if c := resp.Header.Get(h.ContentType); !model.IsContentTypeFeed(c) {
	// 	return nil, nil, appE.
	// 		Comment(model.COMMENT_INVALID_CONTENT_TYPE).
	// 		DetailedError(fmt.Sprintf("invalid content type: %s", c))
	// }

	// parse feed
	parser := &rss.Parser{}
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		return nil, nil, appE.DetailedError(err.Error())
	}

	return feed, map[string]string{
		h.ETag:         resp.Header.Get(h.ETag),
		h.LastModified: resp.Header.Get(h.LastModified),
	}, nil
}
