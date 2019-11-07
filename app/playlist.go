package app

import "github.com/varmamsp/cello/model"

func (app *App) CreatePlaylist(title, privacy, userId string) (*model.Playlist, *model.AppError) {
	playlist := &model.Playlist{
		Title:     title,
		CreatedBy: userId,
		Privacy:   privacy,
	}

	if err := app.Store.Playlist().Save(playlist); err != nil {
		return nil, err
	}
	return playlist, nil
}

func (app *App) GetPlaylist(playlistId string) (*model.Playlist, *model.AppError) {
	return app.Store.Playlist().Get(playlistId)
}

func (app *App) GetPlaylistsByUser(userId string) ([]*model.Playlist, *model.AppError) {
	return app.Store.Playlist().GetAllByUser(userId)
}

func (app *App) GetEpisodesInPlaylist(playlistId string) ([]*model.Episode, *model.AppError) {
	return app.Store.Playlist().GetAllEpisodesInPlaylist(playlistId)
}

func (app *App) AddEpsiodeToPlaylist(episodeId, playlistId string) (*model.PlaylistItem, *model.AppError) {
	playlistItem := &model.PlaylistItem{
		PlaylistId: playlistId,
		EpisodeId:  episodeId,
	}

	if err := app.Store.Playlist().SaveItem(playlistItem); err != nil {
		return nil, err
	}
	return playlistItem, nil
}
