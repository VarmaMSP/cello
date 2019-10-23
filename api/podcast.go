package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.Handler("GET", "/results", api.NewHandler(SearchPodcasts))
	api.router.Handler("GET", "/podcasts/:podcastId", api.NewHandler(GetPodcast))
	api.router.Handler("GET", "/trending", api.NewHandler(GetTrendingPodcasts))
	api.router.Handler("GET", "/trending/:category", api.NewHandler(GetTrendingPodcastsByCategory))
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

	podcast, err := c.app.GetPodcast(podcastId)
	if err != nil {
		c.err = err
		return
	}
	podcast.Sanitize()

	episodes, err := c.app.GetEpisodesInPodcast(podcastId, 1000, 0)
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
		"podcast":  podcast,
		"episodes": episodes,
	}))
}

func GetTrendingPodcasts(c *Context, w http.ResponseWriter) {
	file, err := os.Open("/var/www/static/trending.json")
	if err != nil {
		c.err = model.NewAppError("api.get_trending_podcasts", err.Error(), 400, nil)
		return
	}

	w.Header().Set(headers.CacheControl, "public, max-age=21600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func GetTrendingPodcastsByCategory(c *Context, w http.ResponseWriter) {
	file, err := os.Open("/var/www/static/trending_" + c.Param("category") + ".json")
	if err != nil {
		c.err = model.NewAppError("api.get_trending_podcasts", err.Error(), 400, nil)
		return
	}

	w.Header().Set(headers.CacheControl, "public, max-age=21600")
	w.Header().Set(headers.ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
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
