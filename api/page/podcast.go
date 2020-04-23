package page

import (
	"net/http"

	"github.com/varmamsp/cello/api/browse"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Podcast(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequirePodcastId(); c.Err != nil {
		return
	}

	// feed, err := c.App.GetFeed(c.Params.PodcastId)
	// if err != nil {
	// 	c.Err = err
	// 	return
	// }

	podcast, err := c.App.GetPodcast(c.Params.PodcastId)
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

	c.Params.Endpoint = browse.PODCAST_EPISODES
	if browse.RootHandler(c, w, req); c.Err == nil {
		c.Response.StatusCode = http.StatusOK
		c.Response.Data.Podcasts = []*model.Podcast{podcast}
		c.Response.Data.Categories = categories
	}
}
