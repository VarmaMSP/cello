package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlTaskStore struct {
	sqldb.Broker
}

func (s *sqlTaskStore) GetAll() ([]*model.Task, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlTaskStore) Update(old *model.Task, new *model.Task) *model.AppError {
	panic("not implemented") // TODO: Implement
}
