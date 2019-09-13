package api

import (
	"encoding/json"
	"net/http"
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

	if err := c.app.SaveCuration(curationTitle); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCurationsWithPodcasts(c *Context, w http.ResponseWriter) {
	curations, err := c.app.GetCurations()
	if err != nil {
		c.err = err
		return
	}

	var res []map[string]interface{}
	for _, curation := range curations {
		podcasts, err := c.app.GetPodcastsInCuration(curation.Id)
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

	if err := c.app.DeleteCuration(curationId); err != nil {
		c.err = err
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetPodcastsByCuration(c *Context, w http.ResponseWriter) {
	curationId := c.Param("curationId")

	curation, err := c.app.GetCuration(curationId)
	if err != nil {
		c.err = err
		return
	}
	podcasts, err := c.app.GetPodcastsInCuration(curationId)
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

	if err := c.app.SavePodcastToCuration(curationId, podcastId); err != nil {
		c.err = err
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeletePodcastFromCuration(c *Context, w http.ResponseWriter) {

}
