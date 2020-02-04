package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func GetPodcastPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequirePodcastId()
	if c.Err != nil {
		return
	}

	feed, err := c.App.GetFeed(c.Params.PodcastId)
	if err != nil {
		c.Err = err
		return
	}
	w.Header().Set(headers.CacheControl, "private, max-age=300, must-revalidate")
	if feed.ETag != "" {
		ifNoneMatch := req.Header.Get(headers.IfNoneMatch)
		w.Header().Set(headers.ETag, feed.ETag)
		if ifNoneMatch != "" && ifNoneMatch == feed.ETag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	if feed.LastModified != "" {
		ifModifiedSince := req.Header.Get(headers.IfModifiedSince)
		w.Header().Set(headers.LastModified, feed.LastModified)
		if ifModifiedSince != "" && ifModifiedSince == feed.LastModified {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	podcast, err := c.App.GetPodcast(c.Params.PodcastId, true)
	if err != nil {
		c.Err = err
		return
	}

	episodes, err := c.App.GetEpisodesInPodcast(c.Params.PodcastId, "pub_date_desc", 0, 15)
	if err != nil {
		c.Err = err
		return
	}

	categoryIds := make([]int64, len(podcast.Categories))
	for i, category := range podcast.Categories {
		categoryIds[i] = category.CategoryId
	}
	categories, err := c.App.GetCategoriesByIds(categoryIds)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.Data = &model.ApiResponseData{
		Podcasts:   []*model.Podcast{podcast},
		Episodes:   episodes,
		Categories: categories,
	}
}

func BrowsePodcastEpisodes(c *Context, w http.ResponseWriter, req *http.Request) {
	episodes, err := c.App.GetEpisodesInPodcast(c.Params.PodcastId, c.Params.Order, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		episodeIds := make([]int64, len(episodes))
		for i, episode := range episodes {
			episodeIds[i] = episode.Id
		}

		playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Session.UserId, episodeIds)
		if err != nil {
			c.Err = err
			return
		}
		model.EpisodesJoinPlaybacks(episodes, playbacks)
	}

	c.Response.Data = &model.ApiResponseData{
		Episodes: episodes,
	}
}
