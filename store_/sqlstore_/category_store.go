package sqlstore_

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlCategoryStore struct {
	sqldb.Broker
}

func (s *sqlCategoryStore) Get(categoryId int64) (*model.Category, *model.AppError) {
	res := &model.Category{}
	sql := fmt.Sprintf(
		`SELECT %s FROM category WHERE id = %d`,
		cols(res), categoryId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_category_store.get", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlCategoryStore) GetAll() (res []*model.Category, appE *model.AppError) {
	sql := fmt.Sprintf(`SELECT %s FROM category`, cols(&model.Category{}))
	copyTo := func() []interface{} {
		tmp := &model.Category{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_category_store.get_all", err.Error(), nil)
	}
	return
}

func (s *sqlCategoryStore) GetByIds(categoryIds []int64) (res []*model.Category, appE *model.AppError) {
	if len(categoryIds) == 0 {
		return
	}

	sql := fmt.Sprintf(
		`SELECT %s FROM category WHERE id IN (%s)`,
		cols(&model.Category{}), joinInt64s(categoryIds),
	)
	copyTo := func() []interface{} {
		tmp := &model.Category{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_category_store.get_by_ids", err.Error(), nil)
	}
	return
}

func (s *sqlCategoryStore) SavePodcastCategory(pCategory *model.PodcastCategory) *model.AppError {
	pCategory.PreSave()

	if _, err := s.Insert("podcast_category", pCategory); err != nil {
		return model.New500Error("sql_store.sql_category_store.save_podcast_category", err.Error(), nil)
	}
	return nil
}

func (s *sqlCategoryStore) GetPodcastCategories(podcastId int64) (res []*model.PodcastCategory, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast_category WHERE podcast_id = %d`,
		cols(&model.PodcastCategory{}), podcastId,
	)
	copyTo := func() []interface{} {
		tmp := &model.PodcastCategory{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_category_store.get_podcast_categories", err.Error(), nil)
	}
	return
}
