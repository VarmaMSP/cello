package job

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/varmamsp/gofeed/rss"
	"gopkg.in/jdkato/prose.v2"

	h "github.com/go-http-utils/headers"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type ImportPodcastJob struct {
	*app.App
	log              zerolog.Logger
	input            <-chan amqp.Delivery
	httpClient       *http.Client
	rateLimiter      chan struct{}
	createThumbnailP *rabbitmq.Producer
}

func NewImportPodcastJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.ImportPodcast.WorkerLimit

	importPodcastC, err := rabbitmq.NewConsumer(
		app.RabbitmqConsumerConn,
		&rabbitmq.ConsumerOpts{
			QueueName:     rabbitmq.QUEUE_NAME_IMPORT_PODCAST,
			ConsumerName:  config.Queues.ImportPodcast.ConsumerName,
			AutoAck:       config.Queues.ImportPodcast.ConsumerAutoAck,
			Exclusive:     config.Queues.ImportPodcast.ConsumerExclusive,
			PreFetchCount: config.Queues.ImportPodcast.ConsumerPreFetchCount,
		},
	)
	if err != nil {
		return nil, err
	}

	createThumbnailP, err := rabbitmq.NewProducer(
		app.RabbitmqProducerConn,
		&rabbitmq.ProducerOpts{
			ExchangeName: rabbitmq.EXCHANGE_NAME_PHENOPOD_DIRECT,
			RoutingKey:   rabbitmq.ROUTING_KEY_CREATE_THUMBNAIL,
			DeliveryMode: config.Queues.CreateThumbnail.DeliveryMode,
		},
	)
	if err != nil {
		return nil, err
	}

	return &ImportPodcastJob{
		App:   app,
		log:   app.Log.With().Str("job", "import_podcast").Logger(),
		input: importPodcastC.D,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter:      make(chan struct{}, workerLimit),
		createThumbnailP: createThumbnailP,
	}, nil
}

func (job *ImportPodcastJob) Run() {
	for d := range job.input {
		job.Call(d)
	}
}

type EntitiesToSave struct {
	podcast           *model.Podcast
	episodes          []*model.Episode
	podcastCategories []*model.PodcastCategory
	keywords          []*model.Keyword
}

type EntitiesToIndex struct {
	podcast  *model.Podcast
	episodes []*model.Episode
	keywords []*model.Keyword
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

		now := model.Now()

		// Updated feed
		feedU := feed

		if rssFeed, headers, err := fetchRssFeed(feed.Url, map[string]string{}, job.httpClient); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else if entitiesToSave, err := job.extract(feed.Id, rssFeed); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else if entitiesToIndex, err := job.save(entitiesToSave); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else if err := job.index(entitiesToIndex); err != nil {
			feedU.LastRefreshComment = err.GetComment()
			goto update_feed

		} else {
			feedU.ETag = headers[h.ETag]
			feedU.LastModified = headers[h.LastModified]
			feedU.LastRefreshComment = "SUCCESS"
			if feedU.SetRefershInterval(rssFeed); feedU.RefreshEnabled == 1 {
				feedU.NextRefreshAt = int64(feedU.RefreshInterval) + model.Now()
			}
		}

	update_feed:
		feedU.LastRefreshAt = now
		feedU.UpdatedAt = now
		if err := job.Store.Feed().Update(&feed, &feedU); err != nil {
			job.log.Err(err)
		}
	}()
}

