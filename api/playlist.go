package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPlaylistHandlers() {

	api.router.Handler("GET", "/playlists", api.NewHandlerSessionRequired(GetPlaylists))
	api.router.Handler("GET", "/playlists/:playlistId", api.NewHandler(GetPlaylist))
	api.router.Handler("POST", "/playlists", api.NewHandlerSessionRequired(CreatePlaylist))
	api.router.Handler("POST", "/playlists/episodes", api.NewHandlerSessionRequired(AddEpisodeToPlaylists))
}

func GetPlaylists(c *Context, w http.ResponseWriter) {
	req := &GetPlaylistsReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_playlists_req_load", err.Error(), 400, nil)
		return
	}

	if req.FullDetails {
		GetUserPlaylists(req, c, w)
	} else {
		GetSignedInUserPlaylists(req, c, w)
	}
}

func GetSignedInUserPlaylists(req *GetPlaylistsReq, c *Context, w http.ResponseWriter) {
	playlists, err := c.app.GetPlaylistsByUser(req.UserId, 0, 1000)
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

func GetUserPlaylists(req *GetPlaylistsReq, c *Context, w http.ResponseWriter) {
	playlists, err := c.app.GetPlaylistsByUser(req.UserId, req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	episodesByPlaylist := map[string]([]*model.Episode){}
	for _, playlist := range playlists {
		e, err := c.app.GetEpisodesInPlaylist(playlist.Id, 0, 3)
		if err != nil {
			c.err = err
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

	episodes, err := c.app.GetEpisodesInPlaylist(req.PlaylistId, 0, 1000)
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

func AddEpisodeToPlaylists(c *Context, w http.ResponseWriter) {
	req := &AddEpisodeToPlaylistsReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.add_episode_to_playlist_req.load", err.Error(), 400, nil)
		return
	}

	for _, playlistId := range req.PlaylistIds {
		if _, err := c.app.SaveEpisodeToPlaylist(req.EpisodeId, playlistId); err != nil {
			c.err = err
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
