package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *Api) RegisterPodcastHandlers() {
	api.router.HandlerFunc("GET", "/podcast/:podcastId", api.GetPodcast)
}

func (api *Api) GetPodcast(w http.ResponseWriter, req *http.Request) {
	podcastId := httprouter.ParamsFromContext(req.Context()).ByName("podcastId")

	podcast, err := api.app.GetPodcast(podcastId)
	if err != nil {
		fmt.Print(err)
	}

	episodes, err := api.app.GetEpisodes(podcastId, 200, 0)
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
