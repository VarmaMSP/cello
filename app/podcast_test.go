package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewPodcast(t *testing.T) {
	assert := assert.New(t)

	app := NewApp()
	err := app.AddNewPodcast("https://wtfmnightshyamalan.podbean.com/feed.xml")

	assert.Nil(err)
}
