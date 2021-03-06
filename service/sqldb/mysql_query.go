package sqldb

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/leporo/sqlf"
)

func (splr *supplier) Query(dest interface{}, stmt *sqlf.Stmt) error {
	sql, err := dbr.InterpolateForDialect(stmt.String(), stmt.Args(), dialect.MySQL)
	if err != nil {
		return err
	}
	fmt.Println(sql)
	return splr.db.Select(dest, sql)
}

func (splr *supplier) QueryRow(dest interface{}, stmt *sqlf.Stmt) error {
	sql, err := dbr.InterpolateForDialect(stmt.String(), stmt.Args(), dialect.MySQL)
	if err != nil {
		return err
	}
	fmt.Println(sql)
	return splr.db.Get(dest, sql)
}
