package model

import (
	"encoding/json"
	"io"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/varmamsp/gofeed/rss"
)

const (
	MAX_SUMMARY_SIZE = 230

	COMMENT_INVALID_STATUS_CODE    = "INVALID_STATUS_CODE"
	COMMENT_INVALID_CONTENT_TYPE   = "INVALID_CONTENT_TYPE"
	COMMENT_UNABLE_TO_MAKE_REQUEST = "UNABLE_TO_MAKE_REQUEST"

	secondsInHour  = 60 * 60
	secondsInDay   = 60 * 60 * 24
	secondsInWeek  = 60 * 60 * 24 * 7
	secondsInMonth = 60 * 60 * 24 * 30
	secondsInYear  = 60 * 60 * 24 * 365
)

var (
	regexpHtmlCharacterEntity = regexp.MustCompile(`&[a-zA-Z];`)
)

type DbModel interface {
	PreSave()
	DbColumns() []string
	FieldAddrs() []interface{}
}

type EsModel interface {
	GetId() string
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

// IsValidKeyword checks if a keyword from podcast description is valid
func IsValidKeyword(keyword string) bool {
	words := strings.Split(keyword, " ")

	// Check for number of words
	if len(words) < 2 {
		return false
	}

	// Check for single character words
	for _, word := range words {
		if len([]rune(word)) < 2 {
			return false
		}
	}

	// Check for special characters
	for _, r := range []rune(keyword) {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
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

// Parse Category Url param
func ParseCategoryUrlParam(urlParam string) (string, int64, error) {
	// if urlParam == "" {
	// 	return "", 0, errors.New("UrlParam is empty")
	// }

	// x := strings.Split(urlParam, "-")
	// if len(x) < 2 || len(x[len(x)-1]) < MIN_HASH_ID_LENGTH {
	// 	return "", 0, errors.New("UrlParam is invalid")
	// }

	// id, err := Int64FromHashId(x[len(x)-1])
	// if err != nil {
	// 	return "", 0, err
	// }

	// return strings.Join(x[0:len(x)-1], "-"), id, nil
	panic("")
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
	return regexpHtmlCharacterEntity.ReplaceAllString(
		strip.StripTags(str),
		" ",
	)
}

func IsContentTypeFeed(contentType string) bool {
	for _, v := range strings.Split(contentType, ";") {
		t := strings.TrimSpace(v)

		if t == "text/xml" {
			return true
		}
		if t == "application/xml" {
			return true
		}
		if t == "application/rss+xml" {
			return true
		}
	}
	return false
}
