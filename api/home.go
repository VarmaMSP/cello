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

	categories, err := c.App.GetAllCategories()
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Categories: categories,
	}
	c.Response.Raw = model.EncodeToJson(&struct {
		Recommended json.RawMessage `json:"recommended"`
	}{
		Recommended: (json.RawMessage)(recommended),
	})
}

func GetChart(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireChartId()
	if c.Err != nil {
		return
	}

	chartName, chartId, err_ := model.ParseCategoryUrlParam(c.Params.ChartId)
	if err_ != nil {
		return
	}

	category, err := c.App.GetCategory(chartId)
	if err != nil {
		c.Err = err
		return
	}

	podcasts, err := c.App.GetStaticFile(s3.BUCKET_NAME_PHENOPOD_CHARTS, fmt.Sprintf("%s.json", chartName))
	if err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK
	c.Response.Data = &model.ApiResponseData{
		Categories: []*model.Category{category},
	}
	c.Response.Raw = model.EncodeToJson(&struct {
		Podcasts json.RawMessage `json:"podcasts"`
	}{
		Podcasts: (json.RawMessage)(podcasts),
	})
}
