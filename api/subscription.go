package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterSubscriptionHandlers() {
	api.router.Handler("GET", "/subscriptions/feed", api.NewHandlerSessionRequired(GetSubscriptionsFeed))
	api.router.Handler("PUT", "/podcasts/:podcastId/subscribe", api.NewHandlerSessionRequired(SubscribeToPodcast))
	api.router.Handler("PUT", "/podcasts/:podcastId/unsubscribe", api.NewHandlerSessionRequired(UnsubscribeToPodcast))
}

func GetSubscriptionsFeed(c *Context, w http.ResponseWriter) {
	req := &GetSubscriptionFeedReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_feed_req.load", err.Error(), 400, nil)
		return
	}

	subscriptions, err := c.app.GetUserSubscriptions(req.UserId)
	if err != nil {
		c.err = err
		return
	}

	episodes, err := c.app.GetEpisodesInPodcastIds(model.GetPodcastIds(subscriptions), req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	playbacks, err := c.app.GetUserPlaybacksForEpisodes(req.UserId, model.GetEpisodeIds(episodes))
	if err != nil {
		c.err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
	}))
}

func SubscribeToPodcast(c *Context, w http.ResponseWriter) {
	req := &SubscribeToPodcastReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.subscribe_to_podcast_req.load", err.Error(), 400, nil)
		return
	}
	if err := c.app.SaveSubscription(req.UserId, req.PodcastId); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnsubscribeToPodcast(c *Context, w http.ResponseWriter) {
	req := &UnsubscribeToPodcastReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.unsubscribe_to_podcast_req.load", err.Error(), 400, nil)
		return
	}

	if err := c.app.DeleteSubscription(req.UserId, req.PodcastId); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
