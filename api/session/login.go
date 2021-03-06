package session

import (
	"errors"
	"net/http"

	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	facebook "github.com/huandu/facebook/v2"
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
		} else if err := c.App.LinkGoogleToUser(user, googleUser); err != nil {
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

	res, err := facebook.Get("/me", facebook.Params{
		"fields":       "name,email",
		"access_token": c.Params.FacebookAccessToken,
	})
	if err != nil {
		c.SetError(errors.New("Invalid access token"))
	}
	println(res)

	var id string
	var name string
	var email string
	res.DecodeField("id", &id)
	res.DecodeField("name", &name)
	res.DecodeField("email", &email)

	if c.Params.GuestAccount != nil {
		if user, err := c.App.CreateUserWithGuest(c.Params.GuestAccount); err != nil {
			c.Err = err
		} else if err := c.App.LinkFacebookToUser(user, id, name, email); err != nil {
			c.Err = err
		} else {
			c.App.NewSession(req.Context(), user)
			c.Response.StatusCode = http.StatusOK
			c.Response.Data = &model.ApiResponseData{}
		}
	} else {
		if user, err := c.App.CreateUserWithFacebook(id, name, email); err != nil {
			c.Err = err
		} else {
			c.App.NewSession(req.Context(), user)
			c.Response.StatusCode = http.StatusOK
			c.Response.Data = &model.ApiResponseData{}
		}
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
						// TODO: Handle failure instead of just redirecting the user back to site
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
					if fbUser, err := facebookLogin.UserFromContext(req.Context()); err != nil {
						c.SetError(err)
						return
					} else if user, err := c.App.CreateUserWithFacebook(fbUser.ID, fbUser.Name, fbUser.Email); err != nil {
						// TODO: Handle failure instead of just redirecting the user back to site
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
