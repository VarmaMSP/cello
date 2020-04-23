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

	episodes, err := c.App.GetRecentlyPlayedEpisodes(c.Params.UserId, c.Params.Offset, c.Params.Offset)
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
		Podcasts: podcasts,
		Episodes: episodes,
	}
}
