package app

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

func (app *App) SearchPodcasts(searchQuery string, offset, limit int) ([]*model.Podcast, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index("podcast").
		Query(elastic.NewMultiMatchQuery(searchQuery).
			FieldWithBoost("title", 3).
			Field("author")).
		Highlight(elastic.NewHighlight().
			PreTags("<span class=\"result-highlight\">").
			PostTags("</span>").
			Field("title"),
		).
		From(offset).
		Size(limit).
		Do(context.TODO())

	if err != nil {
		return nil, model.NewAppError("app.search_podcasts", "no results", http.StatusInternalServerError, nil)
	}

	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.Podcast{}, nil
	}

	podcasts := []*model.Podcast{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.Podcast{}
		if err := tmp.LoadFromESHit(hit); err == nil {
			podcasts = append(podcasts, tmp)
		}
	}
	return podcasts, nil
}
