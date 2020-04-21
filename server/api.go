package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Api struct {
	App    *app.App
	Config *model.Config

	Server *http.Server
	Router *httprouter.Router

	Jobs
}

func NewApi(config model.Config) (*Api, error) {
	api := &Api{}

	api.Router = httprouter.New()
	api.Server = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: api.Router,
	}

	api.RegisterHandlers()

	return api, nil
}

func (api *Api) ListenAndServe() {
	api.App.Log.Info().Msg("Server listening on port: 8081")
	err := api.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
