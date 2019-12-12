package api

import "github.com/varmamsp/cello/model"

type GetPlaylistsReq struct {
	UserId      int64 `validate:"-"`
	FullDetails bool  `validate:"-"`
	Offset      int   `validate:"min=0"`
	Limit       int   `validate:"min=5"`
}

func (o *GetPlaylistsReq) Load(c *Context) (err error) {
	if c.Query("user_id") == "" {
		o.UserId = c.session.UserId
	} else {
		o.UserId, err = model.Int64FromHashId(c.Query("user_id"))
		if err != nil {
			return
		}
	}

	o.FullDetails = model.BoolFromStr(c.Query("full_details"))
	o.Offset = model.IntFromStr(c.Query("offset"))
	o.Limit = model.IntFromStr(c.Query("limit"))

	err = c.app.Validate.Struct(o)
	return
}

type GetPlaylistReq struct {
	PlaylistId int64 `validate:"required"`
}

func (o *GetPlaylistReq) Load(c *Context) (err error) {
	if o.PlaylistId, err = model.Int64FromHashId(c.Param("playlistId")); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return nil
}

type CreatePlaylistReq struct {
	Title   string `json:"title" validate:"required,max=100"`
	Privacy string `json:"privacy" validate:"required,oneof=PUBLIC PRIVATE"`
	UserId  int64  `json:"-" validate:"required"`
}

func (o *CreatePlaylistReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId

	if err = c.DecodeBody(o); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type AddEpisodeToPlaylistsReq struct {
	EpisodeId   int64   `validate:"required"`
	PlaylistIds []int64 `validate:"required"`
}

func (o *AddEpisodeToPlaylistsReq) Load(c *Context) (err error) {
	aux := &struct {
		EpisodeId   string   `json:"episode_id"`
		PlaylistIds []string `json:"playlist_ids"`
	}{}
	if err = c.DecodeBody(aux); err != nil {
		return
	}

	if o.EpisodeId, err = model.Int64FromHashId(aux.EpisodeId); err != nil {
		return
	}
	playlistIds := make([]int64, len(aux.PlaylistIds))
	for i, playlistId := range aux.PlaylistIds {
		if playlistIds[i], err = model.Int64FromHashId(playlistId); err != nil {
			return
		}
	}
	o.PlaylistIds = playlistIds

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}
