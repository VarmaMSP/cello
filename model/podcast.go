package model

type Podcast struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	Explicit    int    `json:"explicit"`
	Author      string `json:"author,omitempty"`
	Type        string `json:"type"`
}

type PodcastFeedDetails struct {
	FeedUrl              string
	LastModified         string
	ETag                 string
	TotalEpisodeCount    int
	LatestEpisodeGuid    string
	LatestEpisodePubDate string
}

type PodcastAdminDetails struct {
	Link       string
	OwnerName  string
	OwnerEmail string
}
