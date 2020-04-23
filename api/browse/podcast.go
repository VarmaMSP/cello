package browse

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func podcastEpisodes(c *web.Context, w http.ResponseWriter, req *http.Request) {
	episodes, err := c.App.GetEpisodesFromPodcast(c.Params.PodcastId, c.Params.Order, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		if err := c.App.LoadPlaybacks(c.Session.UserId, episodes); err != nil {
			c.Err = err
			return
		}
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Episodes: episodes,
	}
}
