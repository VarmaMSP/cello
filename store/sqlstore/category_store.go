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

func (s *SqlCategoryStore) SavePodcastCategory(category *model.PodcastCategory) *model.AppError {
	_, err := s.Insert([]DbModel{category}, "podcast_category")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_category_store.save_all",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"podcast_id": category.PodcastId, "category_id": strconv.FormatInt(int64(category.CategoryId), 10)},
		)
	}
	return nil
}
