package elasticsearch

const (
	KeywordIndexName = "keyword"
	KeywordMapping   = `{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 2,
			"index": {
				"analysis": {
					"analyzer": {
						"shingle_analyzer": {
							"tokenizer": "standard",
							"filter": [
								"standard",
								"lowercase",
								"apostrophe",
								"filter_shingle"
							]
						 },
						 "prefix_analyzer": {
							"tokenizer": "standard",
							"filter": [
								"standard",
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
				"text": {
					"type": "text",
					"analyzer":"shingle_analyzer",
					"fields": {
						"prefix": {
							"type": "text",
							"analyzer": "prefix_analyzer"
						}
					}
				},
				"added_by": {
					"type": "keyword"
				}
			}
		}
	}`

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
