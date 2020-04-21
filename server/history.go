package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func GetHistoryPageData(c *Context, w http.ResponseWriter, req *http.Request) {
	playbacks, err := c.App.GetUserPlaybacks(c.Params.UserId, 0, 15)
	if err != nil {
		c.Err = err
		return
	}

	episodeIds := make([]int64, len(playbacks))
	for i, playback := range playbacks {
		episodeIds[i] = playback.EpisodeId
	}
	episodes, err := c.App.GetEpisodesByIds(episodeIds)
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	podcastIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		podcastIds[i] = episode.PodcastId
	}
	podcasts, err := c.App.GetPodcastsByIds(model.RemoveDuplicatesInt64(podcastIds))
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Podcasts: podcasts,
		Episodes: episodes,
	}
}

func BrowseHistoryFeed(c *Context, w http.ResponseWriter, req *http.Request) {
	playbacks, err := c.App.GetUserPlaybacks(c.Params.UserId, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	episodeIds := make([]int64, len(playbacks))
	for i, playback := range playbacks {
		episodeIds[i] = playback.EpisodeId
	}
	episodes, err := c.App.GetEpisodesByIds(episodeIds)
	if err != nil {
		c.Err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	podcastIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		podcastIds[i] = episode.PodcastId
	}
	podcasts, err := c.App.GetPodcastsByIds(model.RemoveDuplicatesInt64(podcastIds))
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Podcasts: podcasts,
		Episodes: episodes,
	}
}
