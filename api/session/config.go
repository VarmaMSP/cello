package session

import (
	"github.com/dghubble/gologin/v2"
	"github.com/varmamsp/cello/web"
	"golang.org/x/oauth2"
	facebookOAuth "golang.org/x/oauth2/facebook"
	googleOAuth "golang.org/x/oauth2/google"
)

func cookieConfig(c *web.Context) gologin.CookieConfig {
	if c.App.Dev {
		return gologin.DebugOnlyCookieConfig
	}
	return gologin.DefaultCookieConfig
}

func googleOAuthConfig(c *web.Context) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.App.Config.OAuth.Google.ClientId,
		ClientSecret: c.App.Config.OAuth.Google.ClientSecret,
		RedirectURL:  c.App.Config.OAuth.Google.RedirectUrl,
		Scopes:       c.App.Config.OAuth.Google.Scopes,
		Endpoint:     googleOAuth.Endpoint,
	}
}

func facebookOAuthConfig(c *web.Context) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.App.Config.OAuth.Facebook.ClientId,
		ClientSecret: c.App.Config.OAuth.Facebook.ClientSecret,
		RedirectURL:  c.App.Config.OAuth.Facebook.RedirectUrl,
		Scopes:       c.App.Config.OAuth.Facebook.Scopes,
		Endpoint:     facebookOAuth.Endpoint,
	}
}
