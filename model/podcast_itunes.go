package model

type PodcastItunes struct {
	ItunesId        string
	FeedUrl         string
	AlbumArt        string
	ScrappedAt      string
	AddedToPhenopod int
	AddedAt         string
}

func (pi *PodcastItunes) DbColumns() []string {
	return []string{
		"itunes_id", "feed_url", "album_art", "scrapped_at",
		"added_to_phenopod", "added_at",
	}
}

func (pi *PodcastItunes) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pi.ItunesId, &pi.FeedUrl, &pi.AlbumArt, &pi.ScrappedAt,
		&pi.AddedToPhenopod, &pi.AddedAt,
	)
}
