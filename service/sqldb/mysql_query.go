package sqldb

func (splr *supplier) Query(copyTo func() []interface{}, sql string, values ...interface{}) error {
	panic("")
}

func (splr *supplier) QueryRow(copyTo []interface{}, sql string, values ...interface{}) error {
	panic("")
}
