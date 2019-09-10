package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type Context struct {
	req *http.Request

	store    store.Store
	esClient *elastic.Client

	err *model.AppError
}

func (c *Context) Query(key string) string {
	return c.req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	return httprouter.ParamsFromContext(c.req.Context()).ByName(key)
}

func (c *Context) Body() (m map[string]string) {
	return model.MapFromJson(c.req.Body)
}

type Handler struct {
	store    store.Store
	esClient *elastic.Client

	handleFunc     func(*Context, http.ResponseWriter)
	enableCors     bool
	requireSession bool
	isStatic       bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		req:      req,
		store:    h.store,
		esClient: h.esClient,
	}

	if h.enableCors {
		w.Header().Set(headers.AccessControlAllowOrigin, req.Header.Get(headers.Origin))
		w.Header().Set(headers.AccessControlAllowCredentials, "true")
	}

	h.handleFunc(c, w)

	if c.err != nil {
		w.WriteHeader(c.err.StatusCode)
		w.Write([]byte(c.err.Error()))
	}
}

func (api *Api) NewHandler(h func(*Context, http.ResponseWriter)) http.Handler {
	return &Handler{
		store:          api.store,
		esClient:       api.esClient,
		handleFunc:     h,
		enableCors:     true,
		requireSession: false,
		isStatic:       false,
	}
}
