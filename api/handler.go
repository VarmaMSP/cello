package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type Context struct {
	store    store.Store
	esClient *elastic.Client
	err      *model.AppError
}

type Handler struct {
	store    store.Store
	esClient *elastic.Client

	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	enableCors     bool
	requireSession bool
	isStatic       bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{store: h.store, esClient: h.esClient}

	if h.enableCors {
		w.Header().Set(headers.AccessControlAllowOrigin, req.Header.Get(headers.Origin))
		w.Header().Set(headers.AccessControlAllowCredentials, "true")
	}

	h.handleFunc(c, w, req)
}

func (api *Api) NewHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		store:          api.store,
		esClient:       api.esClient,
		handleFunc:     h,
		enableCors:     true,
		requireSession: false,
		isStatic:       false,
	}
}
