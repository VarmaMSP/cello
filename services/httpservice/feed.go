package httpservice

import (
	"fmt"
	"net/http"

	"github.com/mmcdole/gofeed/rss"
	"github.com/varmamsp/cello/model"
)

type Feed struct {
	Id           int64
	RssFeed      *rss.Feed
	Etag         *string
	LastModified *string
}

func (client *Client) GetFeeds(feedDetails []*model.PodcastFeedDetails) []*Feed {
	n := len(feedDetails)

	ch := make(chan *Feed)
	ind := make(map[int64]int)
	for i := 0; i < n; i++ {
		tmp := feedDetails[i]
		ind[tmp.Id] = i
		go client.makeFeedRequest(tmp, ch)
	}

	results := make([]*Feed, n)
	for i := 0; i < n; i++ {
		tmp := <-ch
		fmt.Println(tmp.Id)
		results[ind[tmp.Id]] = tmp
	}

	return results
}

func (client *Client) makeFeedRequest(feedDetails *model.PodcastFeedDetails, ch chan<- *Feed) {
	req, _ := http.NewRequest("GET", feedDetails.FeedUrl, nil)
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		ch <- &Feed{Id: feedDetails.Id}
		return
	}

	parser := &rss.Parser{}
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		ch <- &Feed{Id: feedDetails.Id}
		return
	}

	result := Feed{Id: feedDetails.Id, RssFeed: feed}
	if tmp := resp.Header.Get("etag"); tmp != "" {
		result.Etag = &tmp
	}
	if tmp := resp.Header.Get("last-modified"); tmp != "" {
		result.LastModified = &tmp
	}
	ch <- &result
}
