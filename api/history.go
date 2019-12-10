package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterHistoryHandlers() {
	api.router.Handler("GET", "/history/feed", api.NewHandlerSessionRequired(GetHistoryFeed))
}

func GetHistoryFeed(c *Context, w http.ResponseWriter) {
	req := &GetHistoryFeedReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.subscribe_to_podcast_req.load", err.Error(), 400, nil)
		return
	}

	playbacks, err := c.app.GetUserPlaybacks(req.UserId, req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	episodeIds := make([]int64, len(playbacks))
	for i, playback := range playbacks {
		episodeIds[i] = playback.EpisodeId
	}
	episodes, err := c.app.GetEpisodesByIds(episodeIds)
	if err != nil {
		c.err = err
		return
	}
	model.EpisodesJoinPlaybacks(episodes, playbacks)

	podcastIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		podcastIds[i] = episode.PodcastId
	}
	podcasts, err := c.app.GetPodcastsByIs(model.RemoveDuplicatesInt64(podcastIds))
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
		"podcasts": podcasts,
	}))
}
