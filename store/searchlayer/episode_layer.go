package searchlayer

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type searchEpisodeStore struct {
	store.EpisodeStore
	se searchengine.Broker
}

func (s *searchEpisodeStore) Save(episode *model.Episode) *model.AppError {
	if err := s.EpisodeStore.Save(episode); err != nil {
		return err
	}
	if err := s.se.Index(searchengine.EPISODE_INDEX, episode.ForIndexing()); err != nil {
		return model.New500Error("search_layer.search_episode_store.save", err.Error(), nil)
	}
	return nil
}

func (s *searchEpisodeStore) Search(query, sortBy string, offset, limit int) ([]*model.Episode, *model.AppError) {
	q := s.se.C().Search().
		Index(searchengine.EPISODE_INDEX).
		Query(elastic.NewMultiMatchQuery(query).
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
		return nil, model.New500Error("search_layer.search_episode_store.search", err.Error(), nil)
	}
	return episodesFromSearch(results)
}

func (s *searchEpisodeStore) SearchByPodcast(podcastId int64, query string, offset, limit int) ([]*model.Episode, *model.AppError) {
	results, err := s.se.C().Search().
		Index(searchengine.EPISODE_INDEX).
		Query(elastic.NewBoolQuery().
			Should(
				elastic.NewMultiMatchQuery(query).
					FieldWithBoost("title", 1.5).
					Field("description").
					Fuzziness("AUTO").
					TieBreaker(0.3),
			).
			Filter(elastic.NewTermQuery("podcast_id", podcastId)).
			MinimumNumberShouldMatch(1),
		).
		Highlight(elastic.NewHighlight().
			FragmentSize(200).
			NumOfFragments(0).
			PreTags("<span class=\"result-highlight\">").
			PostTags("</span>").
			Fields(
				elastic.NewHighlighterField("title"),
			),
		).
		From(offset).
		Size(limit).
		Do(context.TODO())

	if err != nil {
		return nil, model.New500Error("search_layer.search_episode_store.search_by_podcast", err.Error(), nil)
	}
	return episodesFromSearch(results)
}

func episodesFromSearch(results *elastic.SearchResult) ([]*model.Episode, *model.AppError) {
	if results.Hits == nil || results.Hits.Hits == nil || len(results.Hits.Hits) == 0 {
		return []*model.Episode{}, nil
	}

	res := []*model.Episode{}
	for _, hit := range results.Hits.Hits {
		tmp := &model.Episode{}
		if err := tmp.LoadFromSearchHit(hit); err == nil {
			res = append(res, tmp)
		}
	}
	return res, nil
}
