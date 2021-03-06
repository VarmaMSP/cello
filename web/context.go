package web

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (c *Context) SetNotPermitted(resource string) {
	c.Err = model.NewAppError("api.context.no_permission", resource, http.StatusBadRequest, nil)
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
		c.Params.LoadFromBody(body)
	}

	return c
}

func (c *Context) RequireSession() *Context {
	if c.Err != nil {
		return c
	}

	if c.Session == nil || c.Session.UserId == 0 {
		c.SetSessionExpired()
	} else {
		c.Params.UserId = c.Session.UserId
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
	if c.Err == nil && c.Params.EpisodeId == 0 {
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

func (c *Context) RequireEpisodeIds() *Context {
	if c.Err == nil && (c.Params.EpisodeIds == nil || len(c.Params.EpisodeIds) == 0) {
		c.SetInvalidUrlParam("episode_ids")
	}

	return c
}

func (c *Context) RequireQuery() *Context {
	if c.Err == nil && (c.Params.Query == "" || len(strings.TrimSpace(c.Params.Query)) == 0) {
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

func (c *Context) RequireGoogleIdToken() *Context {
	if c.Err == nil && c.Params.GoogleIdToken == "" {
		c.SetInvalidBodyParam("google_id_token")
	}
	return c
}

func (c *Context) RequireFacebookAccessToken() *Context {
	if c.Err == nil && c.Params.FacebookAccessToken == "" {
		c.SetInvalidBodyParam("facebook_access_token")
	}
	return c
}

func (c *Context) RequireGuestAccount() *Context {
	if c.Err == nil && c.Params.GuestAccount == nil {
		c.SetInvalidQueryParam("guest_account")
	}
	return c
}
