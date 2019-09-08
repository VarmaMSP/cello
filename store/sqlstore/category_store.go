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
	category.PreSave()

	_, err := s.Insert("podcast_category", []DbModel{category})
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_category_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"podcast_id": category.PodcastId, "category_id": strconv.FormatInt(int64(category.CategoryId), 10)},
		)
	}
	return nil
}
