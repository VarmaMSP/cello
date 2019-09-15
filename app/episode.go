package app

import "github.com/varmamsp/cello/model"

func (app *App) GetEpisodeInfo(episodeId string) (*model.Episode, *model.AppError) {
	return app.Store.Episode().Get(episodeId)
}

func (app *App) GetEpisodes(podcastId string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return app.Store.Episode().GetAllByPodcast(podcastId, limit, offset)
}
