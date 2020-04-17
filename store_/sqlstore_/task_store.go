package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store_"
)

type sqlTaskStore struct {
	sqldb.Broker
}

func newSqlTaskStore(broker sqldb.Broker) store_.TaskStore {
	return &sqlTaskStore{broker}
}

func (s *sqlTaskStore) GetAll() ([]*model.Task, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlTaskStore) Update(old *model.Task, new *model.Task) *model.AppError {
	panic("not implemented") // TODO: Implement
}
