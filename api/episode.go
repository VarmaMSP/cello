package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterEpisodeHandlers() {
	api.router.Handler("GET", "/episodes/:episodeId", api.NewHandler(GetEpisode))
	api.router.Handler("GET", "/podcasts/:podcastId/episodes", api.NewHandler(GetPodcastEpisodes))
}

func GetEpisode(c *Context, w http.ResponseWriter) {
	req := &GetEpisodeReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_episode_req.load", err.Error(), 400, nil)
		return
	}

	episode, err := c.app.GetEpisode(req.EpisodeId)
	if err != nil {
		c.err = err
		return
	}

	podcast, err := c.app.GetPodcast(episode.PodcastId)
	if err != nil {
		c.err = err
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episode": episode,
		"podcast": podcast,
	}))
}

func GetPodcastEpisodes(c *Context, w http.ResponseWriter) {
	req := &GetPodcastEpisodesReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_podcast_episodes_req.load", err.Error(), 400, nil)
		return
	}

	episodes, err := c.app.GetEpisodesInPodcast(req.PodcastId, req.Order, req.Offset, req.Limit)
	if err != nil {
		c.err = err
		return
	}

	if c.session != nil && c.session.UserId != 0 {
		playbacks, err := c.app.GetUserPlaybacksForEpisodes(c.session.UserId, model.GetEpisodeIds(episodes))
		if err != nil {
			c.err = err
			return
		}
		model.EpisodesJoinPlaybacks(episodes, playbacks)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"episodes": episodes,
	}))
}
