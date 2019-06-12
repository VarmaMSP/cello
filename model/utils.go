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
	Id            string            `json:"id"`          // Function at which the error occured
	DetailedError string            `json:"error"`       // Internal Error string
	StatusCode    int               `json:"status_code"` // Http status code
	params        map[string]string `json:"parmas"`
}

func (e *AppError) Error() string {
	return e.Id + ": " + e.DetailedError
}

func NewAppError(where string, details string, statusCode int, params map[string]string) *AppError {
	return &AppError{where, details, statusCode, params}
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
