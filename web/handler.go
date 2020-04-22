package web

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/app_"
	"github.com/varmamsp/cello/model"
)

type Handler struct {
	App            *app_.App
	HandleFunc     func(*Context, http.ResponseWriter, *http.Request)
	RequireSession bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{}

	c.App = h.App
	c.Params = ParamsFromRequest(req)
	c.Session = h.App.GetSession(req.Context())
	c.Response = &model.ApiResponse{
		Headers: map[string]string{},
	}

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

	} else if c.Response.StatusCode == http.StatusNotModified {
		for k, v := range c.Response.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(http.StatusNotModified)

	} else if c.Response.StatusCode == http.StatusOK {
		c.Response.Status = "success"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(c.Response.ToJson())
	}
}
