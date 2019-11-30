package api

import (
	"net/http"

	"github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.Handler("GET", "/podcasts/:podcastId", api.NewHandler(GetPodcast))
}

func GetPodcast(c *Context, w http.ResponseWriter) {
	req := &GetPodcastReq{}
	if err := req.Load(c); err != nil {
		c.err = model.NewAppError("api.get_podcast_req.load", err.Error(), 400, nil)
		return
	}

	feed, err := c.app.GetFeed(req.PodcastId)
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

	podcast, err := c.app.GetPodcast(req.PodcastId)
	if err != nil {
		c.err = err
		return
	}

	episodes, err := c.app.GetEpisodesInPodcast(req.PodcastId, "pub_date_desc", 0, 15)
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
