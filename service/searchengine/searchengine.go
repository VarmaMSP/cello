package searchengine

import "github.com/varmamsp/cello/model"

type SearchEngine interface {
	Podcast() PodcastIndex
	Episode() EpisodeIndex
}

type PodcastIndex interface {
	Index(podcast *model.Podcast) *model.AppError
	BulkIndex(podcasts []*model.Podcast) *model.AppError
}

type EpisodeIndex interface {
	Index(episode *model.Episode) *model.AppError
	BulkIndex(episodes []*model.Episode) *model.AppError
}
