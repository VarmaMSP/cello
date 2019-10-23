package app

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func (app *App) GetPodcast(podcastId string) (*model.Podcast, *model.AppError) {
	return app.Store.Podcast().Get(podcastId)
}

func (app *App) GetPodcastsInCuration(curationId string) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetAllByCuration(curationId, 0, 7)
}

func (app *App) GetUserSubscriptions(userId string) ([]*model.Podcast, *model.AppError) {
	return app.Store.Podcast().GetAllSubscribedBy(userId)
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

func (app *App) CreateSubscription(userId, podcastId string) *model.AppError {
	m := &model.PodcastSubscription{
		PodcastId:    podcastId,
		SubscribedBy: userId,
		Active:       1,
	}
	return app.Store.Podcast().SaveSubscription(m)
}

func (app *App) DeleteSubscription(userId, podcastId string) *model.AppError {
	return app.Store.Podcast().DeleteSubscription(userId, podcastId)
}
