package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetEpisode(episodeId string) (*model.Episode, *model.AppError) {
	return app.Store.Episode().Get(episodeId)
}

func (app *App) GetEpisodesByIds(episodeIds []string) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByIds(episodeIds)
}

func (app *App) GetEpisodesInPodcast(podcastId, order string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByPodcast(podcastId, order, offset, limit)
}

func (app *App) GetAllEpisodesPublishedBefore(podcastIds []string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllPublishedBefore(podcastIds, offset, limit)
}

func (app *App) GetAllEpisodePlaybacks(episodeIds []string, userId string) ([]*model.EpisodePlayback, *model.AppError) {
	return app.Store.Episode().GetAllPlaybacks(episodeIds, userId)
}

func (app *App) GetAllEpisodePlaybacksByUser(userId string) ([]*model.EpisodePlayback, *model.AppError) {
	return app.Store.Episode().GetAllPlaybacksByUser(userId)
}

func (app *App) SaveEpisodePlayback(episodeId, userId string) *model.AppError {
	return app.Store.Episode().SavePlayback(&model.EpisodePlayback{
		EpisodeId: episodeId,
		PlayedBy:  userId,
	})
}

func (app *App) SaveEpisodeProgress(episodeId, userId string, currentTime int) {
	app.SyncEpisodePlaybackP.D <- &model.EpisodePlayback{
		EpisodeId:   episodeId,
		PlayedBy:    userId,
		CurrentTime: currentTime,
	}
}
