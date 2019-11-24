package sqlstore

import (
	"net/http"

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

	if _, err := s.Insert("podcast_category", []model.DbModel{category}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_category_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": category.PodcastId, "category_id": category.CategoryId},
		)
	}
	return nil
}
