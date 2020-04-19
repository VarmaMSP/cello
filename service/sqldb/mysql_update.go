package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

func (splr *supplier) Patch(table string, old, new model.DbModel) (sql.Result, error) {
	panic("")
}

func (splr *supplier) Exec(sql string, values ...interface{}) error {
	_, err := splr.db.Exec(sql, values...)
	return err
}
