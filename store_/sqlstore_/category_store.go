package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlCategoryStore struct {
	sqlStore
}

func newSqlCategoryStore(store sqlStore) store_.CategoryStore {
	return &sqlCategoryStore{store}
}

func (s *sqlCategoryStore) Get(categoryId int64) (*model.Category, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlCategoryStore) GetAll() ([]*model.Category, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlCategoryStore) GetByIds(categoryIds []int64) ([]*model.Category, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlCategoryStore) SavePodcastCategory(category *model.PodcastCategory) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlCategoryStore) GetPodcastCategories(podcastId int64) ([]*model.PodcastCategory, *model.AppError) {
	panic("not implemented") // TODO: Implement
}
