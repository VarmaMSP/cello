package sqldb

func (m *mysqlBroker) Query(copyTo func() []interface{}, sql string, values ...interface{}) error {
	panic("")
}

func (m *mysqlBroker) QueryRow(copyTo []interface{}, sql string, values ...interface{}) error {
	panic("")
}
