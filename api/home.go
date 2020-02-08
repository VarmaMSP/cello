package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/s3"
)

func GetHomePageData(c *Context, w http.ResponseWriter, req *http.Request) {
	recommended, err := c.App.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_DISCOVER, "recommended.json")
	if err != nil {
		c.Err = err
		return
	}

	categories, err := c.App.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_DISCOVER, "categories.json")
	if err != nil {
		c.Err = err
		return
	}

	x, err := c.App.GetAllCategories()
	if err != nil {
		c.Err = err
		return
	}

	c.Response.Data = &model.ApiResponseData{
		Categories: x,
	}

	c.Response.Raw = model.EncodeToJson(&struct {
		Recommended json.RawMessage `json:"recommended"`
		Categories  json.RawMessage `json:"categories"`
	}{
		Recommended: (json.RawMessage)(recommended),
		Categories:  (json.RawMessage)(categories),
	})
}

func GetChart(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireChartId()
	if c.Err != nil {
		return
	}

	podcasts, err := c.App.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_CHARTS, fmt.Sprintf("%s.json", c.Params.ChartId))
	if err != nil {
		c.Err = err
		return
	}

	c.Response.Raw = model.EncodeToJson(&struct {
		Podcasts json.RawMessage `json:"podcasts"`
	}{
		Podcasts: (json.RawMessage)(podcasts),
	})
}
