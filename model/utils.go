package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/speps/go-hashids"
	"github.com/varmamsp/gofeed/rss"
)

const (
	MYSQL_DATETIME = "2006-01-02 15:04:05"

	MAX_SUMMARY_SIZE = 230

	MIN_HASH_ID_LENGTH = 6

	secondsInHour  = 60 * 60
	secondsInDay   = 60 * 60 * 24
	secondsInWeek  = 60 * 60 * 24 * 7
	secondsInMonth = 60 * 60 * 24 * 30
	secondsInYear  = 60 * 60 * 24 * 365
)

var (
	regexpNbsp = regexp.MustCompile(`&nbsp;`)
)

type DbModel interface {
	PreSave()
	DbColumns() []string
	FieldAddrs() []interface{}
}

type AppError struct {
	Id            string                 `json:"id"`          // Function at which the error occured
	DetailedError string                 `json:"error"`       // Internal Error string
	StatusCode    int                    `json:"status_code"` // Http status code
	Params        map[string]interface{} `json:"parmas"`
}

func (e *AppError) Error() string {
	return e.Id + ": " + e.DetailedError
}

func NewAppError(where string, details string, statusCode int, params map[string]interface{}) *AppError {
	return &AppError{where, details, statusCode, params}
}

func NewAppErrorC(where string, statusCode int, params map[string]interface{}) func(details string) *AppError {
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

// SecondsSince returns number of seconds elapsed since t
func SecondsSince(t *time.Time) int {
	return int(time.Since(*t).Seconds())
}

// NowDateTime returns UTC time in Mysql datetime format
func NowDateTime() string {
	return time.Now().UTC().Format(MYSQL_DATETIME)
}

// ParseDateTime parses mysql date time string
func ParseDateTime(s string) *time.Time {
	if res, err := time.Parse(MYSQL_DATETIME, s); err == nil {
		return &res
	}
	return nil
}

// FormatDateTime formats time to mysql datetime
func FormatDateTime(t *time.Time) string {
	return t.UTC().Format(MYSQL_DATETIME)
}

// TimeFromTimestamp converts timestamp to time
func TimeFromTimestamp(t int64) *time.Time {
	res := time.Unix(t, 0)
	return &res
}

// DateTimeFromTimestamp converts timestamp to mysql time stamp
func DateTimeFromTimestamp(t int64) string {
	return TimeFromTimestamp(t).UTC().Format(MYSQL_DATETIME)
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

// BoolFromStr decodes bool from string
func BoolFromStr(str string) bool {
	return strings.ToLower(str) == "true"
}

// IntFromStr decodes int from string
func IntFromStr(str string) int {
	if res, err := strconv.Atoi(str); err == nil {
		return res
	}
	return 0
}

// Int64FromStr decodes int64 from string
func Int64FromStr(str string) int64 {
	if res, err := strconv.ParseInt(str, 10, 64); err == nil {
		return res
	}
	return 0
}

func StrFromInt64(i int64) string {
	return strconv.FormatInt(i, 10)
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

func MapToJson(data map[string]string) []byte {
	if res, err := json.Marshal(data); err == nil {
		return res
	}
	return []byte{}
}

//EncodeToJson will encode given map to valid json
func EncodeToJson(data interface{}) []byte {
	if res, err := json.Marshal(data); err == nil {
		return res
	}
	return []byte{}
}

// MinInt returns minimum of two integers
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var hashid, _ = hashids.NewWithData(&hashids.HashIDData{
	Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	MinLength: 6,
})

// HashIdFromInt64 Encodes hashId
func HashIdFromInt64(val int64) string {
	hid, _ := hashid.EncodeInt64(([]int64{val}))
	return hid
}

// Int64FromHashId Decodes hashId
func Int64FromHashId(h string) (int64, error) {
	if h == "" {
		return 0, errors.New("HashId is empty")
	}

	res, err := hashid.DecodeInt64WithError(h)
	if err != nil {
		return 0, err
	}
	if len(res) != 1 {
		return 0, errors.New("Hashid invalid")
	}
	return res[0], nil
}

// UrlParamFromId returns urlparam
func UrlParamFromId(title string, id int64) string {
	var sb strings.Builder

	wordCount, maxWordCount := 0, 10
	runeCount, maxRuneCount := 0, 300
	lastChar, hyphen := rune('-'), rune('-')
	for _, r := range []rune(title) {
		if runeCount == maxRuneCount || wordCount == maxWordCount {
			break
		}
		// replace space in title with hyphen while making sure
		// consequent hyphens do not occur
		if unicode.IsSpace(r) {
			if wordCount == maxWordCount-1 {
				break
			}
			if lastChar != hyphen {
				sb.WriteRune(hyphen)
				wordCount += 1
				runeCount += 1
			}
			lastChar = rune('-')
			continue
		}
		// retain hyphen from title while making sure
		// consequent hyphens do not occur
		if r == hyphen {
			if lastChar != hyphen {
				sb.WriteRune(hyphen)
				runeCount += 1
				lastChar = rune(hyphen)
			}
			continue
		}
		// retain all language alphabet and numbers from title
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(unicode.ToLower(r))
			runeCount += 1
			lastChar = r
			continue
		}
	}

	return fmt.Sprintf(
		"%s-%s",
		sb.String(), HashIdFromInt64(id),
	)
}

// IdFromUrlParam return integer from urlparam's hashId
func IdFromUrlParam(urlParam string) (int64, error) {
	if urlParam == "" {
		return 0, errors.New("UrlParam is empty")
	}

	x := strings.Split(urlParam, "-")
	if len(x) < 2 || len(x[len(x)-1]) < MIN_HASH_ID_LENGTH {
		return 0, errors.New("UrlParam is invalid")
	}

	return Int64FromHashId(x[len(x)-1])
}

// RemoveDuplicatesInt64 removes duplicates from []int64
func RemoveDuplicatesInt64(arr []int64) []int64 {
	m := make(map[int64]struct{}, len(arr))
	res := []int64{}
	for _, a := range arr {
		if _, ok := m[a]; !ok {
			res = append(res, a)
		}
		m[a] = struct{}{}
	}
	return res
}

// StripHTMLTags removes all HTML elements from given string
func StripHTMLTags(str string) string {
	return regexpNbsp.ReplaceAllString(
		strip.StripTags(str),
		" ",
	)
}
