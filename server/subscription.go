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

	podcastIds := make([]int64, len(subscriptions))
	for i, podcast := range subscriptions {
		podcastIds[i] = podcast.Id
	}
	episodes, err := c.App.GetEpisodesInPodcastIds(podcastIds, 0, 15)
	if err != nil {
		c.Err = err
		return
	}

	episodeIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		episodeIds[i] = episode.Id
	}
	playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Params.UserId, episodeIds)
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Episodes: episodes,
	}
}

func BrowseSubscriptionsFeed(c *Context, w http.ResponseWriter, req *http.Request) {
	subscriptions, err := c.App.GetUserSubscriptions(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	podcastIds := make([]int64, len(subscriptions))
	for i, podcast := range subscriptions {
		podcastIds[i] = podcast.Id
	}
	episodes, err := c.App.GetEpisodesInPodcastIds(podcastIds, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	episodeIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		episodeIds[i] = episode.Id
	}
	playbacks, err := c.App.GetUserPlaybacksForEpisodes(c.Params.UserId, episodeIds)
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Episodes: episodes,
	}
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

	c.Response.StatusCode = http.StatusOK
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

	c.Response.StatusCode = http.StatusOK
}
