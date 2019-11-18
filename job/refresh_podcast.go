package job

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	h "github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
	"github.com/varmamsp/gofeed/rss"
)

type RefreshPodcastJob struct {
	*app.App
	input       <-chan amqp.Delivery
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewRefreshPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.RefreshPodcast.WorkerLimit

	refreshPodcastC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_REFRESH_PODCAST,
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

		// updated details
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

func (job *RefreshPodcastJob) updateEpisodes(podcastId string, rssFeed *rss.Feed) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"job.refresh_podcast.update_episodes",
		http.StatusInternalServerError,
		map[string]string{"podcast_id": podcastId},
	)

	episodes, err := job.Store.Episode().GetAllByPodcast(podcastId, "pub_date_desc", 10000, 0)
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
		if item.ITunesExt != nil && item.ITunesExt.Block == "true" {
			continue
		}
		if guid := model.RssItemGuid(item); guid != "" {
			rssItemMap[guid] = item
		}
	}

	// Block episodes
	for episodeGuid, _ := range episodeMap {
		if _, ok := rssItemMap[episodeGuid]; !ok {
			job.Store.Episode().Block(podcastId, episodeGuid)
		}
	}

	// Add New Episodes
	for rssItemGuid, rssItem := range rssItemMap {
		if _, ok := episodeMap[rssItemGuid]; !ok {
			episode := &model.Episode{PodcastId: podcastId}
			if err := episode.LoadDetails(rssItem); err != nil {
				continue
			}
			job.Store.Episode().Save(episode)
		}
	}

	return nil
}
