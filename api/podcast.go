package api

import (
	"encoding/json"
	"net/http"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.Handler("GET", "/results", api.NewHandler(SearchPodcasts))
	api.router.Handler("GET", "/podcasts/:podcastId", api.NewHandler(GetPodcast))
	api.router.Handler("PUT", "/podcasts/:podcastId/subscribe", api.NewHandlerSessionRequired(SubscribeToPodcast))
	api.router.Handler("PUT", "/podcasts/:podcastId/unsubscribe", api.NewHandlerSessionRequired(UnsubscribeToPodcast))
}

func SearchPodcasts(c *Context, w http.ResponseWriter) {
	searchQuery := c.Query("search_query")
	podcasts, err := c.app.SearchPodcasts(searchQuery)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"totalCount": len(podcasts),
		"results":    podcasts,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPodcast(c *Context, w http.ResponseWriter) {
	podcastId := c.Param("podcastId")
	podcast, err := c.app.GetPodcastInfo(podcastId)
	if err != nil {
		c.err = err
		return
	}
	episodes, err := c.app.GetEpisodes(podcastId, 1000, 0)
	if err != nil {
		c.err = err
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"podcast":  podcast,
		"episodes": episodes,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SubscribeToPodcast(c *Context, w http.ResponseWriter) {
	userId := c.session.UserId
	podcastId := c.Param("podcastId")

	if err := c.app.CreateSubscription(userId, podcastId); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnsubscribeToPodcast(c *Context, w http.ResponseWriter) {
	userId := c.session.UserId
	podcastId := c.Param("podcastId")

	if err := c.app.DeleteSubscription(userId, podcastId); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
