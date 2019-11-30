package api

import (
	"github.com/varmamsp/cello/model"
)

type GetPodcastReq struct {
	PodcastId int64 `validate:"required"`
}

func (o *GetPodcastReq) Load(c *Context) (err error) {
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}

	err = c.app.Validate.Struct(o)
	return
}
