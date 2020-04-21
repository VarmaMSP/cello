package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func GetEpisode(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireEpisodeId()
	if c.Err != nil {
		return
	}

	episode, err := c.App.GetEpisode(c.Params.EpisodeId)
	if err != nil {
		c.Err = err
		return
	}

	podcast, err := c.App.GetPodcast(episode.PodcastId, false)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Podcasts: []*model.Podcast{podcast},
		Episodes: []*model.Episode{episode},
	}
}
