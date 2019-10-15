package api

import (
	"net/http"
	"time"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterEpisodeHandlers() {
	api.router.Handler("GET", "/feed", api.NewHandlerSessionRequired(WeeksFeed))
	api.router.Handler("POST", "/sync/:episodeId", api.NewHandler(SyncPlayback))
	api.router.Handler("POST", "/sync/:episodeId/progress", api.NewHandler(SyncPlaybackProgress))
}

func WeeksFeed(c *Context, w http.ResponseWriter) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
	}))
}

func SyncPlayback(c *Context, w http.ResponseWriter) {
	if c.session != nil {
		err := c.app.SaveEpisodePlayback(c.Param("episodeId"), c.session.UserId)
		if err != nil {
			c.err = err
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func SyncPlaybackProgress(c *Context, w http.ResponseWriter) {
	if c.session != nil {
		c.app.SaveEpisodeProgress(
			c.Param("episodeId"),
			c.session.UserId,
			model.IntFromStr(c.Body()["current_time"]),
		)
	}

	w.WriteHeader(http.StatusOK)
}
