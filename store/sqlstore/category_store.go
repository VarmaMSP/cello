package sqlstore

import (
	"net/http"
	"strconv"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlCategoryStore struct {
	SqlStore
}

func NewSqlCategoryStore(store SqlStore) store.CategoryStore {
	return &SqlCategoryStore{store}
}

func (s *SqlCategoryStore) SavePodcastCategories(podcastCategories []*model.PodcastCategory) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		models := make([]DbModel, len(podcastCategories))
		for i := range models {
			models[i] = podcastCategories[i]
		}

		res, err := s.Insert(models, "podcast_category")
		if err != nil {
			r.Err = model.NewAppError(
				"store.sqlstore.sql_podcast_category_store.save_all",
				err.Error(),
				http.StatusInternalServerError,
				map[string]string{
					"podcast_id": strconv.FormatInt(podcastCategories[0].PodcastId, 10),
				},
			)
		}
		r.Data = res
	})
}
