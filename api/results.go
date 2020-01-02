package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func GetResultsPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	podcastSearchResults, err := c.App.SearchPodcasts(c.Params.Query, 0, 25)
	if err != nil {
		c.Err = err
		return
	}

	podcastIds := make([]int64, len(podcastSearchResults))
	for i, podcastSearchResult := range podcastSearchResults {
		podcastIds[i] = podcastSearchResult.Id
	}
	podcasts, err := c.App.GetPodcastsByIds(podcastIds)
	if err != nil {
		c.Err = err
		return
	}

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"podcasts":               podcasts,
		"podcast_search_results": podcastSearchResults,
	}))
}

func BrowseResults(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	podcastSearchResults, err := c.App.SearchPodcasts(c.Params.Query, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	podcastIds := make([]int64, len(podcastSearchResults))
	for i, podcastSearchResult := range podcastSearchResults {
		podcastIds[i] = podcastSearchResult.Id
	}
	podcasts, err := c.App.GetPodcastsByIds(podcastIds)
	if err != nil {
		c.Err = err
		return
	}

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"podcasts":               podcasts,
		"podcast_search_results": podcastSearchResults,
	}))
}
