package app

import "github.com/varmamsp/cello/model"

func (app *App) GetPodcast(podcastId string) (*model.PodcastInfo, *model.AppError) {
	return app.store.Podcast().GetInfo(podcastId)
}

func (app *App) CreatePodcastCuration(title string) *model.AppError {
	return app.store.PodcastCuration().Save(&model.PodcastCuration{
		Title: title,
	})
}

func (app *App) GetAllPodcastCurations() ([]*model.PodcastCuration, *model.AppError) {
	return app.store.PodcastCuration().GetAll()
}
