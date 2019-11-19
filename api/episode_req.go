package api

import (
	"time"

	"github.com/varmamsp/cello/model"
)

type GetPodcastEpisodesReq struct {
	PodcastId string `validate:"required"`
	Limit     int    `validate:"min=5"`
	Offset    int    `validate:"min=0"`
	Order     string `validate:"required,oneof=pub_date_desc pub_date_asc"`
}

func (o *GetPodcastEpisodesReq) Load(c *Context) *model.AppError {
	o.PodcastId = c.Param("podcastId")
	o.Limit = model.IntFromStr(c.Query("limit"))
	o.Offset = model.IntFromStr(c.Query("offset"))
	o.Order = c.Query("order")

	if o.Limit == 0 {
		o.Limit = 10
	}
	if o.Order == "" {
		o.Order = "pub_date_desc"
	}

	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.get_podcast_episodes_req.load", err.Error(), 400, nil)
	}
	return nil
}

type GetFeedReq struct {
	Limit           int        `validate:"min=5"`
	PublishedBefore *time.Time `validate:"-"`
	CurrentUserId   string     `validate:"-"`
}

func (o *GetFeedReq) Load(c *Context) *model.AppError {
	o.Limit = model.IntFromStr(c.Query("limit"))
	o.CurrentUserId = c.session.UserId
	o.PublishedBefore = model.ParseDateTime(c.Query("published_before"))

	if o.Limit == 0 {
		o.Limit = 20
	}
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
