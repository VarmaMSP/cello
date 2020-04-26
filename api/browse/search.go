package browse

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

func searchResults(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireQuery(); c.Err != nil {
		return
	}

	if c.Params.Type == "podcast" {
		results, err := c.App.SearchPodcasts(c.Params.Query, c.Params.Offset, c.Params.Limit)
		if err != nil {
			c.Err = err
			return
		}

		c.Response.StatusCode = http.StatusOK
		c.Response.Data.SearchResults.Podcasts = append(c.Response.Data.SearchResults.Podcasts, results...)

	} else if c.Params.Type == "episode" {
		episodeResults, err := c.App.SearchEpisodes(c.Params.Query, c.Params.SortBy, c.Params.Offset, c.Params.Limit)
		if err != nil {
			c.Err = err
			return
		}
		if err := c.App.LoadMediaUrls(episodeResults); err != nil {
			c.Err = err
			return
		}

		podcastIds := make([]int64, len(episodeResults))
		for i, episode := range episodeResults {
			podcastIds[i] = episode.PodcastId
		}
		podcasts, err := c.App.GetPodcastsByIds(podcastIds)
		if err != nil {
			c.Err = err
			return
		}

		if c.Session != nil && c.Session.UserId != 0 {
			if err := c.App.LoadPlaybacks(c.Session.UserId, episodeResults); err != nil {
				c.Err = err
				return
			}
		}

		c.Response.StatusCode = http.StatusOK
		c.Response.Data.Podcasts = append(c.Response.Data.Podcasts, podcasts...)
		c.Response.Data.SearchResults.Episodes = append(c.Response.Data.SearchResults.Episodes, episodeResults...)

	} else {
		c.SetInvalidQueryParam("type")
	}
}

func podcastSearch(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireQuery().RequirePodcastId(); c.Err != nil {
		return
	}

	episodeResults, err := c.App.SearchEpisodesInPodcast(c.Params.PodcastId, c.Params.Query, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}
	if err := c.App.LoadMediaUrls(episodeResults); err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		if err := c.App.LoadPlaybacks(c.Session.UserId, episodeResults); err != nil {
			c.Err = err
			return
		}
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data.SearchResults.Episodes = append(c.Response.Data.SearchResults.Episodes, episodeResults...)
}
