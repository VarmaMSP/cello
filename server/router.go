package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/api/browse"
	"github.com/varmamsp/cello/api/page"
	"github.com/varmamsp/cello/api/service"
	"github.com/varmamsp/cello/api/session"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/web"
)

func newRouter(app *app.App) http.Handler {
	web := &web.Web{App: app}
	r := httprouter.New()

	// SignIn for Web
	r.Handler("GET", "/signin/google", web.H_(session.LoginWithGoogle))
	r.Handler("GET", "/signin/facebook", web.H_(session.LoginWithFacebook))
	r.Handler("GET", "/signin/twitter", web.H_(session.LoginWithTwitter))
	r.Handler("GET", "/callback/google", web.H_(session.GoogleLoginCallback))
	r.Handler("GET", "/callback/facebook", web.H_(session.FacebookLoginCallback))
	r.Handler("GET", "/callback/twitter", web.H_(session.TwitterLoginCallback))
	// SignIn for mobile
	r.Handler("POST", "/mobile/signin/guest", web.H(session.LoginWithGuest))
	r.Handler("POST", "/mobile/signin/google", web.H(session.LoginWithGoogleMobile))
	r.Handler("POST", "/mobile/signin/facebook", web.H(session.LoginWithGuest))

	r.Handler("GET", "/", web.H(page.Home))
	r.Handler("GET", "/results", web.H(page.Results))
	r.Handler("GET", "/history", web.HAuth(page.History))
	r.Handler("GET", "/subscriptions", web.HAuth(page.Subscriptions))
	r.Handler("GET", "/playlists", web.HAuth(page.Library))

	r.Handler("GET", "/suggest", web.H(page.Suggest))

	r.Handler("GET", "/charts/:chartUrlParam", web.H(page.Chart))
	r.Handler("GET", "/podcasts/:podcastUrlParam", web.H(page.Podcast))
	r.Handler("GET", "/podcasts/:podcastUrlParam/search", web.H(page.PodcastSearch))
	r.Handler("GET", "/episodes/:episodeUrlParam", web.H(page.Episode))
	r.Handler("GET", "/playlists/:playlistUrlParam", web.H(page.Playlist))

	r.Handler("GET", "/ajax/browse", web.H(browse.RootHandler))
	r.Handler("POST", "/ajax/service", web.H(service.RootHandler))

	return r
}
