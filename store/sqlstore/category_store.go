package sqlstore

import (
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlCategoryStore struct {
	sqldb.Broker
}

func newSqlCategoryStore(broker sqldb.Broker) *sqlCategoryStore {
	return &sqlCategoryStore{
		Broker: broker,
	}
}

func (s *sqlCategoryStore) Get(categoryId int64) (*model.Category, *model.AppError) {
	query := sqlf.Select("category.*").
		From("category").
		Where("id = ?", categoryId)

	var category model.Category
	if err := s.QueryRow_(&category, query); err != nil {
		return nil, model.New500Error("sql_store.sql_category_store.get", err.Error(), nil)
	}
	return &category, nil
}

func (s *sqlCategoryStore) GetAll() (res []*model.Category, appE *model.AppError) {
	query := sqlf.Select("category.*").
		From("category")

	var categories []*model.Category
	if err := s.Query_(&categories, query); err != nil {
		return nil, model.New500Error("sql_store.sql_category_store.get_all", err.Error(), nil)
	}
	return categories, nil
}

func (s *sqlCategoryStore) GetByIds(categoryIds []int64) ([]*model.Category, *model.AppError) {
	if len(categoryIds) == 0 {
		return []*model.Category{}, nil
	}

	query := sqlf.Select("category.*").
		From("category").
		Where("id IN (?)", categoryIds)

	var categories []*model.Category
	if err := s.Query_(&categories, query, sqldb.ExpandVars); err != nil {
		return nil, model.New500Error("sql_store.sql_category_store.get_by_ids", err.Error(), nil)
	}
	return categories, nil
}

func (s *sqlCategoryStore) SavePodcastCategory(pCategory *model.PodcastCategory) *model.AppError {
	pCategory.PreSave()

	if _, err := s.Insert("podcast_category", pCategory); err != nil {
		return model.New500Error("sql_store.sql_category_store.save_podcast_category", err.Error(), nil)
	}
	return nil
}

func (s *sqlCategoryStore) GetPodcastCategories(podcastId int64) ([]*model.PodcastCategory, *model.AppError) {
	query := sqlf.Select("podcast_category.*").
		From("podcast_category").
		Where("podcast_id = ?", podcastId)

	var podcastCategories []*model.PodcastCategory
	if err := s.Query_(&podcastCategories, query); err != nil {
		return nil, model.New500Error("sql_store.sql_category_store.get_podcast_categories", err.Error(), nil)
	}
	return podcastCategories, nil
}
