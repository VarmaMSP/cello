package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Context struct {
	req     *http.Request
	app     *app.App
	session *model.Session
	err     *model.AppError
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

func (c *Context) DecodeBody(i interface{}) error {
	decoder := json.NewDecoder(c.req.Body)
	return decoder.Decode(i)
}

type Handler struct {
	app            *app.App
	handleFunc     func(*Context, http.ResponseWriter)
	requireSession bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		app:     h.app,
		req:     req,
		session: h.app.GetSession(req.Context()),
	}

	c.app.Log.Info().
		Str("method", req.Method).
		Str("path", req.URL.String()).
		Str("user_agent", req.Header.Get(headers.UserAgent)).
		Msg("")

	if h.requireSession && c.session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	h.handleFunc(c, w)

	if c.err != nil {
		c.app.Log.Error().
			Str("from", c.err.Id).
			Str("error", c.err.DetailedError).
			Msg("")

		w.WriteHeader(c.err.StatusCode)
		w.Write([]byte(c.err.Error()))
		return
	}
}

func (api *Api) NewHandler(h func(*Context, http.ResponseWriter)) http.Handler {
	return api.app.SessionManager.LoadAndSave(&Handler{
		app:            api.app,
		handleFunc:     h,
		requireSession: false,
	})
}

func (api *Api) NewHandlerSessionRequired(h func(*Context, http.ResponseWriter)) http.Handler {
	return api.app.SessionManager.LoadAndSave(&Handler{
		app:            api.app,
		handleFunc:     h,
		requireSession: true,
	})
}
