package model

import (
	"strconv"
	"strings"
)

// Function to get seconds from time string (HH:MM:SS / MM:SS / SS).
func ParseTime(timeString string) int {
	x := strings.Split(timeString, ":")
	sec := 0
	for i, s := len(x)-1, 1; i >= 0; i, s = i-1, s*60 {
		t, _ := strconv.Atoi(x[i])
		sec = sec + t*s
	}
	return sec
}
