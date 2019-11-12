package sqlstore

import (
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

func UpdateQuery(
	tableName string,
	old, new model.DbModel,
	whereClause string,
	values []interface{},
) (sql string, updateValues []interface{}, noChanges bool) {
	cols := old.DbColumns()
	oldValues := ValuesFromAddrs(old.FieldAddrs())
	newValues := ValuesFromAddrs(new.FieldAddrs())

	var updates []string
	for i := 0; i < len(oldValues); i++ {
		if oldValues[i] != newValues[i] {
			updates = append(updates, cols[i]+" = ?")
			updateValues = append(updateValues, newValues[i])
		}
	}

	if len(updateValues) == 0 {
		noChanges = true
		return
	}

	sql = "UPDATE " + tableName + " SET " + strings.Join(updates, ",") + " WHERE " + whereClause
	updateValues = append(updateValues, values...)
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
