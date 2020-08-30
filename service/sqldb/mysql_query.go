package sqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/leporo/sqlf"
)

func (splr *supplier) Query(copyTo func() []interface{}, sql string, values ...interface{}) error {
	rows, err := splr.db.Query(sql, values...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(copyTo()...); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (splr *supplier) QueryRow(copyTo []interface{}, sql string, values ...interface{}) error {
	return splr.db.QueryRow(sql, values...).Scan(copyTo...)
}

type QueryOption int

const (
	ExpandVars QueryOption = iota
)

func (splr *supplier) Query_(dest interface{}, stmt *sqlf.Stmt, options ...QueryOption) error {
	sql := stmt.String()
	args := stmt.Args()

	if doesOptionExist(ExpandVars, options) {
		var err error
		if sql, args, err = sqlx.In(sql, args); err == nil {
			return err
		}
	}

	return splr.db.Select(dest, sql, args...)
}

func (splr *supplier) QueryRow_(dest interface{}, stmt *sqlf.Stmt, options ...QueryOption) error {
	sql := stmt.String()
	args := stmt.Args()

	if doesOptionExist(ExpandVars, options) {
		var err error
		if sql, args, err = sqlx.In(sql, args); err == nil {
			return err
		}
	}

	return splr.db.Get(dest, sql, args...)
}

func doesOptionExist(option QueryOption, options []QueryOption) bool {
	for _, o := range options {
		if o == option {
			return true
		}
	}
	return false
}
