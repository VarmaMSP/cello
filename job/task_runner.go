package job

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/crawler"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/message_queue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/util/datetime"
)

type TaskRunnerJob struct {
	store store.Store
	se    searchengine.Broker
	mq    message_queue.Broker
	log   zerolog.Logger

	itunesCrawler   *crawler.ItunesCrawler
	refreshPodcastP message_queue.Producer
}

func NewTaskRunnerJob(store store.Store, se searchengine.Broker, mq message_queue.Broker, log zerolog.Logger, config *model.Config) (Job, error) {
	itunesCrawler, err := crawler.NewItunesCrawler(store, mq, log, config)
	if err != nil {
		return nil, err
	}

	refreshPodcastP, err := mq.NewProducer(
		message_queue.EXCHANGE_PHENOPOD_DIRECT,
		message_queue.ROUTING_KEY_REFRESH_PODCAST,
		config.Queues.RefreshPodcast.DeliveryMode,
	)
	if err != nil {
		return nil, err
	}

	return &TaskRunnerJob{
		store:           store,
		se:              se,
		mq:              mq,
		log:             log.With().Str("ctx", "job_server.task_runner").Logger(),
		itunesCrawler:   itunesCrawler,
		refreshPodcastP: refreshPodcastP,
	}, nil
}

func (job *TaskRunnerJob) Start() {
	job.log.Info().Msg("started")
	go func() {
		ticker := time.NewTicker(10 * time.Second)

		for range ticker.C {
			tasks, err := job.store.Task().GetAll()
			if err != nil {
				job.log.Error().Msg(err.Error())
			}

			for _, task := range tasks {
				switch task.Type {
				case model.TASK_TYPE_PERIODIC:
					job.periodic(task)

				case model.TASK_TYPE_ONEOFF:
					job.oneoff(task)

				case model.TASK_TYPE_IMMEDIATE:
					job.immediate(task)
				}
			}
		}
	}()
}

func (job *TaskRunnerJob) periodic(task *model.Task) {
	now := datetime.Unix()
	if task.NextRunAt > now {
		return
	}

	taskU := *task
	taskU.NextRunAt = now + int64(task.Interval)
	taskU.UpdatedAt = now
	if err := job.store.Task().Update(task, &taskU); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}
	job.callTask(task)
}

func (job *TaskRunnerJob) oneoff(task *model.Task) {
	now := datetime.Unix()
	if task.NextRunAt > now {
		return
	}

	taskU := *task
	taskU.Active = 0
	taskU.UpdatedAt = now
	if err := job.store.Task().Update(task, &taskU); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}
	job.callTask(task)
}

func (job *TaskRunnerJob) immediate(task *model.Task) {
	now := datetime.Unix()

	taskU := *task
	taskU.Active = 0
	taskU.UpdatedAt = now
	if err := job.store.Task().Update(task, &taskU); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}
	job.callTask(task)
}

func (job *TaskRunnerJob) callTask(task *model.Task) {
	switch task.Name {
	case "scrape_itunes_directory":
		job.itunesCrawler.Call()

	case "schedule_podcast_refresh":
		go job.schedulePodcastRefresh()

	case "reindex_episodes":
		go job.reindexEpisodes()

	case "reindex_podcasts":
		go job.reindexPodcasts()
	}
}

func (job *TaskRunnerJob) schedulePodcastRefresh() {
	limit := 5000
	lastId := int64(0)

	for {
		feeds, err := job.store.Feed().GetForRefreshPaginated(lastId, limit)
		if err != nil {
			break
		}

		for _, feed := range feeds {
			feedU := feed
			feedU.LastRefreshAt = datetime.Unix()
			feedU.LastRefreshComment = "PENDING"
			if err := job.store.Feed().Update(feed, feedU); err != nil {
				job.log.Error().Msg(err.Error())
			}
			job.refreshPodcastP.Publish(feedU)
		}

		if len(feeds) < limit {
			break
		}
		lastId = feeds[len(feeds)-1].Id
	}
}

func (job *TaskRunnerJob) reindexEpisodes() {
	if err := job.se.DeleteIndex(searchengine.EPISODE_INDEX); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}
	if err := job.se.CreateIndex(searchengine.EPISODE_INDEX, searchengine.EPISODE_INDEX_MAPPING); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}

	lastId, limit := int64(0), 10000
	for {
		episodes, err := job.store.Episode().GetAllPaginated(lastId, limit)
		if err != nil {
			job.log.Error().Msg(err.Error())
			return
		}

		m := make([]model.EsModel, len(episodes))
		for i, episode := range episodes {
			m[i] = episode.ForIndexing()
		}

		if err := job.se.BulkIndex(searchengine.EPISODE_INDEX, m); err != nil {
			job.log.Error().Msg(err.Error())
			return
		}

		if len(episodes) < limit {
			break
		}
		lastId = episodes[len(episodes)-1].Id
	}
}

func (job *TaskRunnerJob) reindexPodcasts() {
	if err := job.se.DeleteIndex(searchengine.PODCAST_INDEX); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}
	if err := job.se.CreateIndex(searchengine.PODCAST_INDEX, searchengine.PODCAST_INDEX_MAPPING); err != nil {
		job.log.Error().Msg(err.Error())
		return
	}

	lastId, limit := int64(0), 10000
	for {
		podcasts, err := job.store.Podcast().GetAllPaginated(lastId, limit)
		if err != nil {
			job.log.Error().Msg(err.Error())
			return
		}

		m := make([]model.EsModel, len(podcasts))
		for i, podcast := range podcasts {
			m[i] = podcast.ForIndexing()
		}

		if err := job.se.BulkIndex(searchengine.PODCAST_INDEX, m); err != nil {
			job.log.Error().Msg(err.Error())
			return
		}

		if len(podcasts) < limit {
			break
		}
		lastId = podcasts[len(podcasts)-1].Id
	}
}
