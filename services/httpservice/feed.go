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
	Error        string
	Etag         string
	LastModified string
}

func (client *Client) GetFeed(feedDetails *model.PodcastFeedDetails) *Feed {
	ch := make(chan *Feed)
	go client.makeFeedRequest(feedDetails, ch)
	return <-ch
}

func (client *Client) GetMultipleFeeds(feedDetails []*model.PodcastFeedDetails) []*Feed {
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
	req.Header.Add("If-None-Match", feedDetails.FeedETag)
	req.Header.Add("If-Modified-Since", feedDetails.FeedLastModified)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		ch <- &Feed{
			Id:    feedDetails.Id,
			Error: err.Error(),
		}
		return
	}

	// Not Modified
	if resp.StatusCode == 304 {
		ch <- &Feed{
			Id:           feedDetails.Id,
			Etag:         resp.Header.Get("ETag"),
			LastModified: resp.Header.Get("Last-Modified"),
		}
		return
	}

	// URL Redirection
	if resp.StatusCode == 302 {
		client.makeFeedRequest(
			&model.PodcastFeedDetails{
				Id:               feedDetails.Id,
				FeedUrl:          resp.Header.Get("Location"),
				FeedETag:         feedDetails.FeedETag,
				FeedLastModified: feedDetails.FeedLastModified,
			},
			ch,
		)
		return
	}

	if resp.StatusCode != 200 {
		ch <- &Feed{
			Id:    feedDetails.Id,
			Error: fmt.Sprintf("Request Unsuccessful: %d", resp.StatusCode),
		}
		return
	}

	parser := &rss.Parser{}
	rssFeed, err := parser.Parse(resp.Body)
	if err != nil {
		ch <- &Feed{
			Id:    feedDetails.Id,
			Error: "Error parsing: " + err.Error(),
		}
		return
	}

	ch <- &Feed{
		Id:           feedDetails.Id,
		RssFeed:      rssFeed,
		Etag:         resp.Header.Get("ETag"),
		LastModified: resp.Header.Get("Last-Modified"),
	}
}
