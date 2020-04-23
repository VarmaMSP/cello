package page

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Suggest(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireQuery(); c.Err != nil {
		return
	}

	suggestions, err := c.App.GetTypeaheadSuggestions(c.Params.Query)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		SearchSuggestions: suggestions,
	}
}
