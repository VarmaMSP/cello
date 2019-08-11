package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"

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

func (job *ImportPodcastJob) Start() *model.AppError {
	return nil
}

func (job *ImportPodcastJob) Stop() *model.AppError {
	return nil
}

func (job *ImportPodcastJob) Call(delivery *amqp.Delivery) {
	var input model.ItunesMeta
	if err := json.Unmarshal(delivery.Body, &input); err != nil {
		return
	}

	job.rateLimiter <- struct{}{}

	go func(meta *model.ItunesMeta) {
		defer func() { <-job.rateLimiter }()

		metaU := *meta
		if err := job.AddToDb(meta.FeedUrl); err != nil {
			metaU.AddedToDb = model.StatusFailure
		} else {
			metaU.AddedToDb = model.StatusSuccess
		}

		job.store.ItunesMeta().Update(meta, &metaU)
	}(&input)
}

func (job *ImportPodcastJob) AddToDb(feedUrl string) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"jobs.podcast_jobort.add_to_db",
		http.StatusInternalServerError,
		map[string]string{"feed_url": feedUrl},
	)

	// fetch rss feed
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return appErrorC(err.Error())
	}
	resp, err := job.httpClient.Do(req)
	if err != nil {
		return appErrorC(err.Error())
	}

	// parse rss feed
	parser := &rss.Parser{}
	feed, err := parser.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return appErrorC(fmt.Sprintf("Cannot parse feed: %s", err.Error()))
	}

	// Save Podcast
	podcast := &model.Podcast{
		Id:               xid.New().String(),
		FeedUrl:          feedUrl,
		FeedETag:         resp.Header.Get("ETag"),
		FeedLastModified: resp.Header.Get("Last-Modified"),
		RefreshEnabled:   1,
		RefreshInterval:  100,
	}
	if err := podcast.LoadDetails(feed); err != nil {
		return err
	}
	if err := job.store.Podcast().Save(podcast); err != nil {
		return err
	}

	// Save Episodes
	for _, item := range feed.Items {
		episode := &model.Episode{
			Id:        xid.New().String(),
			PodcastId: podcast.Id,
		}
		if err := episode.LoadDetails(item); err != nil {
			continue
		}
		if err := job.store.Episode().Save(episode); err != nil {
			fmt.Printf("%s %s: %s\n", feedUrl, episode.Title, err.Error())
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
			job.store.Category().SavePodcastCategory(&model.PodcastCategory{
				PodcastId:  podcast.Id,
				CategoryId: categoryId,
			})
		}
	}

	return nil
}
