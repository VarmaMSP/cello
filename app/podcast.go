package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetFeed(id int64) (*model.Feed, *model.AppError) {
	return app.Store.Feed().Get(id)
}

func (a *App) GetPodcast(podcastId int64) (*model.Podcast, *model.AppError) {
	podcast, err := a.Store.Podcast().Get(podcastId)
	if err != nil {
		return nil, err
	}

	podcast.Categories, err = a.Store.Category().GetPodcastCategories(podcastId)
	if err != nil {
		return nil, err
	}

	return podcast, nil
}

func (a *App) GetPodcastsForEpisodes(episodes []*model.Episode) ([]*model.Podcast, *model.AppError) {
	podcastIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		podcastIds[i] = episode.PodcastId
	}

	return a.GetPodcastsByIds(podcastIds)
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
