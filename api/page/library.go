package page

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Library(c *web.Context, w http.ResponseWriter, req *http.Request) {
	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Playlists: playlists,
	}
}