func (job *ImportPodcastJob) extract(feedId int64, rssFeed *rss.Feed) (*EntitiesToSave, *model.AppError) {
	// Episodes
	var episodes []*model.Episode
	for _, item := range rssFeed.Items {
		episode := &model.Episode{PodcastId: feedId}
		if err := episode.LoadDetails(item); err != nil {
			job.log.Err(err)
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

	// Keywords
	keywords := []*model.Keyword{}
	doc, err := prose.NewDocument(podcast.Description)
	if err != nil {
		job.log.Err(err)
	} else {
		for _, ent := range doc.Entities() {
			if model.IsValidKeyword(ent.Text) {
				keywords = append(keywords, &model.Keyword{Text: strings.ToLower(ent.Text)})
			}
		}
	}

	return &EntitiesToSave{
		podcast:           podcast,
		episodes:          episodes,
		podcastCategories: podcastCategories,
		keywords:          keywords,
	}, nil
}

func (job *ImportPodcastJob) save(toSave *EntitiesToSave) (*EntitiesToIndex, *model.AppError) {
	// Podcast
	if err := job.Store.Podcast().Save(toSave.podcast); err != nil {
		return nil, err
	}

	// Categories
	for _, podcastCategory := range toSave.podcastCategories {
		if err := job.Store.Category().SavePodcastCategory(podcastCategory); err != nil {
			job.log.Err(err)
			continue
		}
	}

	// Episodes
	episodesToIndex := []*model.Episode{}
	for _, episode := range toSave.episodes {
		if err := job.Store.Episode().Save(episode); err != nil {
			job.log.Err(err)
			continue
		}
		episodesToIndex = append(episodesToIndex, episode)
	}

	// Keywords
	keywordsToIndex := []*model.Keyword{}
	for _, keyword := range toSave.keywords {
		k, err := job.Store.Keyword().Upsert(keyword)
		if err != nil {
			job.log.Err(err)
			continue
		}

		if _, err := job.Store.Keyword().SavePodcastKeyword(&model.PodcastKeyword{
			KeywordId: k.Id,
			PodcastId: toSave.podcast.Id,
		}); err != nil {
			job.log.Err(err)
			continue
		}
		keywordsToIndex = append(keywordsToIndex, k)
	}

	// Create thumbnail
	job.createThumbnailP.D <- CreateThumbnailJobInput{
		Id:         toSave.podcast.Id,
		Type:       "PODCAST",
		ImageSrc:   toSave.podcast.ImagePath,
		ImageTitle: model.UrlParamFromId(toSave.podcast.Title, toSave.podcast.Id),
	}

	return &EntitiesToIndex{
		podcast:  toSave.podcast,
		episodes: episodesToIndex,
		keywords: keywordsToIndex,
	}, nil
}

func (job *ImportPodcastJob) index(toIndex *EntitiesToIndex) *model.AppError {
	podcast := toIndex.podcast
	episodes := toIndex.episodes
	keywords := toIndex.keywords
	indexRequests := make([]elastic.BulkableRequest, len(keywords)+len(episodes)+1)

	indexRequests[0] = elastic.NewBulkIndexRequest().
		Index(elasticsearch.PodcastIndexName).
		Id(model.StrFromInt64(podcast.Id)).
		Doc(&model.PodcastIndex{
			Id:          podcast.Id,
			Title:       podcast.Title,
			Author:      podcast.Author,
			Description: podcast.Description,
			Type:        podcast.Type,
			Complete:    podcast.Complete,
		})

	for i, keyword := range keywords {
		indexRequests[i+1] = elastic.NewBulkIndexRequest().
			Index(elasticsearch.KeywordIndexName).
			Id(model.StrFromInt64(keyword.Id)).
			Doc(&model.KeywordIndex{
				Text:    keyword.Text,
				AddedBy: "0",
			})
	}

	for i, episode := range episodes {
		indexRequests[i+1] = elastic.NewBulkIndexRequest().
			Index(elasticsearch.EpisodeIndexName).
			Id(model.StrFromInt64(episode.Id)).
			Doc(&model.EpisodeIndex{
				Id:          episode.Id,
				PodcastId:   episode.PodcastId,
				Title:       episode.Title,
				Description: model.StripHTMLTags(episode.Description),
				PubDate:     episode.PubDate,
				Duration:    episode.Duration,
				Type:        episode.Type,
			})
	}

	bulkIndexSize := 20

	for i := 0; i < len(indexRequests); i += bulkIndexSize {
		end := i + bulkIndexSize
		if end > len(indexRequests) {
			end = len(indexRequests)
		}

		if _, err := job.ElasticSearch.Bulk().Add(indexRequests[i:end]...).Do(context.TODO()); err != nil {
			return model.NewAppError("jobs.podcast_import_job.save_podcast", err.Error(), http.StatusBadRequest, nil)
		}
	}

	return nil
}
