package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/api/browse"
	"github.com/varmamsp/cello/api/service"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/web"
)

func newRouter(app *app.App) http.Handler {
	web := &web.Web{App: app}
	r := httprouter.New()

	r.Handler("GET", "/ajax/browse", web.H(browse.RootHandler))
	r.Handler("POST", "/ajax/service", web.H(service.RootHandler))

	return r
}
