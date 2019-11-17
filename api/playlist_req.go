package api

import "github.com/varmamsp/cello/model"

type GetPlaylistReq struct {
	PlaylistId string `validate:"required"`
}

func (o *GetPlaylistReq) Load(c *Context) *model.AppError {
	o.PlaylistId = c.Param("playlistId")

	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.get_playlist_req.load", err.Error(), 400, map[string]string{"playlist_id": o.PlaylistId})
	}
	return nil
}

type CreatePlaylistReq struct {
	Title         string `json:"title" validate:"required,max=100"`
	Privacy       string `json:"privacy" validate:"required,oneof=PUBLIC PRIVATE"`
	CurrentUserId string `json:"-" validate:"required"`
}

func (o *CreatePlaylistReq) Load(c *Context) *model.AppError {
	o.CurrentUserId = c.session.UserId

	if err := c.DecodeBody(o); err != nil {
		return model.NewAppError("api.create_playlist_req.load", err.Error(), 400, nil)
	}
	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.create_playlist_req.load", err.Error(), 400, nil)
	}
	return nil
}

type AddEpisodeToPlaylistReq struct {
	PlaylistId string `validate:"required"`
	EpisodeId  string `validate:"required"`
}

func (o *AddEpisodeToPlaylistReq) Load(c *Context) *model.AppError {
	o.PlaylistId = c.Param("playlistId")
	o.EpisodeId = c.Param("epsiodeId")

	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.create_playlist_req.load", err.Error(), 400, nil)
	}
	return nil
}
