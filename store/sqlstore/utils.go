package sqlstore

import (
	"reflect"
	"strings"
)

func InsertQuery(tableName string, model DbModel, count int) string {
	cols := model.DbColumns()
	cols_ := strings.Join(cols, ",")
	placeholder_ := "(" + strings.Join(Replicate("?", len(cols)), ",") + ")"
	placeholders_ := strings.Join(Replicate(placeholder_, count), ",")

	return "INSERT INTO " + tableName + cols_ + " VALUES " + placeholders_
}

func Replicate(s string, n int) []string {
	x := make([]string, n)
	for i := range x {
		x[i] = s
	}
	return x
}

func ValuesFromAddrs(addrs []interface{}) []interface{} {
	values := make([]interface{}, len(addrs))
	for i := range values {
		values[i] = reflect.ValueOf(addrs[i]).Elem()
	}
	return values
}
