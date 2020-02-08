package sqlstore

import (
	"fmt"
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

func (s *SqlCategoryStore) GetAll() (res []*model.Category, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM category`,
		joinStrings((&model.Category{}).DbColumns(), ","),
	)

	copyTo := func() []interface{} {
		tmp := &model.Category{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_category_store.get_all", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlCategoryStore) GetByIds(categoryIds []int64) (res []*model.Category, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM category WHERE id IN (%s)`,
		joinStrings((&model.Category{}).DbColumns(), ","), joinInt64s(categoryIds, ","),
	)

	copyTo := func() []interface{} {
		tmp := &model.Category{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_category_store.get_by_ids", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlCategoryStore) GetPodcastCategories(podcastId int64) (res []*model.PodcastCategory, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast_category WHERE podcast_id = %d`,
		joinStrings((&model.PodcastCategory{}).DbColumns(), ","), podcastId,
	)

	copyTo := func() []interface{} {
		tmp := &model.PodcastCategory{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_category_store.get_podcast_categories", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}
