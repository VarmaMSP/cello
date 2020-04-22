package web

import (
	"net/http"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type Web struct {
	App    *app.App
	Config *model.Config
}

func (w *Web) H(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return w.App.SessionManager.LoadAndSave(&Handler{
		App:            w.App,
		HandleFunc:     h,
		RequireSession: false,
	})
}

func (w *Web) HAuth(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return w.App.SessionManager.LoadAndSave(&Handler{
		App:            w.App,
		HandleFunc:     h,
		RequireSession: true,
	})
}
