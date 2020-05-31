package job

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/streadway/amqp"

	h "github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
	"github.com/varmamsp/gofeed/rss"
)

type RefreshPodcastJob struct {
	store            store.Store
	log              zerolog.Logger
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP messagequeue.Producer
	input            messagequeue.Consumer
}

func NewRefreshPodcastJob(store store.Store, mq messagequeue.Broker, log zerolog.Logger, config *model.Config) (Job, error) {
	refreshPodcastC, err := mq.NewConsumer(
		messagequeue.QUEUE_REFRESH_PODCAST,
		config.Queues.RefreshPodcast.ConsumerName,
		config.Queues.RefreshPodcast.ConsumerAutoAck,
		config.Queues.RefreshPodcast.ConsumerExclusive,
		config.Queues.RefreshPodcast.ConsumerPreFetchCount,
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

	workerLimit := config.Jobs.RefreshPodcast.WorkerLimit
	return &RefreshPodcastJob{
		store: store,
		log:   log.With().Str("ctx", "job_server.refresh_podcast").Logger(),
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter:      make(chan struct{}, workerLimit),
		createThumbnailP: createThumbnailP,
		input:            refreshPodcastC,
	}, nil
}

func (job *RefreshPodcastJob) Start() {
	job.log.Info().Msg("started")
	go func() {
		d := job.input.Consume()
		for {
			job.Call(<-d)
		}
	}()
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
		job.log.Error().Msg(err.Error())
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

		if rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{h.ETag: feed.ETag, h.LastModified: feed.LastModified}, job.httpClient); err != nil {
			job.log.Error().Msg(err.Error())
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
			job.log.Error().Msg(err.Error())

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
		feedU.LastRefreshAt = now
		feedU.UpdatedAt = now
		if err := job.store.Feed().Update(&feed, &feedU); err != nil {
			job.log.Error().Msg(err.Error())
		}
	}()
}

func (job *RefreshPodcastJob) getRefreshData(feedId int64, rssFeed *rss.Feed) (*RefreshData, *model.AppError) {
	podcast, err := job.store.Podcast().Get(feedId)
	if err != nil {
		return nil, err
	}

	episodes, err := job.store.Episode().GetByPodcast(podcast.Id)
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

		if err := job.store.Episode().Block(episodeIds); err != nil {
			return err
		}
	}

	// Save Episodes
	if len(episodesToAdd) > 0 {
		for _, episode := range episodesToAdd {
			if err := job.store.Episode().Save(episode); err != nil {
				job.log.Error().Msg(err.Error())
				continue
			}
		}
	}

	// Update thumbnail
	if podcast.ImagePath != podcastU.ImagePath {
		if err := job.createThumbnailP.Publish(&CreateThumbnailJobInput{
			Id:         podcast.Id,
			Type:       "PODCAST",
			ImageSrc:   podcast.ImagePath,
			ImageTitle: hashid.UrlParam(podcast.Title, podcast.Id),
		}); err != nil {
			job.log.Error().Msg(err.Error())
		}
	}

	// Update Stats
	if len(episodesToAdd) > 0 {
		if podcastU.TotalSeasons < episodesToAdd[0].Season {
			podcastU.TotalSeasons = episodesToAdd[0].Season
		}
		podcastU.LastestEpisodePubDate = episodesToAdd[0].PubDate
	}

	if err := job.store.Podcast().Update(podcast, podcastU); err != nil {
		job.log.Error().Msg(err.Error())
	}

	return nil
}
