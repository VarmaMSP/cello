package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"
	"github.com/olivere/elastic/v7"

	h "github.com/go-http-utils/headers"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/elasticsearch"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
)

type ImportPodcastJob struct {
	store            store.Store
	esClient         *elastic.Client
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP *rabbitmq.Producer
}

func NewImportPodcastJob(store store.Store, esClient *elastic.Client, createThumbnailP *rabbitmq.Producer, workerLimit int) (model.Job, error) {
	return &ImportPodcastJob{
		store:    store,
		esClient: esClient,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter:      make(chan struct{}, workerLimit),
		createThumbnailP: createThumbnailP,
	}, nil
}

func (job *ImportPodcastJob) Call(delivery amqp.Delivery) {
	var feed model.Feed
	if err := json.Unmarshal(delivery.Body, &feed); err != nil {
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}

	go func() {
		defer delivery.Ack(false)
		defer func() { <-job.rateLimiter }()

		// updated meta
		feedU := feed
		feedU.LastRefreshAt = model.Now()
		feedU.LastRefreshComment = "SUCCESS"

		rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{}, job.httpClient)
		if err != nil || rssFeed == nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed
		}

		if err := job.savePodcast(rssFeed, feed.Url, headers); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed
		}

		feedU.SetRefershInterval(rssFeed)
		if feedU.RefreshEnabled == 1 {
			feedU.NextRefreshAt = int64(feedU.RefreshInterval) + model.Now()
		}

	update_feed:
		feedU.UpdatedAt = model.Now()
		job.store.Feed().Update(&feed, &feedU)
	}()
}

func (job *ImportPodcastJob) savePodcast(rssFeed *rss.Feed, feedUrl string, headers map[string]string) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import_job.save_podcast",
		http.StatusInternalServerError,
		map[string]string{"title": rssFeed.Title},
	)

	// Save podcast
	podcast := &model.Podcast{
		FeedUrl:          feedUrl,
		FeedETag:         headers[h.ETag],
		FeedLastModified: headers[h.LastModified],
	}
	if err := podcast.LoadDetails(rssFeed); err != nil {
		return appErrorC(err.Error())
	}
	if err := job.store.Podcast().Save(podcast); err != nil {
		return appErrorC(err.Error())
	}

	// Offload creation of podcast thumbnail
	job.createThumbnailP.D <- map[string]string{
		"id":        podcast.Id,
		"image_src": podcast.ImagePath,
		"type":      "PODCAST",
	}

	// Index podcast
	job.esClient.Index().
		Index(elasticsearch.PodcastIndexName).
		Id(podcast.Id).
		BodyJson(&model.PodcastInfo{
			Id:          podcast.Id,
			Title:       podcast.Title,
			Author:      podcast.Author,
			Description: podcast.Description,
			Type:        podcast.Type,
			Complete:    podcast.Complete,
		}).
		Do(context.TODO())

	// Save Episodes
	for _, item := range rssFeed.Items {
		episode := &model.Episode{PodcastId: podcast.Id}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		if err := job.store.Episode().Save(episode); err != nil {
			continue
		}
	}

	// Save Categories
	var categoryIds []int
	if rssFeed.ITunesExt != nil {
		for _, c := range rssFeed.ITunesExt.Categories {
			if c.Subcategory != nil {
				categoryIds = append(categoryIds, model.CategoryId(c.Subcategory.Text))
			}
			categoryIds = append(categoryIds, model.CategoryId(c.Text))
		}
	}
	for _, categoryId := range categoryIds {
		if categoryId != -1 {
			category := &model.PodcastCategory{PodcastId: podcast.Id, CategoryId: categoryId}
			if err := job.store.Category().SavePodcastCategory(category); err != nil {
				continue
			}
		}
	}

	return nil
}
