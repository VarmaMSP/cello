package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type Context struct {
	store store.Store
	err   *model.AppError
}

type Handler struct {
	store          store.Store
	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	enableCors     bool
	requireSession bool
	isStatic       bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{store: h.store}

	if h.enableCors {
		w.Header().Set(headers.AccessControlAllowOrigin, req.Header.Get(headers.Origin))
		w.Header().Set(headers.AccessControlAllowCredentials, "true")
	}

	h.handleFunc(c, w, req)
}

func (api *Api) NewHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		store:          api.store,
		handleFunc:     h,
		enableCors:     true,
		requireSession: false,
		isStatic:       false,
	}
}
