package app

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

func (app *App) SearchPodcasts(searchQuery string, offset, limit int) ([]*model.PodcastSearchResult, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewMultiMatchQuery(searchQuery).
			FieldWithBoost("title", 1.7).
			FieldWithBoost("author", 1.1).
			Field("description").
			TieBreaker(0.4),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(200).
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

func (app *App) SearchPodcastsByPhrase(phrase string) ([]*model.PodcastSearchResult, *model.AppError) {
	results, err := app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewMultiMatchQuery(phrase).
			Type("phrase").
			FieldWithBoost("title", 1.5).
			Field("author").
			TieBreaker(0.4),
		).
		Size(6).
		Do(context.TODO())

	if err != nil {
		return nil, model.NewAppError("app.search_podcasts", err.Error(), http.StatusInternalServerError, nil)
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

func (app *App) SearchEpisodes(searchQuery, sortBy string, offset, limit int) ([]*model.EpisodeSearchResult, *model.AppError) {
	q := app.ElasticSearch.Search().
		Index(elasticsearch.EpisodeIndexName).
		Query(elastic.NewMultiMatchQuery(searchQuery).
			Field("title").
			Field("description").
			TieBreaker(0.2),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(200).
			PreTags("<span class=\"result-highlight\">").
			PostTags("</span>").
			Fields(
				elastic.NewHighlighterField("title"),
				elastic.NewHighlighterField("description"),
			),
		).
		From(offset).
		Size(limit)

	if sortBy == "publish_date" {
		q.Sort("pub_date", false)
	}

	results, err := q.Do(context.TODO())
	if err != nil {
		return nil, model.NewAppError("app.search_podcasts", "no results", http.StatusInternalServerError, nil)
	}

	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.EpisodeSearchResult{}, nil
	}

	episodeSearchResults := []*model.EpisodeSearchResult{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.EpisodeSearchResult{}
		if err := tmp.LoadDetails(hit); err == nil {
			episodeSearchResults = append(episodeSearchResults, tmp)
		}
	}
	return episodeSearchResults, nil
}
