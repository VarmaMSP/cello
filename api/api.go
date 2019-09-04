package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/store"
)

type Api struct {
	store store.Store

	enableCors bool

	server *http.Server
	router *httprouter.Router
}

func NewApi(store store.Store) *Api {
	api := &Api{
		store:  store,
		router: httprouter.New(),
	}

	api.server = &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: api.router,
	}

	api.RegisterStatichandlers()
	api.RegisterPodcastHandlers()

	return api
}

func (api *Api) ListenAndServe() {
	err := api.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
