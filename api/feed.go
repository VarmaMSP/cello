package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterFeedHandlers() {
	api.router.Handler("GET", "/subscriptions/feed", api.NewHandlerSessionRequired(GetSubscriptionsFeed))
	api.router.Handler("GET", "/history/feed", api.NewHandlerSessionRequired(GetHistoryFeed))
}

func GetSubscriptionsFeed(c *Context, w http.ResponseWriter) {
	req := &GetFeedReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	subscriptions, err := c.app.GetUserSubscriptions(req.CurrentUserId)
	if err != nil {
		c.err = err
		return
	}

	podcastIds := make([]string, len(subscriptions))
	for i, podcast := range subscriptions {
		podcastIds[i] = podcast.Id
	}
	episodes, err := c.app.GetAllEpisodesPublishedBefore(podcastIds, req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	episodeIds := make([]string, len(episodes))
	for i, episode := range episodes {
		episode.Sanitize()
		episodeIds[i] = episode.Id
	}
	playbacks, err := c.app.GetAllEpisodePlaybacks(episodeIds, c.session.UserId)
	if err != nil {
		c.err = err
		return
	}
	for _, playback := range playbacks {
		playback.Sanitize()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes":  episodes,
		"playbacks": playbacks,
	}))
}

func GetHistoryFeed(c *Context, w http.ResponseWriter) {
	req := &GetFeedReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	playbacks, err := c.app.GetAllEpisodePlaybacksByUser(req.CurrentUserId, req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	episodeIds := make([]string, len(playbacks))
	for i, playback := range playbacks {
		playback.Sanitize()
		episodeIds[i] = playback.EpisodeId
	}
	episodes, err := c.app.GetEpisodesByIds(episodeIds)
	if err != nil {
		c.err = err
		return
	}
	for _, episode := range episodes {
		episode.Sanitize()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes":  episodes,
		"playbacks": playbacks,
	}))
}
