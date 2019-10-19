package api

import (
	"net/http"
	"time"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterEpisodeHandlers() {
	api.router.Handler("GET", "/feed", api.NewHandlerSessionRequired(GetFeed))
	api.router.Handler("PUT", "/playback", api.NewHandlerSessionRequired(GetEpisodePlaybacks))
	api.router.Handler("POST", "/sync/:episodeId", api.NewHandlerSessionRequired(SyncPlayback))
	api.router.Handler("POST", "/sync/:episodeId/progress", api.NewHandlerSessionRequired(SyncPlaybackProgress))
}

func GetFeed(c *Context, w http.ResponseWriter) {
	subscriptions, err := c.app.GetUserSubscriptions(c.session.UserId)
	if err != nil {
		c.err = err
		return
	}

	now := time.Now().UTC()
	weekAgo := now.AddDate(0, 0, -7)
	podcastIds := make([]string, len(subscriptions))
	for i, podcast := range subscriptions {
		podcastIds[i] = podcast.Id
	}
	episodes, err := c.app.GetAllEpisodesPubblishedBetween(&weekAgo, &now, podcastIds)
	if err != nil {
		c.err = err
		return
	}

	episodeIds := make([]string, len(episodes))
	for i, episode := range episodes {
		episodeIds[i] = episode.Id
	}
	playbacks, err := c.app.GetAllEpisodePlayabacks(episodeIds, c.session.UserId)
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes":  episodes,
		"playbacks": playbacks,
	}))
}

type GetEpisodePlaybacksReq struct {
	EpisodeIds []string `json:"episode_ids"`
}

func GetEpisodePlaybacks(c *Context, w http.ResponseWriter) {
	var body GetEpisodePlaybacksReq
	if err := c.DecodeBody(&body); err != nil {
		return
	}

	playbacks, err := c.app.GetAllEpisodePlayabacks(body.EpisodeIds, c.session.UserId)
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

func SyncPlayback(c *Context, w http.ResponseWriter) {
	if err := c.app.SaveEpisodePlayback(c.Param("episodeId"), c.session.UserId); err != nil {
		c.err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SyncPlaybackProgress(c *Context, w http.ResponseWriter) {
	c.app.SaveEpisodeProgress(
		c.Param("episodeId"),
		c.session.UserId,
		model.IntFromStr(c.Body()["current_time"]),
	)

	w.WriteHeader(http.StatusOK)
}
