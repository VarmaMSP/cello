package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/app"
)

type Api struct {
	app    *app.App
	router *httprouter.Router
}

func NewApi(app *app.App) *Api {
	api := &Api{
		app:    app,
		router: httprouter.New(),
	}

	api.RegisterStatichandlers()
	api.RegisterPodcastHandlers()
	api.RegisterCurationHandlers()

	return api
}

func (api *Api) ListenAndServe() {
	http.ListenAndServe(":8081", api.router)
}
