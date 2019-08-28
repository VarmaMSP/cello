package api

import (
	"encoding/json"
	"net/http"
)

func (api *Api) RegisterCurationHandlers() {
	api.router.HandlerFunc("GET", "/curations", api.GetAllCurations)
	api.router.HandlerFunc("POST", "/curations", api.CreateCuration)
}

type CreateCurationReq struct {
	Title string `json:"title"`
}

func (api *Api) CreateCuration(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data CreateCurationReq
	if err := decoder.Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := api.app.CreatePodcastCuration(data.Title); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *Api) GetAllCurations(w http.ResponseWriter, req *http.Request) {
	curations, err := api.app.GetAllPodcastCurations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err1 := json.Marshal(map[string]interface{}{
		"curations": curations,
	})
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
