package model

// Elasticsearch podcast index type
type PodcastIndex struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Complete    int    `json:"complete,omitempty"`
}
