package elasticsearch_

import (
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

func (e *esEpisodeIndex) Index(episode *model.Episode) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (e *esEpisodeIndex) BulkIndex(episodes []*model.Episode) *model.AppError {
	panic("not implemented") // TODO: Implement
}
