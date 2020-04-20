package datetime

import "time"

// Now returns unix timestamp
func Unix() int64 {
	return time.Now().UTC().Unix()
}

// TimeFromUnix return time from unix timestamp
func TimeFromUnix(t int64) *time.Time {
	res := time.Unix(t, 0)
	return &res
}

// SecondsSince returns number of seconds elapsed since t
func SecondsSince(t *time.Time) int {
	return int(time.Since(*t).Seconds())
}
