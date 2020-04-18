package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch_"
	"github.com/varmamsp/cello/store_"
)

const (
	PODCAST_INDEX         = "podcast"
	PODCAST_INDEX_MAPPING = `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 1,
			"index": {
				"analysis": {
					"analyzer": {	
						"shingle_analyzer": {
							"tokenizer": "standard",
							"filter": [
								"lowercase",
								"filter_shingle"
							]
						 },
						 "prefix_analyzer": {
							"tokenizer": "standard",
							"filter": [
								"lowercase",
								"filter_truncate",
								"filter_edgegram"
							]
						 }
					},
					"filter": {
						"filter_shingle": {
							"type": "shingle",
							"min_shingle_size": 2,
							"max_shingle_size": 4,
							"output_unigrams": "true"
						},
						"filter_truncate": {
							"type": "truncate",
							"length": 8
						},
						"filter_edgegram": {
							"type": "edge_ngram",
							"min_gram": 1,
							"max_gram": 8
						}
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword",
					"index": false
				},
				"title": {
					"type": "text",
					"fields": {
						"shingles": {
							"type": "text",
							"analyzer": "shingle_analyzer"
						},
						"prefix": {
							"type": "text",
							"analyzer": "prefix_analyzer"
						}
					}
				},
				"author": {
					"type": "text",
					"fields": {
						"shingles": {
							"type": "text",
							"analyzer": "shingle_analyzer"
						},
						"prefix": {
							"type": "text",
							"analyzer": "prefix_analyzer"
						}
					}
				},
				"description": {
					"type": "text"
				},
				"type": {
					"type": "keyword"
				},
				"complete": {
					"type": "byte"
				}
			}
		}
	}`
)

type searchPodcastStore struct {
	store_.PodcastStore
	es elasticsearch_.Broker
}

func newSearchPodcastStore(baseStore store_.PodcastStore, broker elasticsearch_.Broker) (store_.PodcastStore, *model.AppError) {
	if err := broker.CreateIndex(PODCAST_INDEX, PODCAST_INDEX_MAPPING); err != nil {
		return nil, err
	}
	return &searchPodcastStore{PodcastStore: baseStore, es: broker}, nil
}

func (s *searchPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	if err := s.PodcastStore.Save(podcast); err != nil {
		return err
	}
	return s.es.IndexDoc(PODCAST_INDEX, podcast.ForIndexing())
}

func (s *searchPodcastStore) SyncIndex() *model.AppError {
	if err := s.es.DeleteIndex(PODCAST_INDEX); err != nil {
		return err
	}
	if err := s.es.CreateIndex(PODCAST_INDEX, PODCAST_INDEX_MAPPING); err != nil {
		return err
	}

	lastId, limit := int64(0), 10000
	for {
		podcasts, err := s.PodcastStore.GetAllPaginated(lastId, limit)
		if err != nil {
			return err
		}

		m := make([]model.EsModel, len(podcasts))
		for i, podcast := range podcasts {
			m[i] = podcast.ForIndexing()
		}

		if err := s.es.BulkIndexDoc(PODCAST_INDEX, m); err != nil {
			return err
		}

		if len(podcasts) < limit {
			break
		}
		lastId = podcasts[len(podcasts)-1].Id
	}

	return nil
}
