package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/google/login", api.app.GoogleSignIn())
	api.router.Handler("GET", "/facebook/login", api.app.FacebookSignIn())
	api.router.Handler("GET", "/twitter/login", api.app.TwitterSignIn())
	api.router.Handler("GET", "/google/callback", api.app.GoogleCallback(api.NewHandler(LoginWithSocial("google"))))
	api.router.Handler("GET", "/facebook/callback", api.app.FacebookCallback(api.NewHandler(LoginWithSocial("facebook"))))
	api.router.Handler("GET", "/twitter/callback", api.app.TwitterCallback(api.NewHandler(LoginWithSocial("twitter"))))
}

func LoginWithSocial(accountType string) func(c *Context, res http.ResponseWriter) {
	return func(c *Context, res http.ResponseWriter) {
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
		res.Header().Set(headers.Location, "http://localhost:8080")
		res.WriteHeader(http.StatusFound)
	}
}
