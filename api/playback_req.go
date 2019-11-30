package api

import (
	"strings"

	"github.com/varmamsp/cello/model"
)

type GetPlaybacksReq struct {
	UserId     int64 `validate:"required"`
	EpisodeIds []int64
}

func (o *GetPlaybacksReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId

	episodeIds := []int64{}
	episodeIdHashs := strings.Split(c.Query("episode_ids"), ",")
	for _, episodeIdHash := range episodeIdHashs {
		if episodeId, err := model.Int64FromHashId(strings.TrimSpace(episodeIdHash)); err == nil {
			episodeIds = append(episodeIds, episodeId)
		}
	}
	o.EpisodeIds = episodeIds

	err = c.app.Validate.Struct(o)
	return
}

type StartPlaybackReq struct {
	UserId    int64 `validate:"required"`
	EpisodeId int64 `validate:"required"`
}

func (o *StartPlaybackReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("episodeId")); err != nil {
		return
	}

	err = c.app.Validate.Struct(o)
	return
}

type SyncPlaybackReq struct {
	UserId    int64   `validate:"required"`
	EpisodeId int64   `validate:"required"`
	Progress  float32 `validate:"required"`
}

func (o *SyncPlaybackReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("episodeId")); err != nil {
		return
	}

	err = c.app.Validate.Struct(o)
	return
}
