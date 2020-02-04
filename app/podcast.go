package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetPodcast(podcastId int64, includeCategories bool) (*model.Podcast, *model.AppError) {
	podcast, err := app.Store.Podcast().Get(podcastId)
	if err != nil {
		return nil, err
	}
	if !includeCategories {
		return podcast, err
	}

	podcast.Categories, err = app.Store.Category().GetPodcastCategories(podcastId)
	if err != nil {
		return nil, err
	}
	if podcast.Categories == nil {
		podcast.Categories = []*model.PodcastCategory{}
	}
	return podcast, nil
}

func (app *App) GetPodcastsByIds(podcastIds []int64) ([]*model.Podcast, *model.AppError) {
	if len(podcastIds) == 0 {
		return []*model.Podcast{}, nil
	}
	return app.Store.Podcast().GetByIds(podcastIds)
}

func (app *App) GetUserSubscriptions(userId int64) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetSubscriptions(userId)
}
