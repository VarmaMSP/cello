package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetEpisode(episodeId int64) (*model.Episode, *model.AppError) {
	return app.Store.Episode().Get(episodeId)
}

func (app *App) GetEpisodesByIds(episodeIds []int64) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetByIds(episodeIds)
}

func (app *App) GetEpisodesInPodcast(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetByPodcastPaginated(podcastId, order, offset, limit)
}

func (app *App) GetEpisodesInPodcastIds(podcastIds []int64, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetByPodcastIdsPaginated(podcastIds, offset, limit)
}
