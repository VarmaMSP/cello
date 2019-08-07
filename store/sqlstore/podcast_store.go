package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastStore struct {
	SqlStore
}

func NewSqlPodcastStore(store SqlStore) store.PodcastStore {
	return &SqlPodcastStore{store}
}

func (s *SqlPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	_, err := s.Insert([]DbModel{podcast}, "podcast")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"title": podcast.Title, "feed_url": podcast.FeedUrl},
		)
	}
	return nil
}
