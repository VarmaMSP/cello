package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/s3"
)

func (api *Api) RegisterHomeHandlers() {
	api.router.Handler("GET", "/home", api.NewHandler(GetHomePageDate))
	api.router.Handler("GET", "/charts/:chartId", api.NewHandler(GetPodcastsInChart))
}

func GetHomePageDate(c *Context, w http.ResponseWriter) {
	recommended, err := c.app.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_DISCOVER, "recommended.json")
	if err != nil {
		c.err = err
		return
	}

	categories, err := c.app.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_DISCOVER, "categories.json")
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(
		model.EncodeToJson(&struct {
			Recommended json.RawMessage `json:"recommended"`
			Categories  json.RawMessage `json:"categories"`
		}{
			Recommended: (json.RawMessage)(recommended),
			Categories:  (json.RawMessage)(categories),
		}),
	)
}

func GetPodcastsInChart(c *Context, w http.ResponseWriter) {
	chartId := c.Param("chartId")
	podcasts, err := c.app.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_CHARTS, fmt.Sprintf("%s.json", chartId))
	if err != nil {
		c.err = err
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(
		model.EncodeToJson(&struct {
			Podcasts json.RawMessage `json:"podcasts"`
		}{
			Podcasts: (json.RawMessage)(podcasts),
		}),
	)
}
