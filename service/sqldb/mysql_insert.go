package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

func (splr *supplier) Insert(table string, item model.DbModel) (*sql.Result, error) {
	panic("")
}

func (splr *supplier) Insert_(table string, item model.DbModel) (*sql.Result, error) {
	panic("")
}

func (splr *supplier) BulkInsert(table string, items []model.DbModel) (*sql.Result, error) {
	panic("")
}
