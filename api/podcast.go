package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.Handler("GET", "/podcasts/:podcastId", api.NewHandler(GetPodcast))
}

func GetPodcast(c *Context, w http.ResponseWriter, req *http.Request) {
	podcastId := httprouter.ParamsFromContext(req.Context()).ByName("podcastId")
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
