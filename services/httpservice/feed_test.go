package httpservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varmamsp/cello/model"
)

func TestGetFeed(t *testing.T) {
	assert := assert.New(t)

	client := NewClient()
	feed1 := client.GetFeed(&model.PodcastFeedDetails{Id: 1})
	feed2 := client.GetFeed(&model.PodcastFeedDetails{Id: 2, FeedUrl: "https://feeds.megaphone.fm/thechernobylpodcast"})

	assert.Equal(feed1.Id, int64(1))
	assert.Nil(feed1.RssFeed)

	assert.Equal(feed2.Id, int64(2))
	assert.Equal(feed2.RssFeed.Title, "The Chernobyl Podcast")
}

func TestGetMultipleFeeds(t *testing.T) {
	assert := assert.New(t)

	client := NewClient()
	feeds := client.GetMultipleFeeds(
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
