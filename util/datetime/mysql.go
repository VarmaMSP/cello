package datetime

import "time"

const _MYSQL_DATETIME = "2006-01-02 15:04:05"

// NowDateTime returns UTC time in Mysql datetime format
func Now() string {
	return time.Now().UTC().Format(_MYSQL_DATETIME)
}

// TimeFromDateTime return time object from mysql datetime string
func ToTime(s string) *time.Time {
	if res, err := time.Parse(_MYSQL_DATETIME, s); err == nil {
		return &res
	}
	return nil
}

// DateTimeFromTime formats time to mysql datetime
func FromTime(t *time.Time) string {
	return t.UTC().Format(_MYSQL_DATETIME)
}

func FromUnix(u int64) string {
	return TimeFromUnix(u).UTC().Format(_MYSQL_DATETIME)
}
