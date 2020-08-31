package sqlstore

import (
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlTaskStore struct {
	sqldb.Broker
}

func (s *sqlTaskStore) GetAll() ([]*model.Task, *model.AppError) {
	query := sqlf.
		Select("*").
		From("task").
		Where("active = 1")

	var tasks []*model.Task
	if err := s.Query_(&tasks, query); err != nil {
		return nil, model.New500Error("sql_store.sql_task_store.get_all_active", err.Error(), nil)
	}
	return tasks, nil
}

func (s *sqlTaskStore) Update(old *model.Task, new *model.Task) *model.AppError {
	if _, err := s.Patch("task", old, new); err != nil {
		return model.New500Error("sql_store.sql_task_store.get_all_active", err.Error(), nil)
	}
	return nil
}
