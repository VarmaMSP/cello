package api

import (
	"github.com/varmamsp/cello/model"
)

type GetEpisodeReq struct {
	EpisodeId int64 `validate:"required"`
}

func (o *GetEpisodeReq) Load(c *Context) (err error) {
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("epsiodeId")); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type GetPodcastEpisodesReq struct {
	PodcastId int64  `validate:"required"`
	Limit     int    `validate:"min=5"`
	Offset    int    `validate:"min=0"`
	Order     string `validate:"required,oneof=pub_date_desc pub_date_asc"`
}

func (o *GetPodcastEpisodesReq) Load(c *Context) (err error) {
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}
	if o.Limit = model.IntFromStr(c.Query("limit")); o.Limit == 0 {
		o.Limit = 10
	}
	if o.Order = c.Query("order"); o.Order == "" {
		o.Order = "pub_date_desc"
	}
	o.Offset = model.IntFromStr(c.Query("offset"))

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}
