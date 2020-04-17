package sqldb

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/varmamsp/cello/model"
)

type mysqlBroker struct {
	db *sql.DB
}

func NewSqlDbBroker(config *model.Config) (Broker, error) {
	db, err := sql.Open("mysql", makeMysqlDSN(config))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &mysqlBroker{db: db}, nil
}

func (m *mysqlBroker) GetMaster() *sql.DB {
	return m.db
}

func makeMysqlDSN(config *model.Config) string {
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
