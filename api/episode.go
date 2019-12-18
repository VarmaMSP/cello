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

	podcast, err := c.App.GetPodcast(episode.PodcastId)
	if err != nil {
		c.Err = err
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episode": episode,
		"podcast": podcast,
	}))
}


