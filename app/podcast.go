package app

import "github.com/varmamsp/cello/model"

func (app *App) GetPodcast(podcastId string) (*model.PodcastInfo, *model.AppError) {
	return app.store.Podcast().GetInfo(podcastId)
}
