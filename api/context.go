package api

import (
	"encoding/json"
	"net/http"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Context struct {
	App      *app.App
	Params   *Params
	Body     map[string]interface{}
	Response *model.ApiResponse
	Session  *model.Session
	Err      *model.AppError
}

func (c *Context) SetSessionExpired() {
	c.Err = model.NewAppError("api.context.session_expired", "", http.StatusUnauthorized, nil)
}

func (c *Context) SetInvalidUrlParam(param string) {
	c.Err = model.NewAppError("api.context.invalid_url_param", param, http.StatusBadRequest, nil)
}

func (c *Context) SetInvalidQueryParam(param string) {
	c.Err = model.NewAppError("api.context.invalid_query_param", param, http.StatusBadRequest, nil)
}

func (c *Context) SetInvalidBodyParam(param string) {
	c.Err = model.NewAppError("api.context.invalid_body_param", param, http.StatusBadRequest, nil)
}

func (c *Context) SetError(err error) {
	c.Err = model.NewAppError("api.generic_error", err.Error(), http.StatusBadRequest, nil)
}

func (c *Context) RequireBody(req *http.Request) *Context {
	if c.Err != nil {
		return c
	}

	body := map[string]interface{}{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&body); err != nil {
		c.Err = model.NewAppError("api.context.invalid_body", err.Error(), http.StatusBadRequest, nil)
	} else {
		c.Body = body
	}

	return c
}

func (c *Context) RequireSession() *Context {
	if c.Err != nil {
		return c
	}

	if c.Params.UserId == 0 {
		c.Params.UserId = c.Session.UserId
	}

	if c.Session.UserId == 0 {
		c.SetSessionExpired()
	}

	return c
}

func (c *Context) RequireUserId() *Context {
	if c.Err == nil && c.Params.UserId == 0 {
		c.Params.UserId = c.Session.UserId
	}
	return c
}

func (c *Context) RequirePodcastId() *Context {
	if c.Err == nil && c.Params.PodcastId == 0 {
		c.SetInvalidUrlParam("podcast_id")
	}
	return c
}

func (c *Context) RequireEpisodeId() *Context {
	if c.Params.EpisodeId == 0 {
		c.SetInvalidUrlParam("episode_id")
	}

	return c
}

func (c *Context) RequirePlaylistId() *Context {
	if c.Err == nil && c.Params.PlaylistId == 0 {
		c.SetInvalidUrlParam("playlist_id")
	}
	return c
}

func (c *Context) RequireChartId() *Context {
	if c.Err == nil && c.Params.ChartId == "" {
		c.SetInvalidUrlParam("chart_id")
	}
	return c
}

func (c *Context) RequireQuery() *Context {
	if c.Err == nil && c.Params.Query == "" {
		c.SetInvalidQueryParam("query")
	}
	return c
}

func (c *Context) RequireEndpoint() *Context {
	if c.Err == nil && c.Params.Endpoint == "" {
		c.SetInvalidQueryParam("endpoint")
	}
	return c
}

func (c *Context) RequireAction() *Context {
	if c.Err == nil && c.Params.Action == "" {
		c.SetInvalidQueryParam("action")
	}
	return c
}
