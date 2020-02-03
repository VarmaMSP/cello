package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Handler struct {
	App            *app.App
	HandleFunc     func(*Context, http.ResponseWriter, *http.Request)
	RequireSession bool
}

func (api *Api) H(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return api.App.SessionManager.LoadAndSave(&Handler{
		App:            api.App,
		HandleFunc:     h,
		RequireSession: false,
	})
}

func (api *Api) HAuth(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return api.App.SessionManager.LoadAndSave(&Handler{
		App:            api.App,
		HandleFunc:     h,
		RequireSession: true,
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{}

	c.App = h.App
	c.Params = ParamsFromRequest(req)
	c.Session = h.App.GetSession(req.Context())
	c.Response = &model.ApiResponse{}

	c.App.Log.Info().
		Str("method", req.Method).
		Str("path", req.URL.String()).
		Str("user_agent", req.Header.Get(headers.UserAgent)).
		Msg("")

	if h.RequireSession {
		c.RequireSession()
	}

	if c.Err == nil {
		h.HandleFunc(c, w, req)
	}

	if c.Err != nil {
		c.App.Log.Error().
			Str("from", c.Err.GetId()).
			Str("error", c.Err.GetComment())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(c.Err.GetStatusCode())
		w.Write((&model.ApiResponse{
			Status:  "error",
			Message: c.Err.Error(),
		}).ToJson())

	} else {
		c.Response.Status = "success"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(c.Response.ToJson())
	}
}
