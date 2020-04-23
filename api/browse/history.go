package browse

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func history(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession(); c.Err != nil {
		return
	}

	episodes, err := c.App.GetRecentlyPlayedEpisodes(c.Params.UserId, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	podcasts, err := c.App.GetPodcastsForEpisodes(episodes)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Podcasts: podcasts,
		Episodes: episodes,
	}
}
