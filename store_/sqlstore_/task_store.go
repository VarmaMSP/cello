package sqlstore_

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlTaskStore struct {
	sqldb.Broker
}

func (s *sqlTaskStore) GetAll() (res []*model.Task, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM task WHERE active = 1`,
		cols(&model.Task{}),
	)
	copyTo := func() []interface{} {
		tmp := &model.Task{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_task_store.get_all_active", err.Error(), nil)
	}
	return
}

func (s *sqlTaskStore) Update(old *model.Task, new *model.Task) *model.AppError {
	panic("not implemented") // TODO: Implement
}
