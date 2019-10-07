package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/me", api.NewHandlerSessionRequired(Me))
	api.router.Handler("GET", "/signin/google", api.app.GoogleSignIn())
	api.router.Handler("GET", "/signin/facebook", api.app.FacebookSignIn())
	api.router.Handler("GET", "/signin/twitter", api.app.TwitterSignIn())
	api.router.Handler("GET", "/callback/google", api.app.GoogleCallback(api.NewHandler(SignInWithSocial("google"))))
	api.router.Handler("GET", "/callback/facebook", api.app.FacebookCallback(api.NewHandler(SignInWithSocial("facebook"))))
	api.router.Handler("GET", "/callback/twitter", api.app.TwitterCallback(api.NewHandler(SignInWithSocial("twitter"))))
	api.router.Handler("GET", "/signout", api.NewHandler(SignOut))
}

func Me(c *Context, w http.ResponseWriter) {
	user, err := c.app.GetUser(c.session.UserId)
	if err != nil {
		c.err = err
		return
	}

	subscriptions, err := c.app.GetUserSubscriptions(c.session.UserId)
	if err != nil {
		c.err = err
		return
	}
	for _, subscription := range subscriptions {
		subscription.SanitizeToMin()
	}

	res, _ := json.Marshal(map[string]interface{}{
		"user":          user,
		"subscriptions": subscriptions,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SignInWithSocial(accountType string) func(*Context, http.ResponseWriter) {
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
		w.Header().Set(headers.Location, c.app.HostName)
		w.WriteHeader(http.StatusFound)
	}
}

func SignOut(c *Context, w http.ResponseWriter) {
	err := c.app.DeleteSession(c.req.Context())
	if err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
