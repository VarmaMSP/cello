package app

import (
	"time"

	"github.com/varmamsp/cello/model"
)

func (app *App) GetEpisodeInfo(episodeId string) (*model.Episode, *model.AppError) {
	return app.Store.Episode().Get(episodeId)
}

func (app *App) GetEpisodes(podcastId string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByPodcast(podcastId, limit, offset)
}

func (app *App) GetAllEpisodesPubblishedBetween(from, to *time.Time, podcastIds []string) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllPublishedBetween(from, to, podcastIds)
}
