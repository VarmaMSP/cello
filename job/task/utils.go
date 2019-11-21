package task

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-collections/go-datastructures/set"
	"github.com/varmamsp/cello/model"
)

var (
	itunesLookupUrl = "https://itunes.apple.com/lookup?id="

	regexpItunesGenrePageUrl      = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/genre\/.*`)
	regexpItunesPodcastPageUrl    = regexp.MustCompile(`https?:\/\/(?:\bpodcasts\b|\bitunes\b).apple.com\/[a-z]+\/podcast(?:\/.+)?\/id([0-9]+).*`)
	regexpChartablePodcastPageUrl = regexp.MustCompile(`https?:\/\/chartable.com\/podcasts\/(.+)`)
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
func isPodcastPage(url string) (bool, string) {
	url = model.RemoveQueryFromUrl(url)
	if regexpItunesPodcastPageUrl.MatchString(url) {
		res := regexpItunesPodcastPageUrl.FindStringSubmatch(url)
		return true, res[1]
	}
	return false, ""
}

// Check if given link points to itunes genre page
// and return url with fragment removed
func isGenrePage(url string) (bool, string) {
	url = model.RemoveFragmentFromUrl(url)
	if regexpItunesGenrePageUrl.MatchString(url) {
		return true, url
	}
	return false, ""
}

// Check if given link points to chartable podcast page
// and return podcast id if true
func isChartablePodcastPage(url string) (bool, string) {
	url = model.RemoveQueryFromUrl(url)
	if regexpChartablePodcastPageUrl.MatchString(url) {
		res := regexpChartablePodcastPageUrl.FindStringSubmatch(url)
		return true, res[1]
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
