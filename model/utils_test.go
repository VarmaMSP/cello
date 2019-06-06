package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5000, ParseTime("5000"))
	assert.Equal(90, ParseTime("1:30"))
	assert.Equal(90, ParseTime("01:30"))
	assert.Equal(90, ParseTime("00:01:30"))
	assert.Equal(3630, ParseTime("01:00:30"))
	assert.Equal(3630, ParseTime("1:0:30"))
}
