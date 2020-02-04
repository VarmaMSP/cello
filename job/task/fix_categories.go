package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

type FixCategories struct {
	*app.App
	httpClient *http.Client
}

func NewFixCategories(app *app.App) (*FixCategories, error) {
	return &FixCategories{
		App: app,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 25,
			},
		},
	}, nil
}

func (s *FixCategories) Call() {
	fmt.Println("Fix categories started")
	go func() {
		limit := 1000
		lastId := int64(0)
		rateLimiter := make(chan struct{}, 100)

		for {
			feeds, err := s.Store.Feed().GetAllPaginated(lastId, limit)
			if err != nil {
				break
			}

			for _, feed := range feeds {
				rateLimiter <- struct{}{}

				go func(feed *model.Feed) {
					defer func() { <-rateLimiter }()

					rssFeed, _, err := fetchRssFeed(feed.Url, map[string]string{}, s.httpClient)
					if err != nil {
						return
					}

					var podcastCategories []*model.PodcastCategory
					if rssFeed.ITunesExt != nil {
						for _, c := range rssFeed.ITunesExt.Categories {
							if model.CategoryId(c.Text) != -1 {
								podcastCategories = append(podcastCategories, &model.PodcastCategory{
									PodcastId:  feed.Id,
									CategoryId: model.CategoryId(c.Text),
								})
							}
							if c.Subcategory != nil && model.CategoryId(c.Subcategory.Text) != -1 {
								podcastCategories = append(podcastCategories, &model.PodcastCategory{
									PodcastId:  feed.Id,
									CategoryId: model.CategoryId(c.Subcategory.Text),
								})
							}
						}
					}

					for _, podcastCategory := range podcastCategories {
						if err := s.Store.Category().SavePodcastCategory(podcastCategory); err != nil {
							return
						}
					}
				}(feed)
			}

			if len(feeds) < limit {
				break
			}
			lastId = feeds[len(feeds)-1].Id
		}
	}()
}
