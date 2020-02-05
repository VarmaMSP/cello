package elasticsearch

const (
	PodcastIndexName = "podcast"
	PodcastMapping   = `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 1
		},
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword",
					"index": false
				},
				"title": {
					"type": "search_as_you_type"
				},
				"author": {
					"type": "search_as_you_type"
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

	EpisodeIndexName = "episode"
	EpisodeMapping   = `{
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
