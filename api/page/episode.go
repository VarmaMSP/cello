package page

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Episode(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireEpisodeId(); c.Err != nil {
		return
	}

	episode, err := c.App.GetEpisode(c.Params.EpisodeId)
	if err != nil {
		c.Err = err
		return
	}

	podcast, err := c.App.GetPodcast(episode.PodcastId)
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
