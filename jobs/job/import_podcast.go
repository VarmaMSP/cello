package job

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type ImportPodcastJob struct {
	I           chan interface{}
	store       store.Store
	httpClient  *http.Client
	workerLimit int
}

func NewImportPodcastJob(store store.Store, workerLimit int) model.Job {
	return &ImportPodcastJob{
		I:     make(chan interface{}, 100),
		store: store,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        2 * workerLimit,
				MaxIdleConnsPerHost: workerLimit,
			},
		},
		workerLimit: workerLimit,
	}
}

func (job *ImportPodcastJob) Start() *model.AppError {
	go job.pollInput()

	return nil
}

func (job *ImportPodcastJob) Stop() *model.AppError {
	return nil
}

func (job *ImportPodcastJob) InputChan() chan interface{} {
	return job.I
}

func (job *ImportPodcastJob) pollInput() {
	semaphore := make(chan int, job.workerLimit)

	for {
		i, ok := (<-job.I).(*model.ImportPodcastInput)
		if !ok || i.Id == "" || i.FeedUrl == "" {
			continue
		}

		semaphore <- 0
		go func(feedUrl, itunesId string) {
			defer func() { <-semaphore }()

			status, _ := job.store.ItunesMeta().GetStatus(itunesId)
			if status != model.StatusPending {
				return
			}
			status = model.StatusSuccess
			if err := job.AddToDb(feedUrl); err != nil {
				status = model.StatusFailure
			}

			job.store.ItunesMeta().SetStatus(itunesId, status)
		}(i.FeedUrl, i.Id)
	}
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
