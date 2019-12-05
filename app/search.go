package app

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func (app *App) SearchPodcasts(searchQuery string) ([]*model.Podcast, *model.AppError) {
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
		podcasts = append(podcasts, &tmp)
	}

	podcasts_ := make([]*model.Podcast, len(podcasts))
	for i, podcast := range podcasts {
		id, _ := model.Int64FromHashId(podcast.Id)
		podcasts_[i] = &model.Podcast{
			Id:     id,
			Title:  podcast.Title,
			Author: podcast.Author,
			Type:   podcast.Type,
		}
	}
	return podcasts_, nil
}
