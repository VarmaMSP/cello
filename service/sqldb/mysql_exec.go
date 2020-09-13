package sqldb

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/leporo/sqlf"
)

func (splr *supplier) Exec(stmt *sqlf.Stmt) error {
	sql, err := dbr.InterpolateForDialect(stmt.String(), stmt.Args(), dialect.MySQL)
	if err != nil {
		return err
	}
	fmt.Println(sql)
	_, err = splr.db.Exec(sql)
	return err
}

func (splr *supplier) ExecRaw(sql_ string, values ...interface{}) error {
	sql, err := dbr.InterpolateForDialect(sql_, values, dialect.MySQL)
	if err != nil {
		return err
	}
	fmt.Println(sql)
	_, err = splr.db.Exec(sql)
	return err
}
