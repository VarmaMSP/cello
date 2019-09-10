package api

import (
	"encoding/json"
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (api *Api) RegisterCurationRoutes() {
	api.router.Handler("GET", "/curations", api.NewHandler(GetCurationsWithPodcasts))
	api.router.Handler("POST", "/curations", api.NewHandler(AddCuration))
	api.router.Handler("DELETE", "/curations/:curationId", api.NewHandler(DeleteCuration))
	api.router.Handler("GET", "/curations/:curationId/podcasts", api.NewHandler(GetPodcastsByCuration))
	api.router.Handler("POST", "/curations/:curationId/podcasts/:podcastId", api.NewHandler(AddPodcastToCuration))
	api.router.Handler("DELETE", "/curations/:curationId/podcasts/:podcastId", api.NewHandler(DeletePodcastFromCuration))
}

func AddCuration(c *Context, w http.ResponseWriter) {
	curationTitle := c.Body()["curation_title"]

	m := &model.Curation{Title: curationTitle}
	if err := c.store.Curation().Save(m); err != nil {
		c.err = err
		return
	}

	w.Header().Set("Location", "/curations/"+m.Id)
	w.WriteHeader(http.StatusCreated)
}

func GetCurationsWithPodcasts(c *Context, w http.ResponseWriter) {
	curations, err := c.store.Curation().GetAll()
	if err != nil {
		c.err = err
		return
	}

	var res []map[string]interface{}
	for _, curation := range curations {
		podcasts, err := c.store.Curation().GetPodcastsByCuration(curation.Id, 0, 10)
		if err != nil {
			continue
		}
		curation.CreatedAt = 0
		for _, podcast := range podcasts {
			podcast.Description = ""
		}
		res = append(res, map[string]interface{}{
			"curation": curation,
			"podcasts": podcasts,
		})
	}

	r, _ := json.Marshal(map[string]interface{}{
		"totalCount": len(res),
		"results":    res,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(r)
}

func DeleteCuration(c *Context, w http.ResponseWriter) {
	curationId := c.Param("curationId")

	if err := c.store.Curation().Delete(curationId); err != nil {
		c.err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetPodcastsByCuration(c *Context, w http.ResponseWriter) {
	curationId := c.Param("curationId")

	curation, err := c.store.Curation().Get(curationId)
	if err != nil {
		c.err = err
		return
	}
	podcasts, err := c.store.Curation().GetPodcastsByCuration(curationId, 0, 100)
	if err != nil {
		c.err = err
		return
	}

	curation.CreatedAt = 0
	for _, podcast := range podcasts {
		podcast.Description = ""
	}

	r, _ := json.Marshal(map[string]interface{}{
		"curation": curation,
		"podcasts": podcasts,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(r)
}

func AddPodcastToCuration(c *Context, w http.ResponseWriter) {
	podcastId, curationId := c.Param("podcastId"), c.Param("curationId")

	m := &model.PodcastCuration{PodcastId: podcastId, CurationId: curationId}
	if err := c.store.Curation().SavePodcastCuration(m); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeletePodcastFromCuration(c *Context, w http.ResponseWriter) {

}
