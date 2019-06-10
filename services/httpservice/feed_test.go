package httpservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varmamsp/cello/model"
)

func TestGetFeeds(t *testing.T) {
	assert := assert.New(t)

	client := NewClient()
	feeds := client.GetFeeds(
		[]*model.PodcastFeedDetails{
			&model.PodcastFeedDetails{Id: 20000, FeedUrl: "https://feeds.megaphone.fm/thechernobylpodcast"},
			&model.PodcastFeedDetails{Id: 20001, FeedUrl: "http://feeds.feedburner.com/dancarlin/history?format=xml"},
			&model.PodcastFeedDetails{Id: 20002, FeedUrl: ""},
		},
	)

	assert.Equal(feeds[0].Id, int64(20000))
	assert.Equal(feeds[0].RssFeed.Title, "The Chernobyl Podcast")

	assert.Equal(feeds[1].Id, int64(20001))
	assert.Equal(feeds[1].RssFeed.Title, "Dan Carlin's Hardcore History")

	assert.Equal(feeds[2].Id, int64(20002))
	assert.Nil(feeds[2].RssFeed)
}
