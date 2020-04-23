package page

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/web"
)

func Home(c *web.Context, w http.ResponseWriter, req *http.Request) {
	recommended, err_ := c.App.FileStorage.ReadFile(filestorage.BUCKET_PHENOPOD_DISCOVER, "home.json")
	if err_ != nil {
		c.SetError(err_)
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
	c.Response.Raw = recommended
}
