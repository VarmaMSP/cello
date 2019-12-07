package api

import (
	"encoding/json"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/s3"
)

func (api *Api) RegisterDiscoverHandlers() {
	api.router.Handler("GET", "/discover", api.NewHandler(GetHomePageDate))
}

func GetHomePageDate(c *Context, w http.ResponseWriter) {
	trending, err := c.app.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_DISCOVER, "discover.json")
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
			Trending   json.RawMessage `json:"trending"`
			Categories json.RawMessage `json:"categories"`
		}{
			Trending:   (json.RawMessage)(trending),
			Categories: (json.RawMessage)(categories),
		}),
	)
}
