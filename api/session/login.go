package session

import (
	"net/http"

	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	twitterLogin "github.com/dghubble/gologin/v2/twitter"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
	google "google.golang.org/api/oauth2/v2"
)

// For Mobile
func LoginWithGuest(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireBody(req).RequireGuestAccount(); c.Err != nil {
		return
	}

	if user, err := c.App.CreateUserWithGuest(c.Params.GuestAccount); err != nil {
		c.Err = err
	} else {
		c.App.NewSession(req.Context(), user)
		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{}
	}
}

func LoginWithGoogleMobile(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireBody(req).RequireGoogleIdToken(); c.Err != nil {
		return
	}

	v := googleAuthIDTokenVerifier.Verifier{}
	if err := v.VerifyIDToken(c.Params.GoogleIdToken, []string{googleOAuthConfig(c).ClientID}); err != nil {
		c.SetError(err)
		return
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(c.Params.GoogleIdToken)
	if err != nil {
		c.SetError(err)
		return
	}
	googleUser := &google.Userinfoplus{
		Id:            claimSet.Sub,
		Name:          claimSet.Name,
		FamilyName:    claimSet.FamilyName,
		GivenName:     claimSet.GivenName,
		Email:         claimSet.Email,
		VerifiedEmail: &claimSet.EmailVerified,
		Picture:       claimSet.Picture,
		Locale:        claimSet.Locale,
	}

	if c.Params.GuestAccount != nil {
		if user, err := c.App.CreateUserWithGuest(c.Params.GuestAccount); err != nil {
			c.Err = err
		} else if err := c.App.LinkUserToGoogle(user, googleUser); err != nil {
			c.Err = err
		} else {
			c.App.NewSession(req.Context(), user)
			c.Response.StatusCode = http.StatusOK
			c.Response.Data = &model.ApiResponseData{}
		}
	} else {
		if user, err := c.App.CreateUserWithGoogle(googleUser); err != nil {
			c.SetError(err)
		} else {
			c.App.NewSession(req.Context(), user)
			c.Response.StatusCode = http.StatusOK
			c.Response.Data = &model.ApiResponseData{}
		}
	}
}

func LoginWithFacebookMobile(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireBody(req).RequireFacebookAccessToken(); c.Err != nil {
		return
	}
}

// For Web
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
					if googleUser, err := googleLogin.UserFromContext(req.Context()); err != nil {
						c.SetError(err)
						return
					} else if user, err := c.App.CreateUserWithGoogle(googleUser); err != nil {
						// Do nothing, just redirect
					} else {
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
