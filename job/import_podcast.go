package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/varmamsp/gofeed/rss"

	h "github.com/go-http-utils/headers"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type ImportPodcastJob struct {
	*app.App
	log              zerolog.Logger
	input            <-chan amqp.Delivery
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP *rabbitmq.Producer
}

func NewImportPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.ImportPodcast.WorkerLimit

	importPodcastC, err := rabbitmq.NewConsumer(
		app.RabbitmqConsumerConn,
		&rabbitmq.ConsumerOpts{
			QueueName:     rabbitmq.QUEUE_NAME_IMPORT_PODCAST,
			ConsumerName:  config.Queues.ImportPodcast.ConsumerName,
			AutoAck:       config.Queues.ImportPodcast.ConsumerAutoAck,
			Exclusive:     config.Queues.ImportPodcast.ConsumerExclusive,
			PreFetchCount: config.Queues.ImportPodcast.ConsumerPreFetchCount,
		},
	)
	if err != nil {
		return nil, err
	}

	createThumbnailP, err := rabbitmq.NewProducer(
		app.RabbitmqProducerConn,
		&rabbitmq.ProducerOpts{
			ExchangeName: rabbitmq.EXCHANGE_NAME_PHENOPOD_DIRECT,
			RoutingKey:   rabbitmq.ROUTING_KEY_CREATE_THUMBNAIL,
			DeliveryMode: config.Queues.CreateThumbnail.DeliveryMode,
		},
	)
	if err != nil {
		return nil, err
	}

	return &ImportPodcastJob{
		App:   app,
		log:   app.Log.With().Str("job", "import_podcast").Logger(),
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

type EntitiesToSave struct {
	podcast           *model.Podcast
	episodes          []*model.Episode
	podcastCategories []*model.PodcastCategory
}

type EntitiesToIndex struct {
	podcast  *model.Podcast
	episodes []*model.Episode
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

		// Updated feed
		feedU := feed
		feedU.LastRefreshAt = model.Now()

		if rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{}, job.httpClient); err != nil || rssFeed == nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if entitiesToSave, err := job.extract(feed.Id, rssFeed); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if entitiesToIndex, err := job.save(entitiesToSave); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if err := job.index(entitiesToIndex); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = "SUCCESS"
			if feedU.SetRefershInterval(rssFeed); feedU.RefreshEnabled == 1 {
				feedU.NextRefreshAt = int64(feedU.RefreshInterval) + model.Now()
			}
		}

	update_feed:
		feedU.UpdatedAt = model.Now()
		job.Store.Feed().Update(&feed, &feedU)
	}()
}

func (job *ImportPodcastJob) extract(feedId int64, rssFeed *rss.Feed) (*EntitiesToSave, *model.AppError) {
	var episodes []*model.Episode
	for _, item := range rssFeed.Items {
		episode := &model.Episode{PodcastId: feedId}
		if err := episode.LoadDetails(item); err != nil {
			job.log.Err(err)
			continue
		}
		episodes = append(episodes, episode)
	}

	var podcastCategories []*model.PodcastCategory
	if rssFeed.ITunesExt != nil {
		for _, c := range rssFeed.ITunesExt.Categories {
			if model.CategoryId(c.Text) != -1 {
				podcastCategories = append(podcastCategories, &model.PodcastCategory{
					PodcastId:  feedId,
					CategoryId: model.CategoryId(c.Text),
				})
			}
			if c.Subcategory != nil && model.CategoryId(c.Subcategory.Text) != -1 {
				podcastCategories = append(podcastCategories, &model.PodcastCategory{
					PodcastId:  feedId,
					CategoryId: model.CategoryId(c.Subcategory.Text),
				})
			}
		}
	}

	podcast := &model.Podcast{Id: feedId}
	if l := len(episodes); l > 0 {
		podcast.TotalSeasons = episodes[0].Season
		podcast.TotalEpisodes = l
		podcast.LastestEpisodePubDate = episodes[0].PubDate
		podcast.EarliestEpisodePubDate = episodes[l-1].PubDate
	}
	if err := podcast.LoadDetails(rssFeed); err != nil {
		return nil, model.NewAppError(
			"jobs.podcast_import_job.save_podcast", err.Error(), http.StatusBadRequest,
			map[string]interface{}{"title": rssFeed.Title},
		)
	}

	return &EntitiesToSave{podcast, episodes, podcastCategories}, nil
}

func (job *ImportPodcastJob) save(toSave *EntitiesToSave) (*EntitiesToIndex, *model.AppError) {
	if err := job.Store.Podcast().Save(toSave.podcast); err != nil {
		return nil, err
	}

	for _, podcastCategory := range toSave.podcastCategories {
		if err := job.Store.Category().SavePodcastCategory(podcastCategory); err != nil {
			job.log.Err(err)
			continue
		}
	}

	episodesToIndex := []*model.Episode{}
	for _, episode := range toSave.episodes {
		if err := job.Store.Episode().Save(episode); err != nil {
			job.log.Err(err)
			continue
		} else {
			episodesToIndex = append(episodesToIndex, episode)
		}
	}

	// Create thumbnail
	job.createThumbnailP.D <- CreateThumbnailJobInput{
		Id:         toSave.podcast.Id,
		Type:       "PODCAST",
		ImageSrc:   toSave.podcast.ImagePath,
		ImageTitle: model.UrlParamFromId(toSave.podcast.Title, toSave.podcast.Id),
	}

	return &EntitiesToIndex{toSave.podcast, episodesToIndex}, nil
}

func (job *ImportPodcastJob) index(toIndex *EntitiesToIndex) *model.AppError {
	podcast := toIndex.podcast
	episodes := toIndex.episodes
	indexRequests := make([]elastic.BulkableRequest, len(episodes)+1)

	indexRequests[0] = elastic.NewBulkIndexRequest().
		Index(elasticsearch.PodcastIndexName).
		Id(model.StrFromInt64(podcast.Id)).
		Doc(&model.PodcastIndex{
			Id:          podcast.Id,
			Title:       podcast.Title,
			Author:      podcast.Author,
			Description: podcast.Description,
			Type:        podcast.Type,
			Complete:    podcast.Complete,
		})

	for i, episode := range episodes {
		indexRequests[i+1] = elastic.NewBulkIndexRequest().
			Index(elasticsearch.EpisodeIndexName).
			Id(model.StrFromInt64(episode.Id)).
			Doc(&model.EpisodeIndex{
				Id:          episode.Id,
				PodcastId:   episode.PodcastId,
				Title:       episode.Title,
				Description: episode.Description,
				PubDate:     episode.PubDate,
				Duration:    episode.Duration,
				Type:        episode.Type,
			})
	}

	bulkIndexSize := 20

	for i := 0; i < len(indexRequests); i += bulkIndexSize {
		end := i + bulkIndexSize
		if end > len(indexRequests) {
			end = len(indexRequests)
		}

		if _, err := job.ElasticSearch.Bulk().Add(indexRequests[i:end]...).Do(context.TODO()); err != nil {
			return model.NewAppError("jobs.podcast_import_job.save_podcast", err.Error(), http.StatusBadRequest, nil)
		}
	}

	return nil
}
