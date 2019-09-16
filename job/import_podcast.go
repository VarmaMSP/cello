package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"

	h "github.com/go-http-utils/headers"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type ImportPodcastJob struct {
	*app.App
	input            <-chan amqp.Delivery
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP *rabbitmq.Producer
}

func NewImportPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.ImportPodcast.WorkerLimit

	importPodcastC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_IMPORT_PODCAST,
		ConsumerName:  config.Queues.ImportPodcast.ConsumerName,
		AutoAck:       config.Queues.ImportPodcast.ConsumerAutoAck,
		Exclusive:     config.Queues.ImportPodcast.ConsumerExclusive,
		PreFetchCount: config.Queues.ImportPodcast.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	createThumbnailP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_CREATE_THUMBNAIL,
		DeliveryMode: config.Queues.CreateThumbnail.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}

	return &ImportPodcastJob{
		App:   app,
		input: importPodcastC.D,
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

func (job *ImportPodcastJob) Run() {
	for d := range job.input {
		job.Call(d)
	}
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

		if err := job.savePodcast(feed.Id, rssFeed); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed
		}

		feedU.ETag = headers[h.ETag]
		feedU.LastModified = headers[h.LastModified]
		if feedU.SetRefershInterval(rssFeed); feedU.RefreshEnabled == 1 {
			feedU.NextRefreshAt = int64(feedU.RefreshInterval) + model.Now()
		}

	update_feed:
		feedU.UpdatedAt = model.Now()
		job.Store.Feed().Update(&feed, &feedU)
	}()
}

func (job *ImportPodcastJob) savePodcast(podcastId string, rssFeed *rss.Feed) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import_job.save_podcast",
		http.StatusInternalServerError,
		map[string]string{"title": rssFeed.Title},
	)

	// Save podcast
	podcast := &model.Podcast{Id: podcastId}
	if err := podcast.LoadDetails(rssFeed); err != nil {
		return appErrorC(err.Error())
	}
	if err := job.Store.Podcast().Save(podcast); err != nil {
		return appErrorC(err.Error())
	}

	// Offload creation of podcast thumbnail
	job.createThumbnailP.D <- map[string]string{
		"id":        podcast.Id,
		"image_src": podcast.ImagePath,
		"type":      "PODCAST",
	}

	// Index podcast
	job.ElasticSearch.Index().
		Index(elasticsearch.PodcastIndexName).
		Id(podcast.Id).
		BodyJson(&model.PodcastIndex{
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
		if err := job.Store.Episode().Save(episode); err != nil {
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
			if err := job.Store.Category().SavePodcastCategory(category); err != nil {
				continue
			}
		}
	}

	return nil
}
