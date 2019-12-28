package sqlstore

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/varmamsp/cello/model"
)

func InsertQuery(
	tableName string,
	models []model.DbModel,
) (sql string, insertValues []interface{}, noValues bool) {
	if len(models) == 0 {
		noValues = true
		return
	}

	// (col1, col2, col3)
	cols := "(" + Cols(models[0]) + ")"
	// (?, ?, ?), (?, ?, ?)...
	placeholders := strings.Join(
		Replicate(
			"("+strings.Join(Replicate("?", len(models[0].DbColumns())), ",")+")",
			len(models),
		),
		",",
	)

	sql = "INSERT INTO " + tableName + cols + " VALUES " + placeholders
	for i := range models {
		insertValues = append(insertValues, ValuesFromAddrs(models[i].FieldAddrs())...)
	}
	return
}

func InsertQueryWithoutPK(
	tableName string,
	item model.DbModel,
) (sql string, insertValues []interface{}, noValues bool) {
	cols := "(" + strings.Join(item.DbColumns()[1:], ",") + ")"
	placeholders := "(" + strings.Join(Replicate("?", len(item.DbColumns())-1), ",") + ")"

	sql = "INSERT INTO " + tableName + cols + " VALUES " + placeholders
	insertValues = item.FieldAddrs()[1:]
	return
}

func UpdateQuery(
	tableName string,
	old, new model.DbModel,
	whereClause string,
) (sql string, noUpdates bool) {
	cols := old.DbColumns()
	oldValues := ValuesFromAddrs(old.FieldAddrs())
	newValues := ValuesFromAddrs(new.FieldAddrs())

	var updates []string
	for i := 0; i < len(oldValues); i++ {
		if oldValues[i] != newValues[i] {
			updates = append(updates, fmt.Sprintf("%s = %v", cols[i], newValues[i]))
		}
	}

	if len(updates) == 0 {
		noUpdates = true
		return
	}

	sql = fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		tableName, joinStrings(updates, ","), whereClause,
	)
	return
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

func Cols(m model.DbModel, prefix ...string) string {
	cols := m.DbColumns()
	if len(prefix) > 0 {
		for i, col := range cols {
			cols[i] = prefix[0] + "." + col
		}
	}
	return strings.Join(cols, ",")
}

func joinStrings(vals []string, sep string) string {
	return strings.Join(vals, sep)
}

func joinInt64s(vals []int64, sep string) string {
	s := make([]string, len(vals))
	for i, val := range vals {
		s[i] = model.StrFromInt64(val)
	}
	return joinStrings(s, sep)
}

func MakeMysqlDSN(config *model.Config) string {
	c := mysql.NewConfig()
	c.Net = "tcp"
	c.Addr = config.Mysql.Address
	c.DBName = config.Mysql.Database
	c.User = config.Mysql.User
	c.Passwd = config.Mysql.Password
	c.AllowNativePasswords = true
	c.Params = map[string]string{"collation": "utf8mb4_unicode_ci"}
	c.ReadTimeout = 2 * time.Minute
	c.WriteTimeout = 2 * time.Minute

	return c.FormatDSN()
}
