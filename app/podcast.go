package app

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func (app *App) GetPodcastInfo(podcastId string) (*model.PodcastInfo, *model.AppError) {
	return app.Store.Podcast().GetInfo(podcastId)
}

func (app *App) GetPodcastFeedDetails(podcastId string) (*model.PodcastFeedDetails, *model.AppError) {
	return app.Store.Podcast().GetFeedDetails(podcastId)
}

func (app *App) GetPodcastsInCuration(curationId string) ([]*model.PodcastInfo, *model.AppError) {
	return app.Store.Curation().GetPodcastsByCuration(curationId, 0, 7)
}

func (app *App) SearchPodcasts(searchQuery string) ([]*model.PodcastInfo, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index("podcast").
		Query(elastic.NewMultiMatchQuery(searchQuery, "title", "author")).
		Size(100).
		Do(context.TODO())
	if err != nil {
		return nil, nil
	}

	podcasts := []*model.PodcastInfo{}
	for _, item := range results.Each(reflect.TypeOf(model.PodcastInfo{})) {
		tmp, _ := item.(model.PodcastInfo)
		tmp.Description = ""
		podcasts = append(podcasts, &tmp)
	}

	return podcasts, nil
}
