package job

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"
	"github.com/olivere/elastic/v7"

	h "github.com/go-http-utils/headers"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/elasticsearch"
	"github.com/varmamsp/cello/store"
)

type ImportPodcastJob struct {
	store       store.Store
	esClient    *elastic.Client
	httpClient  *http.Client
	rateLimiter chan struct{}
}

func NewImportPodcastJob(store store.Store, esClient *elastic.Client, workerLimit int) (model.Job, error) {
	return &ImportPodcastJob{
		store:    store,
		esClient: esClient,
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

func (job *ImportPodcastJob) Call(delivery amqp.Delivery) {
	var meta model.ItunesMeta
	if err := json.Unmarshal(delivery.Body, &meta); err != nil {
		delivery.Ack(false)
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
			metaU.Comment = err.Error()
			goto update_meta
		}

		if podcastId, err := job.savePodcast(feed, meta.FeedUrl, headers); err != nil {
			metaU.AddedToDb = model.StatusFailure
			metaU.Comment = err.Error()
			goto update_meta
		} else {
			job.indexPodcast(podcastId, feed)
			metaU.AddedToDb = model.StatusSuccess
		}

	update_meta:
		metaU.UpdatedAt = model.Now()
		job.store.ItunesMeta().Update(&meta, &metaU)
	}()
}

func (job *ImportPodcastJob) savePodcast(feed *rss.Feed, feedUrl string, headers map[string]string) (string, *model.AppError) {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_import_job.save_podcast",
		http.StatusInternalServerError,
		map[string]string{"title": feed.Title},
	)

	// Save podcast
	podcast := &model.Podcast{
		FeedUrl:          feedUrl,
		FeedETag:         headers[h.ETag],
		FeedLastModified: headers[h.LastModified],
	}
	if err := podcast.LoadDetails(feed); err != nil {
		return "", appErrorC(err.Error())
	}
	podcast.SetRefershInterval(feed.Items)
	if err := job.store.Podcast().Save(podcast); err != nil {
		return "", appErrorC(err.Error())
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

	return podcast.Id, nil
}

func (job *ImportPodcastJob) indexPodcast(podcastId string, feed *rss.Feed) {
	doc := &model.PodcastDocument{Id: podcastId}
	doc.LoadDetails(feed)

	job.esClient.Index().
		Index(elasticsearch.PodcastIndexName).
		Id(doc.Id).
		BodyJson(doc).
		Do(context.TODO())
}
