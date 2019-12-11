package app

import "github.com/varmamsp/cello/model"

func (app *App) SavePlaylist(title, privacy string, userId int64) (*model.Playlist, *model.AppError) {
	playlist := &model.Playlist{
		UserId:  userId,
		Title:   title,
		Privacy: privacy,
	}

	if err := app.Store.Playlist().Save(playlist); err != nil {
		return nil, err
	}
	return playlist, nil
}

func (app *App) GetPlaylist(playlistId int64) (*model.Playlist, *model.AppError) {
	return app.Store.Playlist().Get(playlistId)
}

func (app *App) GetPlaylistsByUser(userId int64) ([]*model.Playlist, *model.AppError) {
	return app.Store.Playlist().GetByUserPaginated(userId, 0, 1000)
}

func (app *App) GetEpisodesInPlaylist(playlistId int64) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetByPlaylistPaginated(playlistId, 0, 1000)
}

func (app *App) SaveEpisodeToPlaylist(episodeId, playlistId int64) (*model.PlaylistMember, *model.AppError) {
	playlistMember := &model.PlaylistMember{
		PlaylistId: playlistId,
		EpisodeId:  episodeId,
	}

	if err := app.Store.Playlist().SaveMember(playlistMember); err != nil {
		return nil, err
	}
	return playlistMember, nil
}

func (app *App) DeleteEpisodeFromPlaylist(episodeId, playlistId int64) *model.AppError {
	return app.Store.Playlist().DeleteMember(playlistId, episodeId)
}
