package sqldb

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/varmamsp/cello/model"
)

type supplier struct {
	db *sqlx.DB
	q  goqu.DialectWrapper
}

func NewBroker(config *model.Config) (Broker, error) {
	db, err := sqlx.Connect("mysql", makeMysqlDSN(config))
	if err != nil {
		return nil, err
	}
	return &supplier{db: db, q: goqu.Dialect("mysql")}, nil
}

func (splr *supplier) C() *sqlx.DB {
	return splr.db
}

func (splr *supplier) GetMaster() *sqlx.DB {
	return splr.db
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
