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

type GetEpisodePlaybacksReq struct {
	EpisodeIds    []int64 `json:"episode_ids" validate:"gt=0,dive,required"`
	CurrentUserId int64   `json:"-" validate:"required"`
}

type GetEpisodePlaybacksReq_ struct {
	EpisodeIds []string `json:"episode_ids"`
}

func (o *GetEpisodePlaybacksReq) Load(c *Context) (err error) {
	var tmp GetEpisodePlaybacksReq_
	if err = c.DecodeBody(tmp); err != nil {
		return
	}

	episodeIds := make([]int64, len(tmp.EpisodeIds))
	for i, id := range tmp.EpisodeIds {
		if episodeIds[i], err = model.Int64FromHashId(id); err != nil {
			return
		}
	}

	o.EpisodeIds = episodeIds
	o.CurrentUserId = c.session.UserId

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type SyncPlaybackReq struct {
	EpisodeId     int64 `validate:"required"`
	CurrentUserId int64 `validate:"required"`
}

func (o *SyncPlaybackReq) Load(c *Context) (err error) {
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("episodeId")); err != nil {
		return
	}
	o.CurrentUserId = c.session.UserId

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type SyncPlaybackProgressReq struct {
	EpisodeId     int64 `json:"-" validate:"required"`
	CurrentTime   int   `json:"current_time" validate:"-"`
	CurrentUserId int64 `json:"-" validate:"required"`
}

func (o *SyncPlaybackProgressReq) Load(c *Context) (err error) {
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("episodeId")); err != nil {
		return
	}
	if err = c.DecodeBody(o); err != nil {
		return
	}
	o.CurrentUserId = c.session.UserId

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}
