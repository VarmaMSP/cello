package model

import (
	"strconv"
	"strings"
)

// ParseTime is a helper method to get seconds from time string (HH:MM:SS / MM:SS / SS).
func ParseTime(timeString string) int {
	var x []int
	for _, v := range strings.Split(timeString, ":") {
		i, _ := strconv.Atoi(v)
		x = append(x, i)
	}
	timeInSec := 0
	for i, s := len(x)-1, 1; i >= 0; i, s = i-1, s*60 {
		timeInSec = timeInSec + x[i]*s
	}
	return timeInSec
}
