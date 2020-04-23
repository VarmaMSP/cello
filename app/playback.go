package app

import "github.com/varmamsp/cello/model"

func (a *App) GetPlaybacks(userId int64, offset, limit int) ([]*model.Playback, *model.AppError) {
	return a.Store.Playback().GetByUserPaginated(userId, offset, limit)
}

func (a *App) GetPlaybacksForEpisodes(userId int64, episodeIds []int64) ([]*model.Playback, *model.AppError) {
	if len(episodeIds) == 0 {
		return []*model.Playback{}, nil
	}
	return a.Store.Playback().GetByUserByEpisodes(userId, episodeIds)
}

func (a *App) LoadPlaybacks(userId int64, episodes []*model.Episode) *model.AppError {
	if len(episodes) == 0 {
		return nil
	}

	episodeIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		episodeIds[i] = episode.Id
	}
	playbacks, err := a.Store.Playback().GetByUserByEpisodes(userId, episodeIds)
	if err != nil {
		return err
	}

	playbackByEpisodeId := map[int64]*model.Playback{}
	for _, playback := range playbacks {
		playbackByEpisodeId[playback.EpisodeId] = playback
	}

	for _, episode := range episodes {
		if playback, ok := playbackByEpisodeId[episode.Id]; ok {
			episode.Progress = playback.CurrentProgress
			episode.LastPlayedAt = playback.LastPlayedAt
		}
	}

	return nil
}

func (app *App) SyncPlaybackBegin(userId, episodeId int64) *model.AppError {
	return app.Store.Playback().Upsert(&model.Playback{
		UserId:    userId,
		EpisodeId: episodeId,
	})
}

func (app *App) SyncPlaybackProgress(episodeId, userId int64, position float64) *model.AppError {
	if err := app.SyncPlaybackP.Publish(&model.PlaybackEvent{
		Event:     model.PLAYBACK_EVENT_PLAYING,
		UserId:    userId,
		EpisodeId: episodeId,
		Position:  position,
	}); err != nil {
		return model.New500Error("sync_playbnack_progress", err.Error(), nil)
	}

	return nil
}
