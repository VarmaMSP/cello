package job

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	h "github.com/go-http-utils/headers"
	"github.com/mmcdole/gofeed/rss"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type RefreshPodcastJob struct {
	store       store.Store
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewRefreshPodcastJob(store store.Store, workerLimit int) (model.Job, error) {
	return &RefreshPodcastJob{
		store: store,
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

func (job *RefreshPodcastJob) Call(delivery amqp.Delivery) {
	var details model.PodcastFeedDetails
	if err := json.Unmarshal(delivery.Body, &details); err != nil {
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}

	go func() {
		defer delivery.Ack(false)
		defer func() { <-job.rateLimiter }()

		// updated details
		detailsU := details

		feed, headers, err := fetchRssFeed(
			details.FeedUrl,
			map[string]string{h.ETag: details.FeedETag, h.LastModified: details.FeedLastModified},
			job.httpClient,
		)
		if err != nil {
			detailsU.LastRefreshStatus = model.StatusFailure
			goto update_details
		}

		detailsU.FeedETag = headers[h.ETag]
		detailsU.FeedLastModified = headers[h.LastModified]

		if feed == nil {
			detailsU.LastRefreshStatus = model.StatusSuccess
			goto update_details
		}

		if err := job.updateEpisodes(details.Id, feed); err != nil {
			detailsU.LastRefreshStatus = model.StatusFailure
			goto update_details
		}

		detailsU.LastRefreshStatus = model.StatusSuccess

	update_details:
		detailsU.UpdatedAt = model.Now()
		job.store.Podcast().UpdateFeedDetails(&details, &detailsU)
	}()
}

func (job *RefreshPodcastJob) updateEpisodes(podcastId string, feed *rss.Feed) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"job.refresh_podcast.fetch_raw_feed",
		http.StatusInternalServerError,
		map[string]string{"podcast_id": podcastId},
	)

	episodeGuidList, err := job.store.Episode().GetAllGuidsByPodcast(podcastId)
	if err != nil {
		return appErrorC(err.Error())
	}

	// Use map to emulate set
	episodes := map[string]interface{}{}
	for _, g := range episodeGuidList {
		episodes[g] = nil
	}

	// Index rss items by their guid
	items := map[string]*rss.Item{}
	for _, i := range feed.Items {
		if i.ITunesExt != nil && i.ITunesExt.Block == "true" {
			continue
		}
		if guid := model.RssItemGuid(i); guid != "" {
			items[guid] = i
		}
	}

	// Block episodes
	for episodeGuid, _ := range episodes {
		if _, ok := items[episodeGuid]; !ok {
			job.store.Episode().Block(podcastId, episodeGuid)
		}
	}

	// Add New Episodes
	for itemGuid, item := range items {
		if _, ok := episodes[itemGuid]; ok {
			episode := &model.Episode{PodcastId: podcastId}
			if err := episode.LoadDetails(item); err != nil {
				continue
			}
			job.store.Episode().Save(episode)
		}
	}

	return nil
}
