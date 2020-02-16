package elasticsearch

const (
	PodcastIndexName = "podcast"
	PodcastMapping   = `{
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
