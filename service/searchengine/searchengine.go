package searchengine

import "github.com/varmamsp/cello/model"

type Broker interface {
	Podcast() PodcastIndex
	Episode() EpisodeIndex
}

type PodcastIndex interface {
	CreateIndex() *model.AppError
	DeleteIndex() *model.AppError

	Save(podcast *model.Podcast) *model.AppError
	BulkSave(podcasts []*model.Podcast) *model.AppError
	Delete(podcastId int64) *model.AppError

	Search(query string) ([]*model.Podcast, *model.AppError)
	GetSuggestions(query string) ([]*model.Podcast, *model.AppError)
}

type EpisodeIndex interface {
	CreateIndex() *model.AppError
	DeleteIndex() *model.AppError

	Save(episode *model.Episode) *model.AppError
	BulkSave(episodes []*model.Episode) *model.AppError

	Search(query string) ([]*model.Episode, *model.AppError)
}
