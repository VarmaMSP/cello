package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

func (m *mysqlBroker) Patch(table string, old, new model.DbModel) (sql.Result, error) {
	panic("")
}
