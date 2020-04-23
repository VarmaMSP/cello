package page

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func Playlist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequirePlaylistId(); c.Err != nil {
		return
	}

	playlist, err := c.App.GetPlaylist(c.Params.PlaylistId)
	if err != nil {
		c.Err = err
		return
	}

	if !c.App.HasPermissionToViewPlaylist(c.Session, playlist) {
		c.Response.StatusCode = http.StatusBadRequest
		return
	}

	episodes, err := c.App.GetPlaylistEpisodes(playlist)
	if err != nil {
		c.Err = err
		return
	}

	podcasts, err := c.App.GetPodcastsForEpisodes(episodes)
	if err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		if err := c.App.LoadPlaybacks(c.Session.UserId, episodes); err != nil {
			c.Err = err
			return
		}
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Playlists: []*model.Playlist{playlist},
		Episodes:  episodes,
		Podcasts:  podcasts,
	}
}
