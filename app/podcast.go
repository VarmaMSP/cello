package app

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func (app *App) GetPodcast(podcastId int64) (*model.Podcast, *model.AppError) {
	return app.Store.Podcast().Get(podcastId)
}

func (app *App) GetUserSubscriptions(userId int64) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetSubscriptions(userId)
}

func (app *App) SearchPodcasts(searchQuery string) ([]*model.PodcastIndex, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index("podcast").
		Query(elastic.NewMultiMatchQuery(searchQuery, "title", "author")).
		Size(100).
		Do(context.TODO())
	if err != nil {
		return nil, nil
	}

	podcasts := []*model.PodcastIndex{}
	for _, item := range results.Each(reflect.TypeOf(model.PodcastIndex{})) {
		tmp, _ := item.(model.PodcastIndex)
		tmp.Description = ""
		podcasts = append(podcasts, &tmp)
	}

	return podcasts, nil
}

func (app *App) SaveSubscription(userId, podcastId int64) *model.AppError {
	return app.Store.Subscription().Save(&model.Subscription{
		UserId:    userId,
		PodcastId: podcastId,
		Active:    1,
	})
}

func (app *App) DeleteSubscription(userId, podcastId int64) *model.AppError {
	return app.Store.Subscription().Delete(userId, podcastId)
}
