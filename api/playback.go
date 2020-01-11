package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

const (
	ACTION_PLAYBACK_BEGIN    = "playback_begin"
	ACTION_PLAYBACK_PROGRESS = "playback_progress"
)

func ServiceGetPlaybacks(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeIds, err_ := GetIds(c.Body["episode_ids"])
	if err_ != nil {
		c.SetInvalidBodyParam("episode_ids")
		return
	}

	playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Params.UserId, episodeIds)
	if err != nil {
		c.Err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playbacks": playbacks,
	}))
}

func ServicePlaybackSync(c *Context, w http.ResponseWriter, req *http.Request) {
	switch c.Params.Action {
	case ACTION_PLAYBACK_BEGIN:
		c.RequireBody(req)
		if c.Err == nil {
			PlaybackBegin(c, w, req)
		}

	case ACTION_PLAYBACK_PROGRESS:
		c.RequireBody(req)
		if c.Err == nil {
			PlaybackProgress(c, w, req)
		}

	default:
		c.SetInvalidQueryParam("action")
	}
}

func PlaybackBegin(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeId, err_ := GetId(c.Body["episode_id"])
	if err_ != nil {
		c.SetInvalidBodyParam("episode_id")
		return
	}

	err := c.App.SyncPlayback(episodeId, c.Params.UserId, model.PLAYBACK_EVENT_BEGIN, 0)
	if err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PlaybackProgress(c *Context, w http.ResponseWriter, req *http.Request) {
	episodeId, err_ := GetId(c.Body["episode_id"])
	if err_ != nil {
		c.SetInvalidBodyParam("episode_id")
		return
	}

	position, ok := c.Body["position"].(float64)
	if !ok {
		c.SetInvalidBodyParam("position")
		return
	}

	err := c.App.SyncPlayback(episodeId, c.Params.UserId, model.PLAYBACK_EVENT_PLAYING, position)
	if err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
