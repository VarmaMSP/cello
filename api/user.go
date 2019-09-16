package api

import (
	"net/http"

	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	"github.com/go-http-utils/headers"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/google/login", api.app.GoogleSignIn())
	api.router.Handler("GET", "/google/callback", api.app.GoogleCallback(api.NewHandler(LoginWithGoogle)))
	api.router.Handler("GET", "/facebook/login", api.app.FacebookSignIn())
	api.router.Handler("GET", "/facebook/callback", api.app.FacebookCallback(api.NewHandler(LoginWithFacebook)))
}

func LoginWithGoogle(c *Context, res http.ResponseWriter) {
	googleUser, err := googleLogin.UserFromContext(c.req.Context())
	if err != nil {
		c.err = nil
		return
	}

	user, err_ := c.app.CreateUserWithGoogle(googleUser)
	if err_ != nil {
		c.err = err_
		return
	}
	c.app.CreateSession(c.req.Context(), user)
	res.Header().Set(headers.Location, "http://localhost:8080")
	res.WriteHeader(http.StatusFound)
}

func LoginWithFacebook(c *Context, res http.ResponseWriter) {
	facebookUser, err := facebookLogin.UserFromContext(c.req.Context())
	if err != nil {
		c.err = nil
		return
	}

	user, err_ := c.app.CreateUserWithFacebook(facebookUser)
	if err_ != nil {
		c.err = err_
		return
	}
	c.app.CreateSession(c.req.Context(), user)
	res.Header().Set(headers.Location, "http://localhost:8080")
	res.WriteHeader(http.StatusFound)
}
