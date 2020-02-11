package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func GetSuggestions(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	suggestions, err := c.App.SuggestKeywords(c.Params.Query)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		SearchSuggestions: suggestions,
	}
}

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

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Podcasts:             podcasts,
			PodcastSearchResults: podcastSearchResults,
		}

	} else if c.Params.Type == "episode" {
		episodeSearchResults, err := c.App.SearchEpisodes(c.Params.Query, c.Params.SortBy, 0, 25)
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

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Podcasts:             podcasts,
			Episodes:             episodes,
			EpisodeSearchResults: episodeSearchResults,
		}

	} else {
		c.SetInvalidQueryParam("type")
	}
}

func BrowseResults(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	if c.Params.Type == "podcast" {
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

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Podcasts:             podcasts,
			PodcastSearchResults: podcastSearchResults,
		}

	} else if c.Params.Type == "episode" {
		episodeSearchResults, err := c.App.SearchEpisodes(c.Params.Query, c.Params.SortBy, c.Params.Offset, c.Params.Limit)
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

		if c.Session != nil && c.Session.UserId != 0 {
			playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Session.UserId, episodeIds)
			if err != nil {
				c.Err = err
				return
			}
			model.EpisodesJoinPlaybacks(episodes, playbacks)
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

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Podcasts:             podcasts,
			Episodes:             episodes,
			EpisodeSearchResults: episodeSearchResults,
		}

	} else {
		c.SetInvalidQueryParam("type")
	}
}

func SearchSuggestions(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	podcastSearchResults, err := c.App.TypeaheadPodcasts(c.Params.Query)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		PodcastSearchResults: podcastSearchResults,
	}
}
