package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
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

	w.Header().Set(headers.CacheControl, "private, max-age=3600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPodcast(c *Context, w http.ResponseWriter) {
	podcastId := c.Param("podcastId")

	feed, err := c.app.GetFeed(podcastId)
	if err != nil {
		c.err = err
		return
	}
	w.Header().Set(headers.CacheControl, "private, max-age=300, must-revalidate")
	if feed.ETag != "" {
		ifNoneMatch := c.req.Header.Get(headers.IfNoneMatch)
		w.Header().Set(headers.ETag, feed.ETag)
		if ifNoneMatch != "" && ifNoneMatch == feed.ETag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	if feed.LastModified != "" {
		ifModifiedSince := c.req.Header.Get(headers.IfModifiedSince)
		w.Header().Set(headers.LastModified, feed.LastModified)
		if ifModifiedSince != "" && ifModifiedSince == feed.LastModified {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(model.EncodeToJson(map[string]interface{}{
		"podcast":  podcast,
		"episodes": episodes,
	}))
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
