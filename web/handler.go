package web

import (
	"fmt"
	"net/http"

	"github.com/avct/uasurfer"
	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Handler struct {
	App            *app.App
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
		Data: &model.ApiResponseData{
			SearchResults: &model.SearchResults{},
		},
	}

	ua := uasurfer.Parse(req.Header.Get(headers.UserAgent))
	c.App.Log.Info().
		Str("method", req.Method).
		Str("path", req.URL.String()).
		Str("device", ua.DeviceType.StringTrimPrefix()).
		Str("user_agent", fmt.Sprintf("%s-%s", ua.OS.Platform.StringTrimPrefix(), ua.Browser.Name.StringTrimPrefix())).
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

type Handler_ struct {
	App        *app.App
	HandleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h *Handler_) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{App: h.App}
	h.HandleFunc(c, w, req)
}
