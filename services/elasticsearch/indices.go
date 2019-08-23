package elasticsearch

const (
	PodcastIndexName = "podcast"
	PodcastMapping   = `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":1
		},
		"mappings":{
			"properties":{
				"id": {
					"type": "keyword"
				},
				"feed_url": {
					"type": "keyword"
				},
				"title": {
					"type": "text"
				},
				"description": {
					"type": "text"
				}
			}
		}
	}`
)
