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
	w := &web.Web{App: app}
	r := httprouter.New()

	r.Handler("GET", "/", w.H(page.Home))
	r.Handler("GET", "/results", w.H(page.Results))
	r.Handler("GET", "/history", w.HAuth(page.History))
	r.Handler("GET", "/subscriptions", w.HAuth(page.Subscriptions))
	r.Handler("GET", "/playlists", w.HAuth(page.Library))

	r.Handler("GET", "/suggest", w.H(page.Suggest))

	r.Handler("GET", "/charts/:chartUrlParam", w.H(page.Chart))
	r.Handler("GET", "/podcasts/:podcastUrlParam", w.H(page.Podcast))
	r.Handler("GET", "/podcasts/:podcastUrlParam/search", w.H(page.PodcastSearch))
	r.Handler("GET", "/episodes/:episodeUrlParam", w.H(page.Episode))
	r.Handler("GET", "/playlists/:playlistUrlParam", w.H(page.Playlist))

	r.Handler("GET", "/ajax/browse", w.H(browse.RootHandler))
	r.Handler("POST", "/ajax/service", w.H(service.RootHandler))

	r.Handler("GET", "/signin/google", w.H_(session.LoginWithGoogle))
	r.Handler("GET", "/signin/facebook", w.H_(session.LoginWithFacebook))
	r.Handler("GET", "/callback/google", w.H_(session.GoogleLoginCallback))
	r.Handler("GET", "/callback/facebook", w.H_(session.FacebookLoginCallback))

	r.Handler("POST", "/mobile/signin/guest", w.H(session.LoginWithGuest))
	r.Handler("POST", "/mobile/signin/google", w.H(session.LoginWithGoogleMobile))
	r.Handler("POST", "/mobile/signin/facebook", w.H(session.LoginWithFacebookMobile))

	r.Handler("GET", "/health", w.H_(
		func(c *web.Context, w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	))

	return r
}
