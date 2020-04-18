package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch_"
	"github.com/varmamsp/cello/store_"
)

const (
	EPISODE_INDEX         = "episode"
	EPISODE_INDEX_MAPPING = `{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 2
		},
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword",
					"index": false
				},
				"podcast_id": {
					"type": "keyword"
				},
				"title": {
					"type": "text"
				},
				"description": {
					"type": "text"
				},
				"pub_date": {
					"type": "date",
					"format": "yyyy-MM-dd HH:mm:ss"
				},
				"duration": {
					"type": "short"
				},
				"type": {
					"type": "keyword"
				}
			}
		}
	}`
)

type searchEpisodeStore struct {
	store_.EpisodeStore
	es elasticsearch_.Broker
}

func newSearchEpisodeStore(baseStore store_.EpisodeStore, broker elasticsearch_.Broker) (store_.EpisodeStore, *model.AppError) {
	if err := broker.CreateIndex(EPISODE_INDEX, EPISODE_INDEX_MAPPING); err != nil {
		return nil, err
	}

	return &searchEpisodeStore{EpisodeStore: baseStore, es: broker}, nil
}

func (s *searchEpisodeStore) Save(episode *model.Episode) *model.AppError {
	if err := s.EpisodeStore.Save(episode); err != nil {
		return err
	}
	return s.es.IndexDoc(EPISODE_INDEX, episode.ForIndexing())
}

func (s *searchEpisodeStore) SyncIndex() *model.AppError {
	if err := s.es.DeleteIndex(EPISODE_INDEX); err != nil {
		return err
	}
	if err := s.es.CreateIndex(EPISODE_INDEX, EPISODE_INDEX_MAPPING); err != nil {
		return err
	}

	lastId, limit := int64(0), 10000
	for {
		episodes, err := s.EpisodeStore.GetAllPaginated(lastId, limit)
		if err != nil {
			return err
		}

		m := make([]model.EsModel, len(episodes))
		for i, episode := range episodes {
			m[i] = episode.ForIndexing()
		}

		if err := s.es.BulkIndexDoc(EPISODE_INDEX, m); err != nil {
			return err
		}

		if len(episodes) < limit {
			break
		}
		lastId = episodes[len(episodes)-1].Id
	}

	return nil
}
