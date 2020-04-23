package service

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

func subscribeToPodcast(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequirePodcastId(); c.Err != nil {
		return
	}

	if err := c.App.AddPodcastToSubscriptions(c.Params.UserId, c.Params.PodcastId); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}

func unsubscribeToPodcast(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequirePodcastId(); c.Err != nil {
		return
	}

	if err := c.App.RemovePodcastFromSubscriptions(c.Params.UserId, c.Params.PodcastId); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}
