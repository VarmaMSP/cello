package job

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type RefreshPodcastJob struct {
	store       store.Store
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewRefreshPodcastJob(store store.Store, workerLimit int) model.Job {
	return &RefreshPodcastJob{
		store: store,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        2 * workerLimit,
				MaxIdleConnsPerHost: workerLimit,
			},
		},
		rateLimiter: make(chan struct{}, workerLimit),
	}
}

func (job *RefreshPodcastJob) Start() *model.AppError {
	return nil
}

func (job *RefreshPodcastJob) Stop() *model.AppError {
	return nil
}

func (job *RefreshPodcastJob) Call(delivery *amqp.Delivery) {
	var input model.PodcastFeedDetails
	if err := json.Unmarshal(delivery.Body, &input); err != nil {
		return
	}

	job.rateLimiter <- struct{}{}

	go func(details *model.PodcastFeedDetails) {
		defer func() { <-job.rateLimiter }()

		resp, err := job.call(details.FeedUrl, map[string]string{
			"Cache-Control":     "no-cache",
			"If-None-Match":     details.FeedETag,
			"If-Modified-Since": details.FeedLastModified,
		})

		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				job.updateEpisodes(details.Id, resp.Body)
			}
		}

		detailsU := *details
		if err != nil {
			detailsU.FeedUrl = resp.Request.URL.String()
			detailsU.FeedETag = resp.Header.Get("ETag")
			detailsU.FeedLastModified = resp.Header.Get("Last-Modified")
			detailsU.LastRefreshStatus = model.StatusSuccess
			detailsU.UpdatedAt = model.Now()
		} else {
			detailsU.LastRefreshStatus = model.StatusFailure
			detailsU.UpdatedAt = model.Now()
		}
		job.store.Podcast().UpdateFeedDetails(details, &detailsU)
	}(&input)
}

func (job *RefreshPodcastJob) call(feedUrl string, headers map[string]string) (*http.Response, *model.AppError) {
	appErrorC := model.NewAppErrorC(
		"job.refresh_podcast.fetch_raw_feed",
		http.StatusInternalServerError,
		map[string]string{"feed_url": feedUrl},
	)

	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := job.httpClient.Do(req)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	return resp, nil
}

func (job *RefreshPodcastJob) updateEpisodes(podcastId string, rawFeed io.ReadCloser) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"job.refresh_podcast.fetch_raw_feed",
		http.StatusInternalServerError,
		map[string]string{"podcast_id": podcastId},
	)

	parser := &rss.Parser{}
	feed, err := parser.Parse(rawFeed)
	if err != nil {
		return appErrorC(err.Error())
	}

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
			episode := &model.Episode{
				Id:        xid.New().String(),
				PodcastId: podcastId,
			}
			if err := episode.LoadDetails(item); err != nil {
				continue
			}
			job.store.Episode().Save(episode)
		}
	}

	return nil
}
