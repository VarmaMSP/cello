package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPlaybackHandlers() {
	api.router.Handler("GET", "/playback", api.NewHandlerSessionRequired(GetPlaybacks))
	api.router.Handler("POST", "/playback/:episodeId", api.NewHandlerSessionRequired(StartPlayback))
	api.router.Handler("POST", "/playback/:episodeId/sync", api.NewHandlerSessionRequired(SyncPlayback))
}

func GetPlaybacks(c *Context, w http.ResponseWriter) {
	req := &GetPlaybacksReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.start_playback_req.load", err.Error(), 400, nil)
		return
	}

	playbacks, err := c.app.GetUserPlaybacksForEpisodes(req.UserId, req.EpisodeIds)
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"playbacks": playbacks,
	}))
}

func StartPlayback(c *Context, w http.ResponseWriter) {
	req := &StartPlaybackReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.start_playback_req.load", err.Error(), 400, nil)
		return
	}

	if err := c.app.SyncPlayback(req.EpisodeId, req.UserId, model.PLAYBACK_EVENT_COMPLETE, 0); err != nil {
		c.err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SyncPlayback(c *Context, w http.ResponseWriter) {
	req := &SyncPlaybackReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.sync_playback_req.load", err.Error(), 400, nil)
		return
	}

	if err := c.app.SyncPlayback(req.EpisodeId, req.UserId, model.PLAYBACK_EVENT_COMPLETE, req.Progress); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
