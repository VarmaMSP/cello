package api

import "github.com/varmamsp/cello/model"

type GetSubscriptionFeedReq struct {
	Offset int   `validate:"min=0"`
	Limit  int   `validate:"min=5"`
	UserId int64 `validate:"-"`
}

func (o *GetSubscriptionFeedReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	o.Offset = model.IntFromStr(c.Query("offset"))
	o.Limit = model.IntFromStr(c.Query("limit"))

	if o.Limit == 0 {
		o.Limit = 20
	}

	err = c.app.Validate.Struct(o)
	return nil
}

type SubscribeToPodcastReq struct {
	UserId    int64 `validate:"required"`
	PodcastId int64 `validate:"required"`
}

func (o *SubscribeToPodcastReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}

	err = c.app.Validate.Struct(o)
	return
}

type UnsubscribeToPodcastReq struct {
	UserId    int64 `validate:"required"`
	PodcastId int64 `validate:"required"`
}

func (o *UnsubscribeToPodcastReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	if o.PodcastId, err = model.Int64FromHashId(c.Param("podcastId")); err != nil {
		return
	}

	err = c.app.Validate.Struct(o)
	return
}
