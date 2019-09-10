package model

import (
	"encoding/json"
	"io"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mmcdole/gofeed/rss"
)

const (
	StatusSuccess = "SUCCESS"
	StatusFailure = "FAILURE"
	StatusPending = "PENDING"

	StateActive   = "ACTIVE"
	StateInactive = "INACTIVE"

	MYSQL_DATETIME      = "2006-01-02 15:04:05"
	MYSQL_BLOB_MAX_SIZE = 65535
)

type AppError struct {
	Id            string            `json:"id"`          // Function at which the error occured
	DetailedError string            `json:"error"`       // Internal Error string
	StatusCode    int               `json:"status_code"` // Http status code
	Params        map[string]string `json:"parmas"`
}

func (e *AppError) Error() string {
	return e.Id + ": " + e.DetailedError
}

func NewAppError(where string, details string, statusCode int, params map[string]string) *AppError {
	return &AppError{where, details, statusCode, params}
}

func NewAppErrorC(where string, statusCode int, params map[string]string) func(details string) *AppError {
	return func(details string) *AppError {
		return &AppError{where, details, statusCode, params}
	}
}

// RssItemGuid returns guid of the item, Enclosure url is returned if no guid is found
func RssItemGuid(item *rss.Item) string {
	if item.GUID != nil && item.GUID.Value != "" {
		return item.GUID.Value
	}
	if item.Enclosure != nil && item.Enclosure.URL != "" {
		return RemoveQueryFromUrl(item.Enclosure.URL)
	}
	return ""
}

// ParseTime parses number of seconds from a string of HH:MM:SS / MM:SS / SS format
func ParseTime(timeString string) int {
	x := strings.Split(timeString, ":")
	sec := 0
	for i, s := len(x)-1, 1; i >= 0; i, s = i-1, s*60 {
		t, _ := strconv.Atoi(x[i])
		sec = sec + t*s
	}
	return sec
}

// Now returns unix timestamp
func Now() int64 {
	return time.Now().UTC().Unix()
}

// NowDateTime returns UTC time in Mysql datetime format
func NowDateTime() string {
	return time.Now().UTC().Format(MYSQL_DATETIME)
}

// SecondsSince returns number of seconds elapsed since t
func SecondsSince(t *time.Time) int {
	return int(time.Since(*t).Seconds())
}

var (
	regexpUrlWithQuery    = regexp.MustCompile(`(https?:\/\/.+)\?.*`)
	regexpUrlWithFragment = regexp.MustCompile(`(https?:\/\/.+)#.*`)
)

// RemoveQueryFromUrl removes query string from given url if present
func RemoveQueryFromUrl(rawUrl string) string {
	if regexpUrlWithQuery.MatchString(rawUrl) {
		capture := regexpUrlWithQuery.FindStringSubmatch(rawUrl)
		return capture[1]
	}
	return rawUrl
}

// RemoveFragementFromUrl removes fragment from given url if present
func RemoveFragmentFromUrl(rawUrl string) string {
	if regexpUrlWithFragment.MatchString(rawUrl) {
		capture := regexpUrlWithFragment.FindStringSubmatch(rawUrl)
		return capture[1]
	}
	return rawUrl
}

// IsValidHttpUrl validates a url
func IsValidHttpUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "http://") != 0 && strings.Index(rawUrl, "https://") != 0 {
		return false
	}
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}
	return true
}

// IsValidEmail validates a email address
func IsValidEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}

// IsValidMediaType validates media type for rss feed item
func IsValidMediaType(mediaType string) bool {
	if ok, err := regexp.MatchString(`(?:audio|video)\/*.`, mediaType); err == nil && ok {
		return true
	}
	if mediaType == "application/pdf" {
		return true
	}
	return true
}

// MapFromJson will decode a map
func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var m map[string]string
	if err := decoder.Decode(&m); err == nil {
		return m
	}
	return map[string]string{}
}

// MinInt returns minimum of two integers
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
