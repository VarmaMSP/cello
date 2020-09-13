package sqldb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/varmamsp/cello/model"
)

func (splr *supplier) Patch(table string, old, new model.DbModel) (sql.Result, error) {
	new.PreSave()

	cols := new.DbColumns()
	oldValues := valuesFromAddrs(old.FieldAddrs())
	newValues := valuesFromAddrs(new.FieldAddrs())

	var updates []string
	var updateValues []interface{}
	for i, col := range cols {
		if oldValues[i] != newValues[i] {
			updates = append(updates, fmt.Sprintf("%s = ?", col))
			updateValues = append(updateValues, newValues[i])
		}
	}

	if len(updateValues) == 0 {
		return nil, nil
	}

	sql := fmt.Sprintf(`UPDATE %s SET %s WHERE id = ?`, table, strings.Join(updates, ","))
	values := append(updateValues, newValues[0])

	return splr.db.Exec(sql, values...)
}

func valuesFromAddrs(addrs []interface{}) []interface{} {
	values := make([]interface{}, len(addrs))
	for i := range values {
		values[i] = reflect.ValueOf(addrs[i]).Elem().Interface()
	}
	return values
}
