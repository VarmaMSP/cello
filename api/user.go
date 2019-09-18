package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/me", api.NewHandlerSessionRequired(Me))
	api.router.Handler("GET", "/google/signin", api.app.GoogleSignIn())
	api.router.Handler("GET", "/facebook/signin", api.app.FacebookSignIn())
	api.router.Handler("GET", "/twitter/signin", api.app.TwitterSignIn())
	api.router.Handler("GET", "/google/callback", api.app.GoogleCallback(api.NewHandler(LoginWithSocial("google"))))
	api.router.Handler("GET", "/facebook/callback", api.app.FacebookCallback(api.NewHandler(LoginWithSocial("facebook"))))
	api.router.Handler("GET", "/twitter/callback", api.app.TwitterCallback(api.NewHandler(LoginWithSocial("twitter"))))
}

func Me(c *Context, w http.ResponseWriter) {
	user, err := c.app.GetUser(c.session.UserId)
	if err != nil {
		c.err = err
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"user": user,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginWithSocial(accountType string) func(*Context, http.ResponseWriter) {
	return func(c *Context, w http.ResponseWriter) {
		var user *model.User
		var err *model.AppError

		if accountType == "google" {
			user, err = c.app.CreateUserWithGoogle(c.req.Context())
		} else if accountType == "facebook" {
			user, err = c.app.CreateUserWithFacebook(c.req.Context())
		} else if accountType == "twitter" {
			user, err = c.app.CreateUserWithTwitter(c.req.Context())
		}

		if err != nil {
			c.err = err
			return
		}
		c.app.CreateSession(c.req.Context(), user)
		w.Header().Set(headers.Location, "http://localhost:8080")
		w.WriteHeader(http.StatusFound)
	}
}
