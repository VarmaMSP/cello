package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetPodcast(podcastId int64) (*model.Podcast, *model.AppError) {
	return app.Store.Podcast().Get(podcastId)
}

func (app *App) GetPodcastsByIs(podcastIds []int64) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetByIds(podcastIds)
}

func (app *App) GetUserSubscriptions(userId int64) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetSubscriptions(userId)
}
