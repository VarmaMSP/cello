package model

import (
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	StatusSuccess = "SUCCESS"
	StatusFailure = "FAILURE"
	StatusPending = "PENDING"

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

// Parse time string (HH:MM:SS / MM:SS / SS) to seconds.
func ParseTime(timeString string) int {
	x := strings.Split(timeString, ":")
	sec := 0
	for i, s := len(x)-1, 1; i >= 0; i, s = i-1, s*60 {
		t, _ := strconv.Atoi(x[i])
		sec = sec + t*s
	}
	return sec
}

// Current unix timestamp
func Now() int64 {
	return time.Now().UTC().Unix()
}

// Current Mysql datetime
func NowDateTime() string {
	return time.Now().UTC().Format(MYSQL_DATETIME)
}

var (
	regexpUrlWithQuery    = regexp.MustCompile(`(https?:\/\/.+)\?.*`)
	regexpUrlWithFragment = regexp.MustCompile(`(https?:\/\/.+)#.*`)
)

func RemoveQueryFromUrl(rawUrl string) string {
	if regexpUrlWithQuery.MatchString(rawUrl) {
		capture := regexpUrlWithQuery.FindStringSubmatch(rawUrl)
		return capture[1]
	}
	return rawUrl
}

func RemoveFragmentFromUrl(rawUrl string) string {
	if regexpUrlWithFragment.MatchString(rawUrl) {
		capture := regexpUrlWithFragment.FindStringSubmatch(rawUrl)
		return capture[1]
	}
	return rawUrl
}

func IsValidHttpUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "http://") != 0 && strings.Index(rawUrl, "https://") != 0 {
		return false
	}
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}
	return true
}

func IsValidEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}

func IsValidAudioType(audioType string) bool {
	if ok, err := regexp.MatchString(`(?:audio|video)\/*.`, audioType); err == nil && ok {
		return true
	}
	if audioType == "application/pdf" {
		return true
	}
	return true
}
