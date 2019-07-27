package itunescrawler

import "regexp"

var REGEXP_ITUNES_GENRE_PAGE = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/genre\/.*`)
var REGEXP_ITUNES_PODCAST_PAGE = regexp.MustCompile(`https?:\/\/podcasts.apple.com\/[a-z]+\/podcast\/.+\/id([0-9]+).*`)
var REGEXP_URL_WITH_FRAGEMENT = regexp.MustCompile(`(https?:\/\/.+)#.*`)

type ItunesLookupResp struct {
	Count   int                 `json:"resultCount"`
	Results []ItunesPodcastMeta `json:"results"`
}

type ItunesPodcastMeta struct {
	Kind    string `json:"kind"`
	Id      int    `json:"collectionId"`
	FeedUrl string `json:"feedUrl"`
}

// Check if given link points to itunes podcast page
// and return podcast id if true
func isPodcastPage(url string) (bool, string) {
	if REGEXP_ITUNES_PODCAST_PAGE.MatchString(url) {
		res := REGEXP_ITUNES_PODCAST_PAGE.FindStringSubmatch(url)
		return true, res[1]
	}
	return false, ""
}

// Check if given link points to itunes genre page
// and return normalised link if true
func isGenrePage(url string) (bool, string) {
	// Remove fragement identifier if present
	if REGEXP_URL_WITH_FRAGEMENT.MatchString(url) {
		res := REGEXP_URL_WITH_FRAGEMENT.FindStringSubmatch(url)
		url = res[1]
	}

	if REGEXP_ITUNES_GENRE_PAGE.MatchString(url) {
		return true, url
	}
	return false, ""
}
