package sqldb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/varmamsp/cello/model"
)

func (splr *supplier) Insert(table string, item model.DbModel) (sql.Result, error) {
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(item.DbColumns(), ","),
		strings.Join(replicate("?", len(item.DbColumns())), ","),
	)

	return splr.db.Exec(sql, getValues(item.FieldAddrs())...)
}

func (splr *supplier) Insert_(table string, item model.DbModel) (sql.Result, error) {
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(item.DbColumns()[1:], ","),
		strings.Join(replicate("?", len(item.DbColumns())-1), ","),
	)

	return splr.db.Exec(sql, getValues(item.FieldAddrs()[1:])...)
}

func (splr *supplier) BulkInsert(table string, items []model.DbModel) (sql.Result, error) {
	panic("Not Implemented")
}

func replicate(s string, n int) []string {
	x := make([]string, n)
	for i := range x {
		x[i] = s
	}
	return x
}

func getValues(addrs []interface{}) []interface{} {
	values := make([]interface{}, len(addrs))
	for i := range values {
		values[i] = reflect.ValueOf(addrs[i]).Elem().Interface()
	}
	return values
}
