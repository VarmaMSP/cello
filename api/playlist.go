package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPlaylistHandlers() {
	api.router.Handler("GET", "/playlists", api.NewHandlerSessionRequired(GetUserPlaylists))
	api.router.Handler("GET", "/playlists/:playlistId", api.NewHandler(GetPlaylist))
	api.router.Handler("POST", "/playlists", api.NewHandlerSessionRequired(CreatePlaylist))
	api.router.Handler("POST", "/playlists/:playlistId/episodes/:episodeId", api.NewHandlerSessionRequired(AddEpisodeToPlaylist))
}

func GetUserPlaylists(c *Context, w http.ResponseWriter) {
	req := &GetUserPlaylistsReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_user_playlists_req_load", err.Error(), 400, nil)
		return
	}

	playlists, err := c.app.GetPlaylistsByUser(req.UserId)
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlists": playlists,
	}))
}

func GetPlaylist(c *Context, w http.ResponseWriter) {
	req := &GetPlaylistReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_playlist_req.load", err.Error(), 400, nil)
		return
	}

	playlist, err := c.app.GetPlaylist(req.PlaylistId)
	if err != nil {
		c.err = err
		return
	}

	episodes, err := c.app.GetEpisodesInPlaylist(req.PlaylistId)
	if err != nil {
		c.err = err
		return
	}

	if c.session != nil && c.session.UserId != 0 {
		playbacks, err := c.app.GetUserPlaybacksForEpisodes(c.session.UserId, model.GetEpisodeIds(episodes))
		if err != nil {
			c.err = err
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

func CreatePlaylist(c *Context, w http.ResponseWriter) {
	req := &CreatePlaylistReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.create_playlist_req.load", err.Error(), 400, nil)
		return
	}

	playlist, err := c.app.SavePlaylist(req.Title, req.Privacy, req.UserId)
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", model.UrlParamFromId(playlist.Title, playlist.Id))
	w.WriteHeader(http.StatusCreated)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlist": playlist,
	}))
}

func AddEpisodeToPlaylist(c *Context, w http.ResponseWriter) {
	req := &AddEpisodeToPlaylistReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.add_episode_to_playlist_req.load", err.Error(), 400, nil)
		return
	}

	if _, err := c.app.SaveEpisodeToPlaylist(req.EpisodeId, req.PlaylistId); err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
