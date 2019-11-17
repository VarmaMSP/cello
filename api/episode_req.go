package api

import (
	"fmt"
	"time"

	"github.com/varmamsp/cello/model"
)

type GetFeedReq struct {
	Limit           int
	PublishedBefore *time.Time
	CurrentUserId   string
}

func (o *GetFeedReq) Load(c *Context) *model.AppError {
	o.CurrentUserId = c.session.UserId
	o.Limit = model.IntFromStr(c.Query("limit"))
	if o.Limit == 0 {
		o.Limit = 20
	}
	o.PublishedBefore = model.ParseDateTime(c.Query("published_before"))
	if o.PublishedBefore == nil {
		now := time.Now()
		o.PublishedBefore = &now
	}

	return nil
}

type GetEpisodePlaybacksReq struct {
	EpisodeIds    []string `json:"episode_ids" validate:"gt=0,dive,required"`
	CurrentUserId string   `json:"-" validate:"required"`
}

func (o *GetEpisodePlaybacksReq) Load(c *Context) *model.AppError {
	o.CurrentUserId = c.session.UserId

	if err := c.DecodeBody(o); err != nil {
		return model.NewAppError("api.get_episode_playbacks_req.load", err.Error(), 400, nil)
	}
	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.get_episode_playbacks_req.load", err.Error(), 400, nil)
	}
	return nil
}

type SyncPlaybackReq struct {
	EpisodeId     string `validate:"required"`
	CurrentUserId string `validate:"required"`
}

func (o *SyncPlaybackReq) Load(c *Context) *model.AppError {
	o.EpisodeId = c.Param("episodeId")
	o.CurrentUserId = c.session.UserId

	fmt.Println(o)

	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.sync_playback_req.load", err.Error(), 400, nil)
	}
	return nil
}

type SyncPlaybackProgressReq struct {
	EpisodeId     string `json:"-" validate:"required"`
	CurrentTime   int    `json:"current_time" validate:"-"`
	CurrentUserId string `json:"-" validate:"required"`
}

func (o *SyncPlaybackProgressReq) Load(c *Context) *model.AppError {
	o.EpisodeId = c.Param("episodeId")
	o.CurrentUserId = c.session.UserId

	if err := c.DecodeBody(o); err != nil {
		return model.NewAppError("api.sync_playback_progress_req.load", err.Error(), 400, nil)
	}
	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.sync_playback_progress_req.load", err.Error(), 400, nil)
	}
	return nil
}
