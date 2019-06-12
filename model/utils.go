package model

import (
	"bytes"
	"database/sql"
	"encoding/base32"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pborman/uuid"
)

const (
	MYSQL_DATETIME = "2006-01-02 15:04:05"
)

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	DetailedError string `json:"detailed_error"`
	StatusCode    int    `json:"status_code,omitempty"`
	Where         string `json:"-"`
	params        map[string]interface{}
}

func (e *AppError) Error() string {
	return e.Where + ": " + e.Message + ", " + e.DetailedError
}

func NewAppError(where string, id string, params map[string]interface{}, details string, status int) *AppError {
	return &AppError{id, id, details, status, where, params}
}

type NullInt64 struct {
	sql.NullInt64
}

func (i *NullInt64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		json.Marshal(i.Int64)
	}
	return []byte("null"), nil
}

func (i *NullInt64) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, str); err != nil {
		i.Valid = false
		return err
	}

	val, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		i.Valid = false
		return err
	}

	i.Int64 = val
	return nil
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

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26)
	return b.String()
}
