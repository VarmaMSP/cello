package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlTaskStore struct {
	sqlStore
}

func newSqlTaskStore(store sqlStore) store_.TaskStore {
	return &sqlTaskStore{store}
}

func (s *sqlTaskStore) GetAll() ([]*model.Task, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlTaskStore) Update(old *model.Task, new *model.Task) *model.AppError {
	panic("not implemented") // TODO: Implement
}
