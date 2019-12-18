package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

const (
	ACTION_ADD_EPISODE    = "add_episode"
	ACTION_DELETE_EPISODE = "delete_episode"
)

func GetPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequirePlaylistId()
	if c.Err != nil {
		return
	}

	playlist, err := c.App.GetPlaylist(c.Params.PlaylistId)
	if err != nil {
		c.Err = err
		return
	}

	episodes, err := c.App.GetEpisodesInPlaylist(c.Params.PlaylistId, 0, 1000)
	if err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Session.UserId, model.GetEpisodeIds(episodes))
		if err != nil {
			c.Err = err
			return
		}
		model.EpisodesJoinPlaybacks(episodes, playbacks)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlist": playlist,
		"episodes": episodes,
	}))
}

func GetUserPlaylists(c *Context, w http.ResponseWriter, req *http.Request) {
	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	episodesByPlaylist := map[string]([]*model.Episode){}
	for _, playlist := range playlists {
		e, err := c.App.GetEpisodesInPlaylist(playlist.Id, 0, 3)
		if err != nil {
			c.Err = err
			return
		}
		episodesByPlaylist[model.HashIdFromInt64(playlist.Id)] = e
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlists":            playlists,
		"episodes_by_playlist": episodesByPlaylist,
	}))
}

func ServiceAddToPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId, 0, 1000)
	if err != nil {
		c.Err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlists": playlists,
	}))
}

func ServiceCreatePlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	title, ok := c.Body["title"].(string)
	if !ok {
		c.SetInvalidBodyParam("title")
	}

	privacy, ok := c.Body["privacy"].(string)
	if !ok {
		c.SetInvalidBodyParam("privacy")
	}

	playlist, err := c.App.SavePlaylist(title, privacy, c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", model.UrlParamFromId(playlist.Title, playlist.Id))
	w.WriteHeader(http.StatusCreated)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlist": playlist,
	}))
}

func ServiceEditPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	switch c.Params.Action {
	case ACTION_ADD_EPISODE:
		c.RequireBody(req)
		if c.Err != nil {
			AddEpisodeToPlaylist(c, w, req)
		}

	case ACTION_DELETE_EPISODE:
		c.RequireBody(req)
		if c.Err != nil {
			DeleteEpisodeFromPlaylist(c, w, req)
		}

	default:
		c.SetInvalidQueryParam("action")
	}
}

func AddEpisodeToPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeId, err := GetId(c.Body["episode_id"])
	if err != nil {
		c.SetInvalidBodyParam("episode_id")
		return
	}

	playlistId, err := GetId(c.Body["playlist_id"])
	if err != nil {
		c.SetInvalidBodyParam("playlist_id")
		return
	}

	if _, err := c.App.SaveEpisodeToPlaylist(episodeId, playlistId); err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteEpisodeFromPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeId, err := GetId(c.Body["episode_id"])
	if err != nil {
		c.SetInvalidBodyParam("episode_id")
		return
	}

	playlistId, err := GetId(c.Body["playlist_id"])
	if err != nil {
		c.SetInvalidBodyParam("playlist_id")
		return
	}

	if _, err := c.App.SaveEpisodeToPlaylist(episodeId, playlistId); err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}
