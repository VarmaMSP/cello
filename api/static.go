package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *Api) RegisterStatichandlers() {
	api.router.HandlerFunc("GET", "/_next/*filepath", api.ServeStatic)
	api.router.NotFound = http.FileServer(http.Dir("/var/www/"))
}

func (api *Api) ServeStatic(w http.ResponseWriter, req *http.Request) {
	filepath := httprouter.ParamsFromContext(req.Context()).ByName("filepath")
	file, _ := http.Get("http://localhost:3000/_next" + filepath)

	w.Header().Set("Content-Type", file.Header.Get("Content-Type"))
	io.Copy(w, file.Body)
}
