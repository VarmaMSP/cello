package searchlayer

import (
	"context"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type searchPodcastStore struct {
	store.PodcastStore
	se searchengine.Broker
}

func (s *searchPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	if err := s.PodcastStore.Save(podcast); err != nil {
		return err
	}
	if err := s.se.Index(searchengine.PODCAST_INDEX, podcast.ForIndexing()); err != nil {
		return model.New500Error("search_layer.search_podcast_store.save", err.Error(), nil)
	}
	return nil
}

func (s *searchPodcastStore) GetTypeaheadSuggestions(query string) ([]*model.SearchSuggestion, *model.AppError) {
	words := strings.Split(query, " ")
	wordCount := len(words)

	phraseWords := words[:wordCount-1]
	phraseNoFuzzy := 0
	if len(phraseWords) > 1 {
		for i := 0; i < len(phraseWords)-2; i++ {
			phraseNoFuzzy += len(phraseWords[i])
		}
		phraseNoFuzzy += len(phraseWords) - 2
	}

	phrase := strings.Join(phraseWords, " ")
	prefix := words[wordCount-1]

	var err error
	var results *elastic.SearchResult

	if phrase != "" && prefix != "" {
		results, err = s.se.C().Search().
			Index(searchengine.PODCAST_INDEX).
			Query(elastic.NewBoolQuery().
				Must(
					elastic.NewFuzzyQuery("title.shingles", phrase).
						Fuzziness("AUTO").
						MaxExpansions(30).
						PrefixLength(phraseNoFuzzy).
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
			Size(6).
			Do(context.TODO())
	} else if phrase != "" && prefix == "" {
		results, err = s.se.C().Search().
			Index(searchengine.PODCAST_INDEX).
			Query(elastic.NewFuzzyQuery("title.shingles", phrase).
				Fuzziness("AUTO").
				MaxExpansions(30).
				PrefixLength(phraseNoFuzzy).
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
			Size(6).
			Do(context.TODO())
	} else {
		results, err = s.se.C().Search().
			Index(searchengine.PODCAST_INDEX).
			Query(elastic.NewTermQuery("title.prefix", prefix)).
			Highlight(elastic.NewHighlight().
				FragmentSize(10).
				PreTags("<em>").
				PostTags("</em>").
				Fields(elastic.NewHighlighterField("text")),
			).
			Size(6).
			Do(context.TODO())
	}

	if err != nil {
		return nil, model.New500Error("search_layer.search_podcast_store.get_typeahead_suggestions", err.Error(), nil)
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

func (s *searchPodcastStore) Search(query string, offset, limit int) ([]*model.Podcast, *model.AppError) {
	results, err := s.se.C().Search().
		Index(searchengine.PODCAST_INDEX).
		Query(elastic.NewMultiMatchQuery(query).
			Type("phrase").
			FieldWithBoost("title", 1.5).
			Field("author").
			TieBreaker(0.4),
		).
		From(offset).
		Size(limit).
		Do(context.TODO())

	if err != nil {
		return nil, model.New500Error("search_layer.search_podcast_store.search", err.Error(), nil)
	}
	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.Podcast{}, nil
	}

	res := []*model.Podcast{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.Podcast{}
		if err := tmp.LoadFromSearchHit(hit); err == nil {
			res = append(res, tmp)
		}
	}
	return res, nil
}
