package service

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func getPlaybacks(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequireEpisodeIds(); c.Err != nil {
		return
	}

	playbacks, err := c.App.GetPlaybacksForEpisodes(c.Params.UserId, c.Params.EpisodeIds)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Playbacks: playbacks,
	}
}

func syncPlaybackBegin(c *web.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Println("syncing")
	if c.RequireSession().RequireBody(req).RequireEpisodeId(); c.Err != nil {
		return
	}

	if err := c.App.SyncPlaybackBegin(c.Params.UserId, c.Params.EpisodeId); err != nil {
		return
	}

	c.Response.StatusCode = http.StatusOK
}

func syncPlaybackProgress(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequireEpisodeId(); c.Err != nil {
		return
	}

	position, ok := c.Body["position"].(float64)
	if !ok {
		c.SetInvalidBodyParam("position")
		return
	}

	if err := c.App.SyncPlaybackProgress(c.Params.UserId, c.Params.EpisodeId, position); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}
