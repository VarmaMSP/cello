package browse

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func searchResults(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireQuery(); c.Err != nil {
		return
	}

	if c.Params.Type == "podcast" {
		results, err := c.App.SearchPodcasts(c.Params.Query, c.Params.Offset, c.Params.Limit)
		if err != nil {
			c.Err = err
			return
		}

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Podcasts: results,
		}

	} else if c.Params.Type == "episode" {
		results, err := c.App.SearchEpisodes(c.Params.Query, c.Params.SortBy, c.Params.Offset, c.Params.Limit)
		if err != nil {
			c.Err = err
			return
		}

		c.Response.StatusCode = http.StatusOK
		c.Response.Data = &model.ApiResponseData{
			Episodes: results,
		}

	} else {
		c.SetInvalidQueryParam("type")
	}
}
