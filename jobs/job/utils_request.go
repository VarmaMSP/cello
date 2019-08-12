package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	h "github.com/go-http-utils/headers"
	"github.com/mmcdole/gofeed/rss"
	"github.com/varmamsp/cello/model"
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

// Fetch RSS feed
func fetchRssFeed(feedUrl string, headers map[string]string, httpClient *http.Client) (*rss.Feed, map[string]string, *model.AppError) {
	appErrorC := model.NewAppErrorC(
		"jobs.job.utils.fetch_rss_feed",
		http.StatusInternalServerError,
		map[string]string{"feed_url": feedUrl},
	)

	// request
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return nil, nil, appErrorC(err.Error())
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
		return nil, nil, appErrorC(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, map[string]string{
			"ETag":          resp.Header.Get("ETag"),
			"Last-Modified": resp.Header.Get("Last-Modified"),
		}, nil
	}

	if resp.StatusCode == http.StatusOK {
		// parse xml
		parser := &rss.Parser{}
		feed, err := parser.Parse(resp.Body)
		if err != nil {
			return nil, nil, appErrorC(fmt.Sprintf("Cannot parse feed: %s", err.Error()))
		}

		return feed, map[string]string{
			"ETag":          resp.Header.Get("ETag"),
			"Last-Modified": resp.Header.Get("Last-Modified"),
		}, nil
	}

	return nil, nil, appErrorC("Invalid http status code.")
}
