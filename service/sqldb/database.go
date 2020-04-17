package sqldb

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
)

type Broker interface {
	// GetMaster returns a connectio to mater db.
	GetMaster() *sql.DB

	// Insert runs insert query for table without a autso generated PK.
	Insert(table string, item model.DbModel) (*sql.Result, error)
	// Insert runs insert query for table with a auto increment PK.
	Insert_(table string, item model.DbModel) (*sql.Result, error)
	// BulkInsert inserts multiple items in single query
	BulkInsert(table string, items []model.DbModel) (*sql.Result, error)

	// Patch persists changes to item
	Patch(table string, old, new model.DbModel) (sql.Result, error)

	// Query runs given sql and copies each result to copyTo func result
	Query(copyTo func() []interface{}, sql string, values ...interface{}) error
	// Query runs given sql and copies it to copyTo addresses
	QueryRow(copyTo []interface{}, sql string, values ...interface{}) error
}
