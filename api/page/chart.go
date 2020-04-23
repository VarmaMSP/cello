package page

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/web"
)

func Chart(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireChartId(); c.Err != nil {
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

	podcasts, err_ := c.App.FileStorage.ReadFile(filestorage.BUCKET_PHENOPOD_CHARTS, fmt.Sprintf("%s.json", chartName))
	if err_ != nil {
		c.SetError(err)
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
