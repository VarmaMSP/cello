package jobserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/varmamsp/gofeed/rss"

	h "github.com/go-http-utils/headers"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type ImportPodcastJob struct {
	store            store.Store
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP messagequeue.Producer
	input            messagequeue.Consumer
}

func NewImportPodcastJob(store store.Store, mq messagequeue.Broker, config *model.Config) (Job, error) {
	importPodcastC, err := mq.NewConsumer(
		messagequeue.QUEUE_IMPORT_PODCAST,
		config.Queues.ImportPodcast.ConsumerName,
		config.Queues.ImportPodcast.ConsumerAutoAck,
		config.Queues.ImportPodcast.ConsumerExclusive,
		config.Queues.ImportPodcast.ConsumerPreFetchCount,
	)
	if err != nil {
		return nil, err
	}

	createThumbnailP, err := mq.NewProducer(
		messagequeue.EXCHANGE_PHENOPOD_DIRECT,
		messagequeue.ROUTING_KEY_CREATE_THUMBNAIL,
		config.Queues.CreateThumbnail.DeliveryMode,
	)
	if err != nil {
		return nil, err
	}

	workerLimit := config.Jobs.ImportPodcast.WorkerLimit
	return &ImportPodcastJob{
		store: store,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter:      make(chan struct{}, workerLimit),
		createThumbnailP: createThumbnailP,
		input:            importPodcastC,
	}, nil
}

func (job *ImportPodcastJob) Start() {
	job.input.Consume(job.Call)
}

type EntitiesToSave struct {
	podcast           *model.Podcast
	episodes          []*model.Episode
	podcastCategories []*model.PodcastCategory
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

		now := datetime.Unix()

		// Updated feed
		feedU := feed

		if rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{}, job.httpClient); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else if entitiesToSave, err := job.extract(feed.Id, rssFeed); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else if err := job.save(entitiesToSave); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = "SUCCESS"
			if feedU.SetRefershInterval(rssFeed); feedU.RefreshEnabled == 1 {
				feedU.NextRefreshAt = int64(feedU.RefreshInterval) + datetime.Unix()
			}
		}

	update_feed:
		feedU.LastRefreshAt = now
		feedU.UpdatedAt = now
		if err := job.store.Feed().Update(&feed, &feedU); err != nil {

		}
	}()
}

func (job *ImportPodcastJob) extract(feedId int64, rssFeed *rss.Feed) (*EntitiesToSave, *model.AppError) {
	// Episodes
	var episodes []*model.Episode
	for _, item := range rssFeed.Items {
		episode := &model.Episode{PodcastId: feedId}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		episodes = append(episodes, episode)
	}

	// Categories
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

	// Podcast
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

	return &EntitiesToSave{
		podcast:           podcast,
		episodes:          episodes,
		podcastCategories: podcastCategories,
	}, nil
}

func (job *ImportPodcastJob) save(toSave *EntitiesToSave) *model.AppError {
	// Podcast
	if err := job.store.Podcast().Save(toSave.podcast); err != nil {
		return err
	}

	// Categories
	for _, podcastCategory := range toSave.podcastCategories {
		if err := job.store.Category().SavePodcastCategory(podcastCategory); err != nil {
			continue
		}
	}

	// Episodes
	episodesToIndex := []*model.Episode{}
	for _, episode := range toSave.episodes {
		if err := job.store.Episode().Save(episode); err != nil {
			continue
		}
		episodesToIndex = append(episodesToIndex, episode)
	}

	// Create thumbnail
	job.createThumbnailP.Publish(CreateThumbnailJobInput{
		Id:         toSave.podcast.Id,
		Type:       "PODCAST",
		ImageSrc:   toSave.podcast.ImagePath,
		ImageTitle: hashid.UrlParam(toSave.podcast.Title, toSave.podcast.Id),
	})

	return nil
}
