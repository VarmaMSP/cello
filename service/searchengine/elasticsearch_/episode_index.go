package elasticsearch_

import (
	"context"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
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

type esEpisodeIndex struct {
	esBroker
}

func newESEpisodeIndex(broker esBroker) (searchengine.EpisodeIndex, error) {
	if err := broker.createIndex(EPISODE_INDEX, EPISODE_INDEX_MAPPING); err != nil {
		return nil, err
	}

	return &esEpisodeIndex{broker}, nil
}

func (e *esEpisodeIndex) CreateIndex() *model.AppError {
	result, err := e.getClient().CreateIndex(EPISODE_INDEX).
		Body(EPISODE_INDEX_MAPPING).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("create_index", err.Error(), http.StatusInternalServerError, nil)
	} else if result == nil {
		return model.NewAppError("create_index", "result is nil", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esEpisodeIndex) DeleteIndex() *model.AppError {
	result, err := e.getClient().DeleteIndex(EPISODE_INDEX).
		Do(context.TODO())

	if err != nil {
		return model.NewAppError("create_index", err.Error(), http.StatusInternalServerError, nil)
	} else if result == nil {
		return model.NewAppError("create_index", "result is nil", http.StatusInternalServerError, nil)
	}

	return nil
}

func (e *esEpisodeIndex) Save(episode *model.Episode) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (e *esEpisodeIndex) BulkSave(episodes []*model.Episode) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (e *esEpisodeIndex) Search(query string) ([]*model.Episode, *model.AppError) {
	return nil, nil
}
