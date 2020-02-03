package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func ServiceLoadSession(c *Context, w http.ResponseWriter, req *http.Request) {
	if c.Session == nil {
		return
	}

	user, err := c.App.GetUser(c.Session.UserId)
	if err != nil {
		c.Err = err
		return
	}
	user.Sanitize()

	subscriptions, err := c.App.GetUserSubscriptions(c.Session.UserId)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.Data = &model.ApiResponseData{
		Users:    []*model.User{user},
		Podcasts: subscriptions,
	}
}

func ServiceEndSession(c *Context, w http.ResponseWriter, req *http.Request) {
	err := c.App.DeleteSession(req.Context())
	if err != nil {
		c.Err = err
		return
	}
}
