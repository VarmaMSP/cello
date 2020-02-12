package app

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

func (app *App) SuggestKeywords(phrase, prefix string) ([]*model.SearchSuggestion, *model.AppError) {
	var err error
	var results *elastic.SearchResult

	if phrase != "" && prefix != "" {
		results, err = app.KeywordPhrasePrefixSearch(phrase, prefix)
	} else if phrase != "" && prefix == "" {
		results, err = app.KeywordPhraseSearch(phrase)
	} else {
		results, err = app.KeywordPrefixSearch(prefix)
	}

	if err != nil {
		return nil, model.NewAppError(
			"app.suggestKeywords", err.Error(), http.StatusInternalServerError, nil,
		)
	}

	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.SearchSuggestion{}, nil
	}

	searchSuggestions := []*model.SearchSuggestion{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.SearchSuggestion{}
		if err := tmp.LoadFromKeyword(hit); err == nil {
			searchSuggestions = append(searchSuggestions, tmp)
		}
	}

	return searchSuggestions, nil
}

func (app *App) SuggestPodcasts(phrase string) ([]*model.SearchSuggestion, *model.AppError) {
	results, err := app.PodcastPharseSearch(phrase)
	if err != nil {
		return nil, model.NewAppError(
			"app.typeahead_podcasts", "no results", http.StatusInternalServerError, nil,
		)
	}

	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.SearchSuggestion{}, nil
	}

	searchSuggestions := []*model.SearchSuggestion{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.SearchSuggestion{}
		if err := tmp.LoadFromPodcast(hit); err == nil {
			searchSuggestions = append(searchSuggestions, tmp)
		}
	}

	return searchSuggestions, nil
}

func (app *App) KeywordPhrasePrefixSearch(phrase, prefix string) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.KeywordIndexName).
		Query(elastic.NewBoolQuery().
			Must(
				elastic.NewFuzzyQuery("text", phrase).
					Fuzziness("AUTO").
					MaxExpansions(30).
					PrefixLength(0).
					Transpositions(true),
				elastic.NewTermQuery("text.prefix", prefix),
			),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(10).
			PreTags("<em>").
			PostTags("</em>").
			Fields(elastic.NewHighlighterField("text")),
		).
		Size(5).
		Do(context.TODO())
}

func (app *App) KeywordPhraseSearch(phrase string) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.KeywordIndexName).
		Query(elastic.NewFuzzyQuery("text", phrase).
			Fuzziness("AUTO").
			MaxExpansions(30).
			PrefixLength(0).
			Transpositions(true),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(10).
			PreTags("<em>").
			PostTags("</em>").
			Fields(elastic.NewHighlighterField("text")),
		).
		Size(5).
		Do(context.TODO())
}

func (app *App) KeywordPrefixSearch(prefix string) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.KeywordIndexName).
		Query(elastic.NewTermQuery("text.prefix", prefix)).
		Highlight(elastic.NewHighlight().
			FragmentSize(10).
			PreTags("<em>").
			PostTags("</em>").
			Fields(elastic.NewHighlighterField("text")),
		).
		Size(5).
		Do(context.TODO())
}

func (app *App) PodcastPharseSearch(phrase string) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewMultiMatchQuery(phrase).
			Type("bool_prefix").
			Field("title").
			Field("title._2gram").
			Field("title._3gram").
			Field("author").
			Field("author._2gram").
			Field("author._3gram"),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(200).
			PreTags("<em>").
			PostTags("</em>").
			Fields(
				elastic.NewHighlighterField("title"),
				elastic.NewHighlighterField("author"),
			),
		).
		Size(4).
		Do(context.TODO())
}

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
