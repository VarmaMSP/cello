package sqlstore

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/varmamsp/cello/model"
)

// cols returns comma separated column names for given model.
func cols(m model.DbModel, prefix ...string) string {
	if len(prefix) == 1 {
		return colsWithPrefix(m, prefix[0])
	}
	return strings.Join(m.DbColumns(), ",")
}

func colsWithPrefix(m model.DbModel, prefix string) string {
	columns := m.DbColumns()
	columnsP := make([]string, len(columns))
	for i, column := range columns {
		columnsP[i] = fmt.Sprintf("%s.%s", prefix, column)
	}

	return strings.Join(columnsP, ",")
}

// vals returns comma separated values for given model.
func vals(m model.DbModel) string {
	addrs := m.FieldAddrs()

	s := make([]string, len(addrs))
	for i, addr := range addrs {
		s[i] = formatToSqlValue(reflect.ValueOf(addr).Elem().Interface())
	}
	return strings.Join(s, ",")
}

func joinInt64s(vals []int64) string {
	s := make([]string, len(vals))
	for i, val := range vals {
		s[i] = model.StrFromInt64(val)
	}
	return strings.Join(s, ",")
}

func formatToSqlValue(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("%d", v)

	case int64:
		return fmt.Sprintf("%d", v)

	case string:
		return fmt.Sprintf("'%s'", v)

	default:
		return fmt.Sprintf("%v", v)
	}
}
