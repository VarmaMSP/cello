package datetime

import "time"

const _MYSQL_DATETIME = "2006-01-02 15:04:05"

// Now returns MYSQL datetime (UTC)
func Now() string {
	return time.Now().UTC().Format(_MYSQL_DATETIME)
}

// ToTime converts a MYSQL datetime to time
func ToTime(s string) *time.Time {
	if res, err := time.Parse(_MYSQL_DATETIME, s); err == nil {
		return &res
	}
	return nil
}

// FromTime returns MYSQL datetime from time
func FromTime(t *time.Time) string {
	return t.UTC().Format(_MYSQL_DATETIME)
}

// FromUnix returns MYSQL datetime from unix timestamp
func FromUnix(u int64) string {
	return TimeFromUnix(u).UTC().Format(_MYSQL_DATETIME)
}
