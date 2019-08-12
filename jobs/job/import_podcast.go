package job

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"

	"github.com/streadway/amqp"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type ImportPodcastJob struct {
	store       store.Store
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewImportPodcastJob(store store.Store, workerLimit int) model.Job {
	return &ImportPodcastJob{
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

func (job *ImportPodcastJob) Stop() *model.AppError {
	return nil
}

func (job *ImportPodcastJob) Call(delivery *amqp.Delivery) {
	var meta model.ItunesMeta
	if err := json.Unmarshal(delivery.Body, &meta); err != nil {
		return
	}

	job.rateLimiter <- struct{}{}

	go func() {
		defer delivery.Ack(false)
		defer func() { <-job.rateLimiter }()

		// updated meta
		metaU := meta

		feed, headers, err := fetchRssFeed(meta.FeedUrl, map[string]string{}, job.httpClient)
		if err != nil || feed == nil {
			metaU.AddedToDb = model.StatusFailure
			goto update_meta
		}

		if err := job.savePodcast(feed, meta.FeedUrl, headers); err != nil {
			metaU.AddedToDb = model.StatusFailure
			goto update_meta
		}

		metaU.AddedToDb = model.StatusSuccess

	update_meta:
		metaU.UpdatedAt = model.Now()
		job.store.ItunesMeta().Update(&meta, &metaU)
	}()
}

func (job *ImportPodcastJob) savePodcast(feed *rss.Feed, feedUrl string, headers map[string]string) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import_job.save_podcast",
		http.StatusInternalServerError,
		map[string]string{"title": feed.Title},
	)

	// Save podcast
	podcast := &model.Podcast{
		FeedUrl:          feedUrl,
		FeedETag:         headers["ETag"],
		FeedLastModified: headers["Last-Modified"],
	}
	if err := podcast.LoadDetails(feed); err != nil {
		return appErrorC(err.Error())
	}
	podcast.SetRefershInterval(feed.Items)
	if err := job.store.Podcast().Save(podcast); err != nil {
		return appErrorC(err.Error())
	}

	// Save Episodes
	for _, item := range feed.Items {
		episode := &model.Episode{PodcastId: podcast.Id}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		if err := job.store.Episode().Save(episode); err != nil {
			continue
		}
	}

	// Save Categories
	var categoryIds []int
	if feed.ITunesExt != nil {
		for _, c := range feed.ITunesExt.Categories {
			if c.Subcategory != nil {
				categoryIds = append(categoryIds, model.CategoryId(c.Subcategory.Text))
			}
			categoryIds = append(categoryIds, model.CategoryId(c.Text))
		}
	}
	for _, categoryId := range categoryIds {
		if categoryId != -1 {
			category := &model.PodcastCategory{PodcastId: podcast.Id, CategoryId: categoryId}
			if err := job.store.Category().SavePodcastCategory(category); err != nil {
				continue
			}
		}
	}

	return nil
}
