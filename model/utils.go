package model

import (
	"strconv"
	"strings"
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
