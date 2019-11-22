package api

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterEpisodeHandlers() {
	api.router.Handler("GET", "/podcasts/:podcastId/episodes", api.NewHandler(GetPodcastEpisodes))
	api.router.Handler("GET", "/episodes/:episodeId", api.NewHandler(GetEpisode))
	api.router.Handler("GET", "/subscriptions/feed", api.NewHandlerSessionRequired(GetFeed))
	api.router.Handler("GET", "/history", api.NewHandlerSessionRequired(GetHistory))
	api.router.Handler("PUT", "/playback", api.NewHandlerSessionRequired(GetEpisodePlaybacks))
	api.router.Handler("POST", "/sync/:episodeId", api.NewHandlerSessionRequired(SyncPlayback))
	api.router.Handler("POST", "/sync/:episodeId/progress", api.NewHandlerSessionRequired(SyncPlaybackProgress))
}

func GetEpisode(c *Context, w http.ResponseWriter) {
	req := &GetEpisodeReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	episode, err := c.app.GetEpisode(req.EpisodeId)
	if err != nil {
		c.err = err
		return
	}
	episode.Sanitize()

	podcast, err := c.app.GetPodcast(episode.PodcastId)
	if err != nil {
		c.err = err
		return
	}
	podcast.Sanitize()

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
		c.err = err
		return
	}

	episodes, err := c.app.GetEpisodesInPodcast(req.PodcastId, req.Order, req.Offset, req.Offset)
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

func GetFeed(c *Context, w http.ResponseWriter) {
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

func GetHistory(c *Context, w http.ResponseWriter) {
	playbacks, err := c.app.GetAllEpisodePlaybacksByUser(c.session.UserId)
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
		"history":   episodes,
		"playbacks": playbacks,
	}))
}

func GetEpisodePlaybacks(c *Context, w http.ResponseWriter) {
	req := &GetEpisodePlaybacksReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	playbacks, err := c.app.GetAllEpisodePlaybacks(req.EpisodeIds, req.CurrentUserId)
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
		"playbacks": playbacks,
	}))
}

func SyncPlayback(c *Context, w http.ResponseWriter) {
	req := &SyncPlaybackReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	if err := c.app.SaveEpisodePlayback(req.EpisodeId, req.CurrentUserId); err != nil {
		c.err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SyncPlaybackProgress(c *Context, w http.ResponseWriter) {
	req := &SyncPlaybackProgressReq{}
	if err := req.Load(c); err != nil {
		c.err = err
		return
	}

	c.app.SaveEpisodeProgress(req.EpisodeId, req.CurrentUserId, req.CurrentTime)

	w.WriteHeader(http.StatusOK)
}
