package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlTaskStore struct {
	SqlStore
}

func NewSqlTaskStore(store SqlStore) store.TaskStore {
	return &SqlTaskStore{store}
}

func (s *SqlTaskStore) GetAllActive() (res []*model.Task, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Task{}) + " FROM task WHERE active = 1"

	copyTo := func() []interface{} {
		tmp := &model.Task{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_task_store.get_all_active", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return
}

func (s *SqlTaskStore) Update(old, new *model.Task) *model.AppError {
	if _, err := s.UpdateChanges("task", old, new, "name = ?", new.Name); err != nil {
		return model.NewAppError(
			"sqlstore.sql_task_store.update", err.Error(), http.StatusInternalServerError,
			map[string]string{"name": new.Name},
		)
	}
	return nil
}