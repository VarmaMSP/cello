package app

import "github.com/varmamsp/cello/model"

func (app *App) GetCuration(curationId string) (*model.Curation, *model.AppError) {
	return app.Store.Curation().Get(curationId)
}

func (app *App) GetCurations() ([]*model.Curation, *model.AppError) {
	return app.Store.Curation().GetAll()
}

func (app *App) SaveCuration(title string) *model.AppError {
	return app.Store.Curation().Save(&model.Curation{Title: title})
}

func (app *App) DeleteCuration(curationId string) *model.AppError {
	return app.Store.Curation().Delete(curationId)
}

func (app *App) SavePodcastToCuration(curationId, podcastId string) *model.AppError {
	return app.Store.Curation().SavePodcastCuration(&model.PodcastCuration{PodcastId: podcastId, CurationId: curationId})
}
