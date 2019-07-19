package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastItunesStore struct {
	SqlStore
}

func NewSqlPodcastItunesStore(store SqlStore) *SqlPodcastItunesStore {
	return &SqlPodcastItunesStore{store}
}

func (s *SqlPodcastItunesStore) SaveAll(podcastItunesMeta []*model.PodcastItunes) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		models := make([]DbModel, len(podcastItunesMeta))
		for i := range models {
			models[i] = podcastItunesMeta[i]
		}

		res, err := s.Insert(models, "podcast_itunes")
		if err != nil {
			r.Err = model.NewAppError(
				"store.sqlstore.sql_podcast_itunes_store.save_all",
				err.Error(),
				http.StatusInternalServerError,
				nil,
			)
		}
		r.Data = res
	})
}
