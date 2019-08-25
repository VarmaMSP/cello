package app

import "github.com/varmamsp/cello/model"

func (app *App) GetEpisodes(podcastId string, limit, offset int) ([]*model.EpisodeInfo, *model.AppError) {
	return app.store.Episode().GetAllByPodcast(podcastId, limit, offset)
}
