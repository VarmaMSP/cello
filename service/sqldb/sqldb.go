package sqldb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
)

// Broker implements some common helper functions while exposing
// the underling Mysql handle.
type Broker interface {
	// C returns underlying db handle.
	C() *sqlx.DB

	// Returns master connection
	GetMaster() *sqlx.DB

	// Insert model
	Insert(table string, item interface{}) (sql.Result, error)
	// Bulk Insert
	BulkInsert(table string, items ...interface{}) (sql.Result, error)

	// Patch persists changes to item
	Patch(table string, old, new model.DbModel) (sql.Result, error)

	// Exec executes given sql
	Exec(stmt *sqlf.Stmt) error
	// Exec raw sql query
	ExecRaw(sql string, values ...interface{}) error

	// Runs a query and copy results to desctination
	Query(dest interface{}, stmt *sqlf.Stmt) error
	// Runs a query that returns a single row and copy it to destination
	QueryRow(dest interface{}, stmt *sqlf.Stmt) error
}
