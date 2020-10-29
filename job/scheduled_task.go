package job

import (
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/message_queue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/util/datetime"
)

type SchdeuledTaskJob struct {
	store store.Store
	se    searchengine.Broker
	mq    message_queue.Broker
	log   zerolog.Logger
	input message_queue.Consumer
}

func (job *SchdeuledTaskJob) schedulePodcastRefresh() {
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
			// job.refreshPodcastP.Publish(feedU)
		}

		if len(feeds) < limit {
			break
		}
		lastId = feeds[len(feeds)-1].Id
	}
}

func (job *SchdeuledTaskJob) reindexEpisodes() {
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

func (job *SchdeuledTaskJob) reindexPodcasts() {
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
