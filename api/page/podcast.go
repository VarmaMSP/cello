package page

import (
	"net/http"

	hdrs "github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/api/browse"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Podcast(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequirePodcastId(); c.Err != nil {
		return
	}

	feed, err := c.App.GetFeed(c.Params.PodcastId)
	if err != nil {
		c.Err = err
		return
	}

	// Set Etag / LastModified headers of the feed
	c.Response.Headers[hdrs.CacheControl] = "private, max-age=300, must-revalidate"
	if feed.ETag != "" {
		c.Response.Headers[hdrs.ETag] = feed.ETag
		if ifNoneMatch := req.Header.Get(hdrs.IfNoneMatch); ifNoneMatch != "" && ifNoneMatch == feed.ETag {
			c.Response.StatusCode = http.StatusNotModified
			return
		}
	}
	if feed.LastModified != "" {
		c.Response.Headers[hdrs.LastModified] = feed.LastModified
		if ifModifiedSince := req.Header.Get(hdrs.IfModifiedSince); ifModifiedSince != "" && ifModifiedSince == feed.LastModified {
			c.Response.StatusCode = http.StatusNotModified
			return
		}
	}

	podcast, err := c.App.GetPodcast(c.Params.PodcastId)
	if err != nil {
		c.Err = err
		return
	}
	podcast.LoadFeedDetails(feed)

	if c.Params.MobileClient && c.Session != nil && c.Session.UserId != 0 {
		podcast.IsSubscribed, _ = c.App.IsUserSubscribedToPodcast(c.Session.UserId, c.Params.PodcastId)
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

	c.Params.Endpoint = browse.PODCAST_EPISODES
	if browse.RootHandler(c, w, req); c.Err == nil {
		c.Response.StatusCode = http.StatusOK
		c.Response.Data.Podcasts = []*model.Podcast{podcast}
		c.Response.Data.Categories = categories
	}
}

func PodcastSearch(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequirePodcastId(); c.Err != nil {
		return
	}

	podcast, err := c.App.GetPodcast(c.Params.PodcastId)
	if err != nil {
		c.Err = err
		return
	}

	c.Params.Endpoint = browse.PODCAST_SEARCH_RESULTS
	c.Params.Offset = 0
	c.Params.Limit = 25
	if browse.RootHandler(c, w, req); c.Err == nil {
		c.Response.StatusCode = http.StatusOK
		c.Response.Data.Podcasts = []*model.Podcast{podcast}
	}
}
