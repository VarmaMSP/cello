package elasticsearch_

import (
	"context"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
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

type esPodcastIndex struct {
	esBroker
}

func newESPodcastIndex(broker esBroker) (searchengine.PodcastIndex, error) {
	if err := broker.createIndex(PODCAST_INDEX, PODCAST_INDEX_MAPPING); err != nil {
		return nil, err
	}

	return &esPodcastIndex{broker}, nil
}

func (e *esPodcastIndex) CreateIndex() *model.AppError {
	createIndexResult, err := e.getClient().CreateIndex(PODCAST_INDEX).
		Body(PODCAST_INDEX_MAPPING).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("create_index", err.Error(), http.StatusInternalServerError, nil)
	} else if createIndexResult == nil {
		return model.NewAppError("create_index", "no result", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esPodcastIndex) DeleteIndex() *model.AppError {
	deleteIndexResult, err := e.getClient().DeleteIndex(PODCAST_INDEX).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("create_index", err.Error(), http.StatusInternalServerError, nil)
	} else if deleteIndexResult == nil {
		return model.NewAppError("create_index", "no result", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esPodcastIndex) Save(podcast *model.Podcast) *model.AppError {
	indexResult, err := e.getClient().Index().
		Index(PODCAST_INDEX).
		Id(model.StrFromInt64(podcast.Id)).
		BodyJson(&model.PodcastIndex{
			Id:          podcast.Id,
			Title:       podcast.Title,
			Author:      podcast.Author,
			Description: podcast.Description,
			Type:        podcast.Type,
			Complete:    podcast.Complete,
		}).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("esPodcastIndex", err.Error(), http.StatusInternalServerError, nil)
	} else if indexResult == nil {
		return model.NewAppError("esPodcastIndex", "index result is nil", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esPodcastIndex) BulkSave(podcasts []*model.Podcast) *model.AppError {
	if len(podcasts) == 0 {
		return nil
	}

	indexRequests := make([]elastic.BulkableRequest, len(podcasts))
	for i, podcast := range podcasts {
		indexRequests[i] = elastic.NewBulkIndexRequest().
			Index(PODCAST_INDEX).
			Id(model.StrFromInt64(podcast.Id)).
			Doc(&model.PodcastIndex{
				Id:          podcast.Id,
				Title:       podcast.Title,
				Author:      podcast.Author,
				Description: podcast.Description,
				Type:        podcast.Type,
				Complete:    podcast.Complete,
			})
	}

	bulkIndexSize := 100
	for i := 0; i < len(indexRequests); i += bulkIndexSize {
		end := i + bulkIndexSize
		if end > len(indexRequests) {
			end = len(indexRequests)
		}

		bulkResult, err := e.getClient().Bulk().Add(indexRequests[i:end]...).Do(context.TODO())
		if err != nil {
			return model.NewAppError("bulkindex", err.Error(), http.StatusInternalServerError, nil)
		} else if bulkResult == nil {
			return model.NewAppError("bulkindex", "bulk result is empty", http.StatusInternalServerError, nil)
		}
	}

	return nil
}

func (e *esPodcastIndex) Delete(podcastId int64) *model.AppError {
	deleteResult, err := e.getClient().Delete().
		Index(PODCAST_INDEX).
		Id(model.StrFromInt64(podcastId)).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("delete", err.Error(), http.StatusInternalServerError, nil)
	} else if deleteResult == nil {
		return model.NewAppError("delete", "result is empty", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esPodcastIndex) Search(query string) ([]*model.Podcast, *model.AppError) {
	return nil, nil
}

func (e *esPodcastIndex) GetSuggestions(query string) ([]*model.Podcast, *model.AppError) {
	return nil, nil
}
