package api

import (
	"net/http"

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
	requireSession bool
	isStatic       bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{store: h.store}
	h.handleFunc(c, w, req)
}

func (api *Api) NewHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		store:          api.store,
		handleFunc:     h,
		requireSession: false,
		isStatic:       false,
	}
}
