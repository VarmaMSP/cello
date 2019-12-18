package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func GetSubscriptionsPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	subscriptions, err := c.App.GetUserSubscriptions(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	episodes, err := c.App.GetEpisodesInPodcastIds(model.GetPodcastIds(subscriptions), 0, 15)
	if err != nil {
		c.Err = err
		return
	}

	playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Params.UserId, model.GetEpisodeIds(episodes))
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
	}))
}

func BrowseSubscriptionsFeed(c *Context, w http.ResponseWriter, req *http.Request) {
	subscriptions, err := c.App.GetUserSubscriptions(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	episodes, err := c.App.GetEpisodesInPodcastIds(model.GetPodcastIds(subscriptions), c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Params.UserId, model.GetEpisodeIds(episodes))
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
	}))
}

func ServiceSubscribePodcast(c *Context, w http.ResponseWriter, req *http.Request) {
	podcastHashId, ok := c.Body["podcast_id"].(string)
	if !ok {
		c.SetInvalidBodyParam("podcast_id")
		return
	}
	podcastId, err := model.Int64FromHashId(podcastHashId)
	if !ok {
		c.SetError(err)
		return
	}

	if err := c.App.SaveSubscription(c.Params.UserId, podcastId); err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ServiceUnsubscribePodcast(c *Context, w http.ResponseWriter, req *http.Request) {
	podcastHashId, ok := c.Body["podcast_id"].(string)
	if !ok {
		c.SetInvalidBodyParam("podcast_id")
		return
	}
	podcastId, err := model.Int64FromHashId(podcastHashId)
	if !ok {
		c.SetError(err)
		return
	}

	if err := c.App.DeleteSubscription(c.Params.UserId, podcastId); err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
