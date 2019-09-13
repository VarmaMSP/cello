package api

import (
	"fmt"
	"net/http"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"github.com/go-http-utils/headers"
)

func (api *Api) RegisterUserHandlers() {
	api.router.Handler("GET", "/google/login", google.StateHandler(gologin.DebugOnlyCookieConfig, google.LoginHandler(api.app.GoogleOAuthConfig, issueSession())))
	api.router.Handler("GET", "/google/callback", google.StateHandler(gologin.DebugOnlyCookieConfig, google.CallbackHandler(api.app.GoogleOAuthConfig, api.NewHandler(LoginWithGoogle), nil)))
}

func LoginWithGoogle(c *Context, res http.ResponseWriter) {
	googleUser, _ := google.UserFromContext(c.req.Context())
	fmt.Println(googleUser)
	res.Header().Set(headers.Location, "http://localhost:8080/")
	res.WriteHeader(http.StatusFound)
}

func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		err := gologin.ErrorFromContext(ctx)
		fmt.Println(err)
	}
	return http.HandlerFunc(fn)
}
