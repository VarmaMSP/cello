package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPlaylistHandlers() {
	api.router.Handler("GET", "/playlists/:playlistId", api.NewHandler(GetPlaylist))
	api.router.Handler("GET", "/playlists", api.NewHandlerSessionRequired(GetCurrentUserPlaylists))
	api.router.Handler("POST", "/playlists", api.NewHandlerSessionRequired(CreatePlaylist))
	api.router.Handler("POST", "/playlists/:playlistId/episodes/:episodeId", api.NewHandlerSessionRequired(AddEpisodeToPlaylist))
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

	episodes, err := c.app.GetEpisodesInPlaylist((playlist.Id))
	if err != nil {
		c.err = err
		return
	}

	episodeIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		episode.Sanitize()
		episodeIds[i] = episode.Id
	}
	var playbacks []*model.EpisodePlayback
	if c.session != nil {
		playbacks, err := c.app.GetAllEpisodePlaybacks(episodeIds, c.session.UserId)
		if err != nil {
			c.err = err
			return
		}
		for _, playback := range playbacks {
			playback.Sanitize()
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlist":  playlist,
		"episodes":  episodes,
		"playbacks": playbacks,
	}))
}

func GetCurrentUserPlaylists(c *Context, w http.ResponseWriter) {
	playlists, err := c.app.GetPlaylistsByUser(c.session.UserId)
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

func CreatePlaylist(c *Context, w http.ResponseWriter) {
	req := &CreatePlaylistReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.create_playlist_req.load", err.Error(), 400, nil)
		return
	}

	playlist, err := c.app.CreatePlaylist(req.Title, req.Privacy, req.CurrentUserId)
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

	if _, err := c.app.AddEpsiodeToPlaylist(req.EpisodeId, req.PlaylistId); err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
