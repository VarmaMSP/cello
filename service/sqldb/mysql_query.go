package sqldb

func (splr *supplier) Query(copyTo func() []interface{}, sql string, values ...interface{}) error {
	rows, err := splr.db.Query(sql, values...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(copyTo()...); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (splr *supplier) QueryRow(copyTo []interface{}, sql string, values ...interface{}) error {
	return splr.db.QueryRow(sql, values...).Scan(copyTo...)
}
