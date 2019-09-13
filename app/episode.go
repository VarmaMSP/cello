package app

import "github.com/varmamsp/cello/model"

func (app *App) GetEpisodeInfo(episodeId string) (*model.EpisodeInfo, *model.AppError) {
	return app.Store.Episode().GetInfo(episodeId)
}

func (app *App) GetEpisodes(podcastId string, offset, limit int) ([]*model.EpisodeInfo, *model.AppError) {
	return app.Store.Episode().GetAllByPodcast(podcastId, limit, offset)
}
