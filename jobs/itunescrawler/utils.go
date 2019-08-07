package itunescrawler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/varmamsp/cello/model"
)

var (
	itunesLookupUrl            = "https://itunes.apple.com/lookup?id="
	regexpItunesGenrePageUrl   = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/genre\/.*`)
	regexpItunesPodcastPageUrl = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/podcast\/.+\/id([0-9]+).*`)
)

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

// Check if given link points to itunes podcast page
// and return podcast id if true
func isPodcastPage(url string) (bool, string) {
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
