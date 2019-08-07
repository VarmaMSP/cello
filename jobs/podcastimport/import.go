package podcastimport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"
	"github.com/streadway/amqp"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
)

type PodcastImport struct {
	store       store.Store
	consumer    *rabbitmq.Consumer
	httpClient  *http.Client
	workerLimit int
}

func New(store store.Store, consumer *rabbitmq.Consumer, workerLimit int) (*PodcastImport, error) {
	return &PodcastImport{
		store:    store,
		consumer: consumer,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        2 * workerLimit,
				MaxIdleConnsPerHost: 2 * workerLimit,
			},
		},
		workerLimit: workerLimit,
	}, nil
}

func (p *PodcastImport) Run() {
	semaphore := make(chan int, p.workerLimit)

	for {
		semaphore <- 0

		go func(d amqp.Delivery) {
			defer d.Ack(false)
			defer func() { <-semaphore }()

			msg := model.MapFromJson(d.Body)
			itunesId, feedUrl := msg["id"], msg["feed_url"]
			if itunesId == "" || feedUrl == "" {
				return
			}

			if status, _ := p.store.ItunesMeta().GetStatus(itunesId); status != model.StatusPending {
				return
			}

			newStatus := model.StatusSuccess
			if err := p.AddToDb(feedUrl); err != nil {
				newStatus = model.StatusFailure
			}
			p.store.ItunesMeta().SetStatus(itunesId, newStatus)
		}(<-p.consumer.D)
	}
}

func (p *PodcastImport) AddToDb(feedUrl string) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import.add_to_db",
		http.StatusInternalServerError,
		map[string]string{"feed_url": feedUrl},
	)

	// fetch rss feed
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return appErrorC(err.Error())
	}
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return appErrorC(err.Error())
	}

	// parse rss feed
	parser := &rss.Parser{}
	feed, err := parser.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return appErrorC(fmt.Sprintf("Cannot parse feed: %s", err.Error()))
	}

	// Save Podcast
	podcast := &model.Podcast{
		Id:               xid.New().String(),
		FeedUrl:          feedUrl,
		FeedETag:         resp.Header.Get("ETag"),
		FeedLastModified: resp.Header.Get("Last-Modified"),
	}
	if err := podcast.LoadDetails(feed); err != nil {
		return err
	}
	if err := p.store.Podcast().Save(podcast); err != nil {
		return err
	}

	// Save Episodes
	for _, item := range feed.Items {
		episode := &model.Episode{
			Id:        xid.New().String(),
			PodcastId: podcast.Id,
		}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		if err := p.store.Episode().Save(episode); err != nil {
			fmt.Printf("%s %s: %s\n", feedUrl, episode.Title, err.Error())
		}
	}

	// Save Categories
	var categoryIds []int
	if feed.ITunesExt != nil {
		for _, c := range feed.ITunesExt.Categories {
			if c.Subcategory != nil {
				categoryIds = append(categoryIds, model.CategoryId(c.Subcategory.Text))
			}
			categoryIds = append(categoryIds, model.CategoryId(c.Text))
		}
	}
	for _, categoryId := range categoryIds {
		if categoryId != -1 {
			p.store.Category().SavePodcastCategory(&model.PodcastCategory{
				PodcastId:  podcast.Id,
				CategoryId: categoryId,
			})
		}
	}

	return nil
}
