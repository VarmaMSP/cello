package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Api struct {
	app        *app.App
	enableCors bool

	server *http.Server
	router *httprouter.Router
}

func NewApi(config model.Config) *Api {
	api := &Api{}
	api.app, _ = app.NewApp(config)
	api.router = httprouter.New()
	api.server = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: api.router,
	}

	api.RegisterPodcastHandlers()
	api.RegisterCurationRoutes()
	api.RegisterUserHandlers()

	return api
}

func (api *Api) ListenAndServe() {
	api.app.Log.Info().Msg("Server listening on port: 8081")
	err := api.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
