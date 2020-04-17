package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

func (m *mysqlBroker) Insert(table string, item model.DbModel) (*sql.Result, error) {
	panic("")
}

func (m *mysqlBroker) Insert_(table string, item model.DbModel) (*sql.Result, error) {
	panic("")
}

func (m *mysqlBroker) BulkInsert(table string, items []model.DbModel) (*sql.Result, error) {
	panic("")
}
