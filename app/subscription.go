package app

import "github.com/varmamsp/cello/model"

func (app *App) IsUserSubscribedToPodcast(userId, podcastId int64) (bool, *model.AppError) {
	return app.Store.Subscription().IsUserSubscribed(userId, podcastId)
}

func (app *App) AddPodcastToSubscriptions(userId, podcastId int64) *model.AppError {
	return app.Store.Subscription().Save(&model.Subscription{
		UserId:    userId,
		PodcastId: podcastId,
		Active:    1,
	})
}

func (app *App) RemovePodcastFromSubscriptions(userId, podcastId int64) *model.AppError {
	return app.Store.Subscription().Delete(userId, podcastId)
}
