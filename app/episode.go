package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetEpisode(episodeId int64) (*model.Episode, *model.AppError) {
	return app.Store.Episode().Get(episodeId)
}

func (app *App) GetEpisodesByIds(episodeIds []int64) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByIds(episodeIds)
}

func (app *App) GetEpisodesInPodcast(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByPodcast(podcastId, order, offset, limit)
}

func (app *App) GetAllEpisodesPublishedBefore(podcastIds []int64, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllPublishedBefore(podcastIds, offset, limit)
}

func (app *App) GetAllEpisodePlaybacks(episodeIds []int64, userId int64) ([]*model.EpisodePlayback, *model.AppError) {
	return app.Store.Episode().GetAllPlaybacks(episodeIds, userId)
}

func (app *App) GetAllEpisodePlaybacksByUser(userId int64, offset, limit int) ([]*model.EpisodePlayback, *model.AppError) {
	return app.Store.Episode().GetAllPlaybacksByUser(userId, offset, limit)
}

func (app *App) SaveEpisodePlayback(episodeId, userId int64) *model.AppError {
	return app.Store.Episode().SavePlayback(&model.EpisodePlayback{
		EpisodeId: episodeId,
		PlayedBy:  userId,
	})
}

func (app *App) SaveEpisodeProgress(episodeId, userId int64, currentTime int) {
	app.SyncEpisodePlaybackP.D <- &model.EpisodePlayback{
		EpisodeId:   episodeId,
		PlayedBy:    userId,
		CurrentTime: currentTime,
	}
}
