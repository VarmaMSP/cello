package service

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func loadSession(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.Session == nil {
		return
	}

	user, err := c.App.GetUser(c.Session.UserId)
	if err != nil {
		c.Err = err
		return
	}
	user.Sanitize()

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Users: []*model.User{user},
	}
}

func endSession(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if err := c.App.DeleteSession(req.Context()); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}
