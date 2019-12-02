package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/varmamsp/gofeed/rss"

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
		QueueName:     rabbitmq.QUEUE_NAME_IMPORT_PODCAST,
		ConsumerName:  config.Queues.ImportPodcast.ConsumerName,
		AutoAck:       config.Queues.ImportPodcast.ConsumerAutoAck,
		Exclusive:     config.Queues.ImportPodcast.ConsumerExclusive,
		PreFetchCount: config.Queues.ImportPodcast.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	createThumbnailP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.EXCHANGE_NAME_PHENOPOD_DIRECT,
		RoutingKey:   rabbitmq.ROUTING_KEY_CREATE_THUMBNAIL,
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

func (job *ImportPodcastJob) savePodcast(podcastId int64, rssFeed *rss.Feed) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import_job.save_podcast",
		http.StatusInternalServerError,
		map[string]interface{}{"title": rssFeed.Title},
	)

	// Episodes
	var episodes []*model.Episode
	for _, item := range rssFeed.Items {
		episode := &model.Episode{PodcastId: podcastId}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		episodes = append(episodes, episode)
		if err := job.Store.Episode().Save(episode); err != nil {
			continue
		}
	}

	// Podcast Categories
	var podcastCategories []*model.PodcastCategory
	if rssFeed.ITunesExt != nil {
		for _, c := range rssFeed.ITunesExt.Categories {
			if model.CategoryId(c.Text) != -1 {
				podcastCategories = append(podcastCategories, &model.PodcastCategory{
					PodcastId:  podcastId,
					CategoryId: model.CategoryId(c.Text),
				})
			}
			if c.Subcategory != nil && model.CategoryId(c.Subcategory.Text) != -1 {
				podcastCategories = append(podcastCategories, &model.PodcastCategory{
					PodcastId:  podcastId,
					CategoryId: model.CategoryId(c.Subcategory.Text),
				})
			}
		}
	}

	// Podcast
	podcast := &model.Podcast{Id: podcastId}
	if l := len(episodes); l > 0 {
		podcast.TotalSeasons = episodes[0].Season
		podcast.TotalEpisodes = l
		podcast.LastestEpisodePubDate = episodes[0].PubDate
		podcast.EarliestEpisodePubDate = episodes[l-1].PubDate
	}
	if err := podcast.LoadDetails(rssFeed); err != nil {
		return appErrorC(err.Error())
	}

	// Persist
	if err := job.Store.Podcast().Save(podcast); err != nil {
		return appErrorC(err.Error())
	}
	for _, episode := range episodes {
		if err := job.Store.Episode().Save(episode); err != nil {
			continue
		}
	}
	for _, podcastCategory := range podcastCategories {
		if err := job.Store.Category().SavePodcastCategory(podcastCategory); err != nil {
			continue
		}
	}

	podcastHashId := model.HashIdFromInt64(podcast.Id)

	// Create thumbnail
	job.createThumbnailP.D <- map[string]interface{}{
		"id":          podcast.Id,
		"type":        "PODCAST",
		"image_src":   podcast.ImagePath,
		"image_title": podcastHashId,
	}

	// Index Podcast
	job.ElasticSearch.Index().
		Index(elasticsearch.PodcastIndexName).
		Id(podcastHashId).
		BodyJson(&model.PodcastIndex{
			Id:          podcastHashId,
			Title:       model.UrlParamFromId(podcast.Title, podcast.Id),
			Author:      podcast.Author,
			Description: podcast.Description,
			Type:        podcast.Type,
			Complete:    podcast.Complete,
		}).
		Do(context.TODO())

	return nil
}
