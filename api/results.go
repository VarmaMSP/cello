package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
)

func GetResultsPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireQuery()
	if c.Err != nil {
		return
	}

	podcasts, err := c.App.SearchPodcasts(c.Params.Query)
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
