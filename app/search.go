package app

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

func (app *App) SearchPodcasts(searchQuery string, offset, limit int) ([]*model.PodcastSearchResult, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index("podcast").
		Query(elastic.NewMultiMatchQuery(searchQuery).Type("best_fields").
			FieldWithBoost("title", 1).
			FieldWithBoost("author", 2).
			Field("description"),
		).
		Highlight(elastic.NewHighlight().
			PreTags("<span class=\"result-highlight\">").
			PostTags("</span>").
			Fields(
				elastic.NewHighlighterField("title"),
				elastic.NewHighlighterField("author"),
				elastic.NewHighlighterField("description"),
			),
		).
		From(offset).
		Size(limit).
		Do(context.TODO())

	if err != nil {
		return nil, model.NewAppError("app.search_podcasts", "no results", http.StatusInternalServerError, nil)
	}

	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.PodcastSearchResult{}, nil
	}

	podcastSearchResults := []*model.PodcastSearchResult{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.PodcastSearchResult{}
		if err := tmp.LoadDetails(hit); err == nil {
			podcastSearchResults = append(podcastSearchResults, tmp)
		}
	}
	return podcastSearchResults, nil
}
