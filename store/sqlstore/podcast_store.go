package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastStore struct {
	SqlStore
}

func (s *SqlPodcastStore) Save(podcast *model.Podcast) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		res, err := s.Insert([]DbModel{podcast}, "podcast")
		if err != nil {
			r.Err = model.NewAppError(
				"PodcastStore.Save",
				"store.sqlstore.sql_podcast_store.save",
				map[string]interface{}{
					"title": podcast.Title,
				},
				"Cannot save podcast",
				http.StatusInternalServerError,
			)
			return
		}
		r.Data = res
	})
}

func NewSqlPodcastStore(store SqlStore) store.PodcastStore {
	return &SqlPodcastStore{store}
}
