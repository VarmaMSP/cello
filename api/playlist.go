package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

const (
	ACTION_ADD_EPISODE    = "add_episode"
	ACTION_REMOVE_EPISODE = "remove_episode"
)

func GetPlaylistsPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId)
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

func GetPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequirePlaylistId()
	if c.Err != nil {
		return
	}

	playlist, err := c.App.GetPlaylist(c.Params.PlaylistId, true)
	if err != nil {
		c.Err = err
		return
	}

	episodeIds := make([]int64, len(playlist.Members))
	for i, member := range playlist.Members {
		episodeIds[i] = member.EpisodeId
	}
	episodes, err := c.App.GetEpisodesByIds(episodeIds)
	if err != nil {
		c.Err = err
		return
	}

	podcastIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		podcastIds[i] = episode.PodcastId
	}
	podcasts, err := c.App.GetPodcastsByIds(podcastIds)
	if err != nil {
		c.Err = err
		return
	}

	if c.Session != nil && c.Session.UserId != 0 {
		playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Session.UserId, episodeIds)
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
		"podcasts": podcasts,
	}))
}

func ServiceAddToPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeIds, err := GetIds(c.Body["episode_ids"])
	if err != nil {
		c.Err = err
		return
	}

	playlists, err := c.App.GetPlaylistsByUser(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	if err := c.App.JoinPlaylistsToEpisodes(playlists, episodeIds); err != nil {
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

	description, ok := c.Body["description"].(string)
	if !ok {
		c.SetInvalidBodyParam("description")
	}

	episodeIds, err := GetIds(c.Body["episode_ids"])
	if err != nil {
		c.SetInvalidBodyParam("episode_ids")
	}

	playlist, err := c.App.CreatePlaylist(title, privacy, description, c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	for i, episodeId := range episodeIds {
		if err := c.App.AddEpisodeToPlaylist(playlist.Id, episodeId); err != nil {
			c.Err = err
			return
		}
		playlist.Members = append(playlist.Members, &model.PlaylistMember{EpisodeId: episodeId, Position: i + 1})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playlist": playlist,
	}))
}

func ServiceEditPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	switch c.Params.Action {
	case ACTION_ADD_EPISODE:
		c.RequireBody(req)
		if c.Err == nil {
			AddEpisodeToPlaylist(c, w, req)
		}

	case ACTION_REMOVE_EPISODE:
		c.RequireBody(req)
		if c.Err == nil {
			RemoveEpisodeFromPlaylist(c, w, req)
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

	if err := c.App.AddEpisodeToPlaylist(playlistId, episodeId); err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveEpisodeFromPlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
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

	if err := c.App.RemoveEpisodeFromPlaylist(playlistId, episodeId); err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ServiceDeletePlaylist(c *Context, w http.ResponseWriter, req *http.Request) {
	playlistId, err := GetId(c.Body["playlist_id"])
	if err != nil {
		c.SetInvalidBodyParam("playlist_id")
		return
	}

	if err := c.App.DeletePlaylist(playlistId); err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}
