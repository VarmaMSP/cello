package api

import (
	"io"
	"net/http"
)

func (api *Api) RegisterStatichandlers() {
	api.router.Handler("GET", "/_next/*filepath", api.NewHandler(ServeStatic))
	api.router.NotFound = http.FileServer(http.Dir("/var/www/"))
}

func ServeStatic(c *Context, w http.ResponseWriter) {
	filepath := c.Param("filepath")
	file, _ := http.Get("http://localhost:3000/_next" + filepath)

	w.Header().Set("Content-Type", file.Header.Get("Content-Type"))
	io.Copy(w, file.Body)
}
