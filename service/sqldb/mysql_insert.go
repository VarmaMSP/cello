package sqldb

import (
	"database/sql"
	"fmt"
)

func (splr *supplier) Insert(table string, item interface{}) (sql.Result, error) {
	sql, _, err := splr.q.Insert(table).Rows(item).ToSQL()
	if err != nil {
		return nil, err
	}
	fmt.Println(sql)
	return splr.db.Exec(sql)
}

func (splr *supplier) BulkInsert(table string, items ...interface{}) (sql.Result, error) {
	sql, _, err := splr.q.Insert(table).Rows(items...).ToSQL()
	if err != nil {
		return nil, err
	}
	fmt.Println(sql)
	return splr.db.Exec(sql)
}
