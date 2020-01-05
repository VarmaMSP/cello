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

	if c.Params.Type == "podcast" {
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

	} else if c.Params.Type == "episode" {
		episodeSearchResults, err := c.App.SearchEpisodes(c.Params.Query, "", 0, 25)
		if err != nil {
			c.Err = err
			return
		}

		episodeIds := make([]int64, len(episodeSearchResults))
		for i, episodeSearchResult := range episodeSearchResults {
			episodeIds[i] = episodeSearchResult.Id
		}
		episodes, err := c.App.GetEpisodesByIds(episodeIds)
		if err != nil {
			c.Err = err
			return
		}

		podcastIds := make([]int64, len(episodes))
		for i, episode := range episodes {
			podcastIds[i] = episode.PodcastId
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
			"episodes":               episodes,
			"episode_search_results": episodeSearchResults,
		}))

	} else {
		c.SetInvalidQueryParam("type")
	}
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
