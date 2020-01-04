package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"

	h "github.com/go-http-utils/headers"
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
	"github.com/varmamsp/cello/service/rabbitmq"
	"github.com/varmamsp/gofeed/rss"
)

type RefreshPodcastJob struct {
	*app.App
	log              zerolog.Logger
	input            <-chan amqp.Delivery
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP *rabbitmq.Producer
}

func NewRefreshPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.RefreshPodcast.WorkerLimit

	refreshPodcastC, err := rabbitmq.NewConsumer(
		app.RabbitmqConsumerConn,
		&rabbitmq.ConsumerOpts{
			QueueName:     rabbitmq.QUEUE_NAME_REFRESH_PODCAST,
			ConsumerName:  config.Queues.RefreshPodcast.ConsumerName,
			AutoAck:       config.Queues.RefreshPodcast.ConsumerAutoAck,
			Exclusive:     config.Queues.RefreshPodcast.ConsumerExclusive,
			PreFetchCount: config.Queues.RefreshPodcast.ConsumerPreFetchCount,
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

	return &RefreshPodcastJob{
		App:   app,
		log:   app.Log.With().Str("job", "refresh_podcast").Logger(),
		input: refreshPodcastC.D,
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

func (job *RefreshPodcastJob) Run() {
	for d := range job.input {
		job.Call(d)
	}
}

type RefreshData struct {
	podcast         *model.Podcast
	podcastU        *model.Podcast
	episodesToAdd   []*model.Episode
	episodesToBlock []*model.Episode
}

func (job *RefreshPodcastJob) Call(delivery amqp.Delivery) {
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

		if rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{h.ETag: feed.ETag, h.LastModified: feed.LastModified}, job.httpClient); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if rssFeed == nil {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = "NOT MODIFIED"
			goto update_feed

		} else if refreshData, err := job.getRefreshData(feed.Id, rssFeed); err != nil {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if err := job.save(refreshData); err != nil {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else if err := job.index(refreshData); err != nil {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = err.Error()
			goto update_feed

		} else {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = "MODIFIED"
		}

	update_feed:
		feedU.UpdatedAt = model.Now()
		job.Store.Feed().Update(&feed, &feedU)
	}()
}

func (job *RefreshPodcastJob) getRefreshData(feedId int64, rssFeed *rss.Feed) (*RefreshData, *model.AppError) {
	podcast, err := job.Store.Podcast().Get(feedId)
	if err != nil {
		return nil, err
	}

	episodes, err := job.Store.Episode().GetByPodcast(podcast.Id)
	if err != nil {
		return nil, err
	}

	podcastU := *podcast
	if err := podcastU.LoadDetails(rssFeed); err != nil {
		return nil, err
	}
	podcastU.TotalEpisodes = len(episodes)

	// Index episodes by their guids
	episodeByGuid := map[string]*model.Episode{}
	for _, episode := range episodes {
		episodeByGuid[episode.Guid] = episode
	}

	// Index rss items by their guids
	rssItemByGuid := map[string]*rss.Item{}
	for _, item := range rssFeed.Items {
		if guid := model.RssItemGuid(item); guid != "" {
			rssItemByGuid[guid] = item
		}
	}

	// New episodes
	episodesToAdd := []*model.Episode{}
	for guid, rssItem := range rssItemByGuid {
		if _, ok := episodeByGuid[guid]; !ok {
			episode := &model.Episode{PodcastId: podcast.Id}
			if err := episode.LoadDetails(rssItem); err != nil {
				job.log.Err(err)
				continue
			}
			episodesToAdd = append(episodesToAdd, episode)
		}
	}

	// Blocked Episodes
	episodesToBlock := []*model.Episode{}
	for guid, episode := range episodeByGuid {
		if _, ok := rssItemByGuid[guid]; !ok {
			episodesToBlock = append(episodesToBlock, episode)
		}
	}

	return &RefreshData{podcast, &podcastU, episodesToAdd, episodesToBlock}, nil
}

func (job *RefreshPodcastJob) save(refreshData *RefreshData) *model.AppError {
	podcast := refreshData.podcast
	podcastU := refreshData.podcastU
	episodesToAdd := refreshData.episodesToAdd
	episodesToBlock := refreshData.episodesToBlock

	// Block Episodes
	if len(episodesToBlock) > 0 {
		episodeIds := make([]int64, len(episodesToBlock))
		for i, episode := range episodesToBlock {
			episodeIds[i] = episode.Id
		}

		if err := job.Store.Episode().Block(episodeIds); err != nil {
			return err
		}
	}

	// Save Episodes
	if len(episodesToAdd) > 0 {
		for _, episode := range episodesToAdd {
			if err := job.Store.Episode().Save(episode); err != nil {
				job.log.Err(err)
				continue
			}
		}
	}

	// Update thumbnail
	if podcast.ImagePath != podcastU.ImagePath {
		job.createThumbnailP.D <- CreateThumbnailJobInput{
			Id:         podcast.Id,
			Type:       "PODCAST",
			ImageSrc:   podcast.ImagePath,
			ImageTitle: model.UrlParamFromId(podcast.Title, podcast.Id),
		}
	}

	// Update Stats
	if len(episodesToAdd) > 0 {
		podcastU.TotalEpisodes = podcastU.TotalEpisodes + len(episodesToAdd) - len(episodesToBlock)
		if podcastU.TotalSeasons < episodesToAdd[0].Season {
			podcastU.TotalSeasons = episodesToAdd[0].Season
		}
		podcastU.LastestEpisodePubDate = episodesToAdd[0].PubDate
	}

	if err := job.Store.Podcast().Update(podcast, podcastU); err != nil {
		job.log.Err(err)
	}

	return nil
}

func (job *RefreshPodcastJob) index(data *RefreshData) *model.AppError {
	podcast := data.podcastU
	episodesToAdd := data.episodesToAdd
	episodesToBlock := data.episodesToBlock
	requests := make([]elastic.BulkableRequest, len(episodesToAdd)+len(episodesToBlock))

	for i, episode := range episodesToAdd {
		requests[i] = elastic.NewBulkIndexRequest().
			Index(elasticsearch.PodcastIndexName).
			Id(model.StrFromInt64(episode.Id)).
			Doc(&model.EpisodeIndex{
				Id:          episode.Id,
				PodcastId:   podcast.Id,
				Title:       episode.Title,
				Description: model.StripHTMLTags(episode.Description),
				PubDate:     episode.PubDate,
				Duration:    episode.Duration,
				Type:        episode.Type,
			})
	}

	for i, episode := range episodesToBlock {
		requests[i+len(episodesToAdd)] = elastic.NewBulkDeleteRequest().
			Index(elasticsearch.EpisodeIndexName).
			Id(model.StrFromInt64(episode.Id))
	}

	bulkIndexSize := 20

	for i := 0; i < len(requests); i += bulkIndexSize {
		end := i + bulkIndexSize
		if end > len(requests) {
			end = len(requests)
		}

		if _, err := job.ElasticSearch.Bulk().Add(requests[i:end]...).Do(context.TODO()); err != nil {
			return model.NewAppError("jobs.podcast_import_job.save_podcast", err.Error(), http.StatusBadRequest, nil)
		}
	}

	return nil
}
