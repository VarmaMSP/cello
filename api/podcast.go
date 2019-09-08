package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.Handler("GET", "/results", api.NewHandler(SearchPodcasts))
	api.router.Handler("GET", "/podcasts/:podcastId", api.NewHandler(GetPodcast))
}

func GetPodcast(c *Context, w http.ResponseWriter) {
	podcastId := c.Param("podcastId")
	podcast, err := c.store.Podcast().GetInfo(podcastId)
	if err != nil {
		c.err = err
		return
	}
	episodes, err := c.store.Episode().GetAllByPodcast(podcastId, 1000, 0)
	if err != nil {
		c.err = err
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"podcast":  podcast,
		"episodes": episodes,
	})
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// body, err1 := c.app.UI.RenderPodcastPage(map[string]interface{}{
	// 	"podcast":  podcast,
	// 	"episodes": episodes,
	// })
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }

	// w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusOK)
	// io.Copy(w, body)
}

func SearchPodcasts(c *Context, w http.ResponseWriter) {
	searchQuery := c.Query("search_query")
	searchResult, err := c.esClient.Search().
		Index("podcast").
		Query(elastic.NewMultiMatchQuery(searchQuery, "title", "author", "description")).
		Size(50).
		Do(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	var podcasts []*model.PodcastInfo
	for _, item := range searchResult.Each(reflect.TypeOf(model.PodcastInfo{})) {
		tmp, _ := item.(model.PodcastInfo)
		tmp.Description = ""
		podcasts = append(podcasts, &tmp)
	}

	res, _ := json.Marshal(map[string]interface{}{
		"totalCount": len(podcasts),
		"results":    podcasts,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
