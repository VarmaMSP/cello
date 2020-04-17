package sqlstore_

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/varmamsp/cello/model"
)

// cols returns comma separated column names for given model.
func cols(m model.DbModel) string {
	return strings.Join(m.DbColumns(), ",")
}

// cols_ returns comma separated column names (excluding primary key name) for given model.
func cols_(m model.DbModel) string {
	return strings.Join(m.DbColumns()[1:], ",")
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

// vals_ returns comma separated values (excluding primary key value) for given model.
func vals_(m model.DbModel) string {
	addrs := m.FieldAddrs()[1:]

	s := make([]string, len(addrs))
	for i, addr := range addrs {
		s[i] = formatToSqlValue(reflect.ValueOf(addr).Elem().Interface())
	}
	return strings.Join(s, ",")
}

// ph returns enough comma separated placeholders for values of given model.
func ph(m model.DbModel) string {
	n := len(m.FieldAddrs())

	s := make([]string, n)
	for i := range s {
		s[i] = "?"
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
