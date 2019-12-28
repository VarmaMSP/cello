package job

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	h "github.com/go-http-utils/headers"
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
	"github.com/varmamsp/gofeed/rss"
)

type RefreshPodcastJob struct {
	*app.App
	log         zerolog.Logger
	input       <-chan amqp.Delivery
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewRefreshPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.RefreshPodcast.WorkerLimit

	refreshPodcastC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     rabbitmq.QUEUE_NAME_REFRESH_PODCAST,
		ConsumerName:  config.Queues.RefreshPodcast.ConsumerName,
		AutoAck:       config.Queues.RefreshPodcast.ConsumerAutoAck,
		Exclusive:     config.Queues.RefreshPodcast.ConsumerExclusive,
		PreFetchCount: config.Queues.RefreshPodcast.ConsumerPreFetchCount,
	})
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
		rateLimiter: make(chan struct{}, workerLimit),
	}, nil
}

func (job *RefreshPodcastJob) Run() {
	for d := range job.input {
		job.Call(d)
	}
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

		rssFeed, headers, err := fetchRssFeed(
			feed.Url,
			map[string]string{h.ETag: feed.ETag, h.LastModified: feed.LastModified},
			job.httpClient,
		)
		if err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed
		}

		feedU.ETag = headers[h.ETag]
		feedU.LastModified = headers[h.LastModified]
		if rssFeed == nil {
			feedU.LastRefreshComment = "NOT MODIFIED"
			goto update_feed
		}

		if err := job.updateEpisodes(feed.Id, rssFeed); err != nil {
			feedU.LastRefreshComment = err.Error()
			goto update_feed
		}

		feedU.LastRefreshComment = "MODIFIED"

	update_feed:
		feedU.UpdatedAt = model.Now()
		job.Store.Feed().Update(&feed, &feedU)
	}()
}

func (job *RefreshPodcastJob) updateEpisodes(podcastId int64, rssFeed *rss.Feed) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"job.refresh_podcast.update_episodes",
		http.StatusInternalServerError,
		map[string]interface{}{"podcast_id": podcastId},
	)

	episodes, err := job.Store.Episode().GetByPodcast(podcastId)
	if err != nil {
		return appErrorC(err.Error())
	}

	// Index episodes by their guids
	episodeMap := map[string]*model.Episode{}
	for _, episode := range episodes {
		episodeMap[episode.Guid] = episode
	}

	// Index rss items by their guid
	rssItemMap := map[string]*rss.Item{}
	for _, item := range rssFeed.Items {
		if guid := model.RssItemGuid(item); guid != "" {
			rssItemMap[guid] = item
		}
	}

	// blocked episodes
	blockedEpisodesCount := 0
	for episodeGuid, episode := range episodeMap {
		if _, ok := rssItemMap[episodeGuid]; !ok {
			job.Store.Episode().Block(episode.Id)
			blockedEpisodesCount += 1
		}
	}

	// new episodes
	newEpisodesCount := 0
	for rssItemGuid, rssItem := range rssItemMap {
		if _, ok := episodeMap[rssItemGuid]; !ok {
			episode := &model.Episode{PodcastId: podcastId}
			if err := episode.LoadDetails(rssItem); err != nil {
				continue
			}
			if err := job.Store.Episode().Save(episode); err != nil {
				continue
			}
			newEpisodesCount += 1
		}
	}

	if len(rssFeed.Items) > 0 && (newEpisodesCount > 0 || blockedEpisodesCount > 0) {
		latestEpisode := &model.Episode{}
		latestEpisode.LoadDetails(rssFeed.Items[0])

		err := job.Store.Podcast().UpdateEpisodeStats(&model.PodcastEpisodeStats{
			Id:                    podcastId,
			TotalEpisodes:         len(episodes) + newEpisodesCount - blockedEpisodesCount,
			TotalSeasons:          latestEpisode.Season,
			LastestEpisodePubDate: latestEpisode.PubDate,
		})
		if err != nil {
			job.log.Err(err)
		}
	}

	return nil
}
