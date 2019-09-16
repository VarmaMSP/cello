package api

import (
	"net/http"

	"github.com/dghubble/gologin/v2/google"
	"github.com/go-http-utils/headers"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/google/login", api.app.GoogleSignIn())
	api.router.Handler("GET", "/google/callback", api.app.GoogleCallback(api.NewHandler(LoginWithGoogle)))
}

func LoginWithGoogle(c *Context, res http.ResponseWriter) {
	googleUser, err := google.UserFromContext(c.req.Context())
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
	res.Header().Set(headers.Location, "/")
	res.WriteHeader(http.StatusFound)
}
