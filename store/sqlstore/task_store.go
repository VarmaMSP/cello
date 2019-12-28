package sqlstore

import (
	"fmt"
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

func (s *SqlTaskStore) GetAll() (res []*model.Task, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM task WHERE active = 1`,
		joinStrings((&model.Task{}).DbColumns(), ","),
	)

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
	if _, err := s.Update_("task", old, new, fmt.Sprintf("id = %d", new.Id)); err != nil {
		return model.NewAppError(
			"sqlstore.sql_task_store.update", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"name": new.Name},
		)
	}
	return nil
}
