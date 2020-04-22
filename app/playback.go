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

func (app *App) SyncPlaybackBegin(episodeId, userId int64) *model.AppError {
	return app.Store.Playback().Upsert(&model.Playback{
		UserId:    userId,
		EpisodeId: episodeId,
	})
}

func (app *App) SyncPlaybackProgress(episodeId, userId int64, event string, position float64) *model.AppError {

	// if event == model.PLAYBACK_EVENT_PLAYING {
	// 	app.SyncEpisodePlaybackP.D <- &model.PlaybackEvent{
	// 		Event:     model.PLAYBACK_EVENT_PLAYING,
	// 		UserId:    userId,
	// 		EpisodeId: episodeId,
	// 		Position:  position,
	// 	}
	// }
	return nil
}
