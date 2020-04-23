package session

import (
	"net/http"

	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	twitterLogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/varmamsp/cello/web"
)

func LoginWithGoogle(c *web.Context, w http.ResponseWriter, req *http.Request) {
	googleLogin.StateHandler(
		cookieConfig(c),
		googleLogin.LoginHandler(googleOAuthConfig(c), nil),
	).ServeHTTP(w, req)
}

func LoginWithFacebook(c *web.Context, w http.ResponseWriter, req *http.Request) {
	facebookLogin.StateHandler(
		cookieConfig(c),
		facebookLogin.LoginHandler(facebookOAuthConfig(c), nil),
	).ServeHTTP(w, req)
}

func LoginWithTwitter(c *web.Context, w http.ResponseWriter, req *http.Request) {
	twitterLogin.LoginHandler(twitterOAuthConfig(c), nil).ServeHTTP(w, req)
}

func GoogleLoginCallback(c *web.Context, w http.ResponseWriter, req *http.Request) {
	googleLogin.StateHandler(
		cookieConfig(c),
		googleLogin.CallbackHandler(
			googleOAuthConfig(c),
			c.App.SessionManager.LoadAndSave(
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					if user, err := c.App.CreateUserWithGoogle(req.Context()); err == nil {
						c.App.NewSession(req.Context(), user)
					}

					w.Header().Set("Location", c.App.HostName)
					w.WriteHeader(http.StatusFound)
				}),
			),
			nil,
		),
	).ServeHTTP(w, req)
}

func FacebookLoginCallback(c *web.Context, w http.ResponseWriter, req *http.Request) {
	facebookLogin.StateHandler(
		cookieConfig(c),
		facebookLogin.CallbackHandler(
			facebookOAuthConfig(c),
			c.App.SessionManager.LoadAndSave(
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					if user, err := c.App.CreateUserWithFacebook(req.Context()); err == nil {
						c.App.NewSession(req.Context(), user)
					}

					w.Header().Set("Location", c.App.HostName)
					w.WriteHeader(http.StatusFound)
				}),
			),
			nil,
		),
	).ServeHTTP(w, req)
}

func TwitterLoginCallback(c *web.Context, w http.ResponseWriter, req *http.Request) {
	twitterLogin.CallbackHandler(
		twitterOAuthConfig(c),
		c.App.SessionManager.LoadAndSave(
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if user, err := c.App.CreateUserWithTwitter(req.Context()); err == nil {
					c.App.NewSession(req.Context(), user)
				}

				w.Header().Set("Location", c.App.HostName)
				w.WriteHeader(http.StatusFound)
			}),
		),
		nil,
	).ServeHTTP(w, req)
}
