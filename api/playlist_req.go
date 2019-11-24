package api

import "github.com/varmamsp/cello/model"

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
	Title         string `json:"title" validate:"required,max=100"`
	Privacy       string `json:"privacy" validate:"required,oneof=PUBLIC PRIVATE"`
	CurrentUserId int64  `json:"-" validate:"required"`
}

func (o *CreatePlaylistReq) Load(c *Context) (err error) {
	o.CurrentUserId = c.session.UserId

	if err = c.DecodeBody(o); err != nil {
		return
	}
	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}

type AddEpisodeToPlaylistReq struct {
	PlaylistId int64 `validate:"required"`
	EpisodeId  int64 `validate:"required"`
}

func (o *AddEpisodeToPlaylistReq) Load(c *Context) (err error) {
	if o.PlaylistId, err = model.Int64FromHashId(c.Param("playlistId")); err != nil {
		return
	}
	if o.EpisodeId, err = model.Int64FromHashId(c.Param("epsiodeId")); err != nil {
		return
	}

	if err = c.app.Validate.Struct(o); err != nil {
		return
	}
	return
}
