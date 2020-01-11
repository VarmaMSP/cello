package app

import "github.com/varmamsp/cello/model"

func (app *App) GetUserPlaybacks(userId int64, offset, limit int) ([]*model.Playback, *model.AppError) {
	return app.Store.Playback().GetByUserPaginated(userId, offset, limit)
}

func (app *App) GetUserPlaybacksForEpisodes(userId int64, episodeIds []int64) ([]*model.Playback, *model.AppError) {
	if len(episodeIds) == 0 {
		return []*model.Playback{}, nil
	}
	return app.Store.Playback().GetByUserByEpisodes(userId, episodeIds)
}

func (app *App) SyncPlayback(episodeId, userId int64, event string, position float64) *model.AppError {
	if event == model.PLAYBACK_EVENT_BEGIN {
		return app.Store.Playback().Upsert(&model.Playback{
			UserId:    userId,
			EpisodeId: episodeId,
		})
	}

	if event == model.PLAYBACK_EVENT_PLAYING {
		app.SyncEpisodePlaybackP.D <- &model.PlaybackEvent{
			Event:     model.PLAYBACK_EVENT_PLAYING,
			UserId:    userId,
			EpisodeId: episodeId,
			Position:  position,
		}
	}
	return nil
}
