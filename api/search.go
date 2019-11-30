package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterSearchHandlers() {
	api.router.Handler("GET", "/results", api.NewHandler(SearchPodcasts))
}

func SearchPodcasts(c *Context, w http.ResponseWriter) {
	req := &SearchPodcastsReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.search_podcasts_req.load", err.Error(), 400, nil)
		return
	}

	podcasts, err := c.app.SearchPodcasts(req.SearchQuery)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"totalCount": len(podcasts),
		"results":    podcasts,
	})

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
