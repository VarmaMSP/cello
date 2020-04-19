package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

// Broker implements some common helper functions while exposing
// the underling Mysql handle.
type Broker interface {
	// C returns underlying db handle.
	C() *sql.DB

	// Insert runs insert query for table without a autso generated PK.
	Insert(table string, item model.DbModel) (sql.Result, error)
	// Insert runs insert query for table with a auto increment PK.
	Insert_(table string, item model.DbModel) (sql.Result, error)
	// BulkInsert inserts multiple items in single query
	BulkInsert(table string, items []model.DbModel) (sql.Result, error)
	// Patch persists changes to item
	Patch(table string, old, new model.DbModel) (sql.Result, error)
	// Exec executes given sql
	Exec(sql string, values ...interface{}) error

	// Query runs given sql and copies each result to copyTo func result
	Query(copyTo func() []interface{}, sql string, values ...interface{}) error
	// Query runs given sql and copies it to copyTo addresses
	QueryRow(copyTo []interface{}, sql string, values ...interface{}) error
}
