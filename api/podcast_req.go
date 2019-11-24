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
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type SubscribeToPodcastReq struct {
	PodcastId     int64 `validate:"required"`
	CurrentUserId int64 `validate:"required"`
}

func (o *SubscribeToPodcastReq) Load(c *Context) (err error) {
	o.CurrentUserId = c.session.UserId
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type UnsubscribeToPodcastReq struct {
	PodcastId     int64 `validate:"required"`
	CurrentUserId int64 `validate:"required"`
}

func (o *UnsubscribeToPodcastReq) Load(c *Context) (err error) {
	o.CurrentUserId = c.session.UserId
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}
