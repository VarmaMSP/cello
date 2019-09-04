package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/store"
)

type Api struct {
	store    store.Store
	esClient *elastic.Client

	enableCors bool

	server *http.Server
	router *httprouter.Router
}

func NewApi(store store.Store, client *elastic.Client) *Api {
	api := &Api{
		store:    store,
		esClient: client,
		router:   httprouter.New(),
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
