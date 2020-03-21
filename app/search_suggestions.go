package app

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

func (app *App) SuggestKeywords(phrase, prefix string, noFuzzy int) ([]*model.SearchSuggestion, *model.AppError) {
	var err error
	var results *elastic.SearchResult

	if phrase != "" && prefix != "" {
		results, err = app.KeywordPhrasePrefixSearch(phrase, prefix, noFuzzy)
	} else if phrase != "" && prefix == "" {
		results, err = app.KeywordPhraseSearch(phrase, noFuzzy)
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

func (app *App) SuggestPodcasts(phrase, prefix string, noFuzzy int) ([]*model.SearchSuggestion, *model.AppError) {
	var err error
	var results *elastic.SearchResult

	if phrase != "" && prefix != "" {
		results, err = app.PodcastPhrasePrefixSearch(phrase, prefix, noFuzzy)
	} else if phrase != "" && prefix == "" {
		results, err = app.PodcastPhraseSearch(phrase, noFuzzy)
	} else {
		results, err = app.PodcastPrefixSearch(prefix)
	}

	if err != nil {
		return nil, model.NewAppError(
			"app.typeahead_podcasts", err.Error(), http.StatusInternalServerError, nil,
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

func (app *App) KeywordPhrasePrefixSearch(phrase, prefix string, noFuzzy int) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.KeywordIndexName).
		Query(elastic.NewBoolQuery().
			Must(
				elastic.NewFuzzyQuery("text", phrase).
					Fuzziness("AUTO").
					MaxExpansions(30).
					PrefixLength(noFuzzy).
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

func (app *App) KeywordPhraseSearch(phrase string, noFuzzy int) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.KeywordIndexName).
		Query(elastic.NewFuzzyQuery("text", phrase).
			Fuzziness("AUTO").
			MaxExpansions(30).
			PrefixLength(noFuzzy).
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

func (app *App) PodcastPhrasePrefixSearch(phrase, prefix string, noFuzzy int) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewBoolQuery().
			Must(
				elastic.NewFuzzyQuery("title.shingles", phrase).
					Fuzziness("AUTO").
					MaxExpansions(30).
					PrefixLength(noFuzzy).
					Transpositions(true),
				elastic.NewTermQuery("title.prefix", prefix),
			),
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

func (app *App) PodcastPhraseSearch(phrase string, noFuzzy int) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewFuzzyQuery("title.shingles", phrase).
			Fuzziness("AUTO").
			MaxExpansions(30).
			PrefixLength(noFuzzy).
			Transpositions(true),
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

func (app *App) PodcastPrefixSearch(prefix string) (*elastic.SearchResult, error) {
	return app.ElasticSearch.Search().
		Index(elasticsearch.PodcastIndexName).
		Query(elastic.NewTermQuery("title.prefix", prefix)).
		Highlight(elastic.NewHighlight().
			FragmentSize(10).
			PreTags("<em>").
			PostTags("</em>").
			Fields(elastic.NewHighlighterField("text")),
		).
		Size(5).
		Do(context.TODO())
}
