package service

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func addToPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequireEpisodeIds(); c.Err != nil {
		return
	}

	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId, c.Params.EpisodeIds...)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Playlists: playlists,
	}
}

func createPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequireEpisodeIds(); c.Err != nil {
		return
	}

	title, ok := c.Body["title"].(string)
	if !ok {
		c.SetInvalidBodyParam("title")
	}

	privacy, ok := c.Body["privacy"].(string)
	if !ok || (privacy != "PRIVATE" && privacy != "PUBLIC") {
		c.SetInvalidBodyParam("privacy")
	}

	description, ok := c.Body["description"].(string)
	if !ok {
		c.SetInvalidBodyParam("description")
	}

	playlist, err := c.App.CreatePlaylistWithEpisodes(&model.Playlist{
		Title:       title,
		Description: description,
		Privacy:     privacy,
	}, c.Params.EpisodeIds)
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Playlists: []*model.Playlist{playlist},
	}
}

func deletePlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequirePlaylistId(); c.Err != nil {
		return
	}

	if ok, err := c.App.HasPermissionToPlaylist(c.Params.UserId, c.Params.PlaylistId); err != nil {
		c.Err = err
		return
	} else if !ok {
		c.SetNotPermitted("playlist")
		return
	}

	if err := c.App.DeletePlaylist(c.Params.PlaylistId); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}

func addEpisodeToPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequirePlaylistId().RequireEpisodeId(); c.Err != nil {
		return
	}

	if ok, err := c.App.HasPermissionToPlaylist(c.Params.UserId, c.Params.PlaylistId); err != nil {
		c.Err = err
		return
	} else if !ok {
		c.SetNotPermitted("playlist")
		return
	}

	if err := c.App.AddEpisodeToPlaylist(c.Params.PlaylistId, c.Params.EpisodeId); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}

func removeEpisodeFromPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession().RequireBody(req).RequirePlaylistId().RequireEpisodeId(); c.Err != nil {
		return
	}

	if ok, err := c.App.HasPermissionToPlaylist(c.Params.UserId, c.Params.PlaylistId); err != nil {
		c.Err = err
		return
	} else if !ok {
		c.SetNotPermitted("playlist")
		return
	}

	if err := c.App.RemoveEpisodeFromPlaylist(c.Params.PlaylistId, c.Params.EpisodeId); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
}
