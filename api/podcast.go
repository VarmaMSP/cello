package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.HandlerFunc("GET", "/podcasts/:podcastId", api.GetPodcast)
	api.router.HandlerFunc("GET", "/fuck", api.ProxyTest)
}

func (api *Api) GetPodcast(w http.ResponseWriter, req *http.Request) {
	podcastId := httprouter.ParamsFromContext(req.Context()).ByName("podcastId")

	podcast, err := api.app.GetPodcast(podcastId)
	if err != nil {
		fmt.Print(err)
	}

	episodes, err := api.app.GetEpisodes(podcastId, 1000, 0)
	if err != nil {
		fmt.Print(err)
	}
	for _, episode := range episodes {
		episode.Description = ""
	}

	res, err1 := json.Marshal(map[string]interface{}{
		"podcast":  podcast,
		"episodes": episodes,
	})
	if err != nil {
		fmt.Print(err1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (api *Api) ProxyTest(w http.ResponseWriter, req *http.Request) {
	x, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(w, x.Body)
}
