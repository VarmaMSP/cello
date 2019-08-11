package sqlstore

import (
	"reflect"
	"strings"
)

func InsertQuery(tableName string, model DbModel, count int) string {
	cols := model.DbColumns()
	cols_ := "(" + strings.Join(cols, ",") + ")"
	placeholder_ := "(" + strings.Join(Replicate("?", len(cols)), ",") + ")"
	placeholders_ := strings.Join(Replicate(placeholder_, count), ",")

	return "INSERT INTO " + tableName + cols_ + " VALUES " + placeholders_
}

func UpdateQuery(tableName string, old, new DbModel) (string, []interface{}) {
	cols := old.DbColumns()
	oldValues := ValuesFromAddrs(old.FieldAddrs())
	newValues := ValuesFromAddrs(new.FieldAddrs())

	var updates []string
	var updateValues []interface{}
	for i := 0; i < len(oldValues); i++ {
		if oldValues[i] != newValues[i] {
			updates = append(updates, cols[i]+" = ?")
			updateValues = append(updateValues, newValues[i])
		}
	}

	return "UPDATE " + tableName + " SET " + strings.Join(updates, ","), updateValues
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
		values[i] = reflect.ValueOf(addrs[i]).Elem().Interface()
	}
	return values
}
