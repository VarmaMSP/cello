package api

import (
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/store"
)

type Api struct {
	store store.Store

	server     *http.Server
	router     *httprouter.Router
	listenAddr *net.TCPAddr
}

func NewApi(store store.Store) *Api {
	api := &Api{
		store:  store,
		router: httprouter.New(),
		listenAddr: &net.TCPAddr{
			IP:   []byte("127.0.0.1"),
			Port: 8081,
		},
	}

	api.server = &http.Server{
		Addr:    api.listenAddr.String(),
		Handler: api.router,
	}

	api.RegisterStatichandlers()
	api.RegisterPodcastHandlers()

	return api
}

func (api *Api) ListenAndServe() {
	api.server.ListenAndServe()
}
