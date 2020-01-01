package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func GetResultsPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	podcasts, err := c.App.SearchPodcasts(c.Params.Query, 0, 25)
	if err != nil {
		return
	}

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"results": podcasts,
	}))
}

func BrowseResults(c *Context, w http.ResponseWriter, req *http.Request) {
	podcasts, err := c.App.SearchPodcasts(c.Params.Query, c.Params.Offset, c.Params.Limit)
	if err != nil {
		return
	}

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"results": podcasts,
	}))
}
