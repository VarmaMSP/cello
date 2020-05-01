package job

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
)

type SyncPlaybackJob struct {
	store          store.Store
	log            zerolog.Logger
	input          messagequeue.Consumer
	inputBatchSize int
	buffer         chan amqp.Delivery
}

func NewSyncPlaybackJob(store store.Store, mq messagequeue.Broker, log zerolog.Logger, config *model.Config) (Job, error) {
	syncPlaybackC, err := mq.NewConsumer(
		messagequeue.QUEUE_SYNC_PLAYBACK,
		config.Queues.SyncPlayback.ConsumerName,
		config.Queues.SyncPlayback.ConsumerAutoAck,
		config.Queues.SyncPlayback.ConsumerExclusive,
		config.Queues.SyncPlayback.ConsumerPreFetchCount,
	)
	if err != nil {
		return nil, err
	}

	return &SyncPlaybackJob{
		store:          store,
		log:            log.With().Str("ctx", "job_server.sync_playback_job").Logger(),
		input:          syncPlaybackC,
		inputBatchSize: 1000,
		buffer:         make(chan amqp.Delivery),
	}, nil
}

func (job *SyncPlaybackJob) Start() {
	job.log.Info().Msg("started")
	job.input.Consume(func(d amqp.Delivery) { job.buffer <- d })
	go job.Run()
}

func (job *SyncPlaybackJob) Run() {
	timeout := time.NewTimer(10 * time.Second)

	for {
		var deliveries []amqp.Delivery
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(10*time.Second); i < job.inputBatchSize; i++ {
			select {
			case delivery := <-job.buffer:
				deliveries = append(deliveries, delivery)
				if len(deliveries) == job.inputBatchSize && !timeout.Stop() {
					<-timeout.C
				}

			case <-timeout.C:
				break BATCH_LOOP
			}
		}

		if len(deliveries) == 0 {
			continue
		}

		job.Call(deliveries)
	}
}

func (job *SyncPlaybackJob) Call(deliveries []amqp.Delivery) {
	eventsByUserByEpisode := map[int64](map[int64][]*model.PlaybackEvent){}
	for _, delivery := range deliveries {
		event := &model.PlaybackEvent{}
		if err := json.Unmarshal(delivery.Body, event); err != nil {
			continue
		}

		if _, ok := eventsByUserByEpisode[event.UserId]; !ok {
			eventsByUserByEpisode[event.UserId] = map[int64][]*model.PlaybackEvent{}
		}
		if _, ok := eventsByUserByEpisode[event.UserId][event.EpisodeId]; !ok {
			eventsByUserByEpisode[event.UserId][event.EpisodeId] = []*model.PlaybackEvent{}
		}
		eventsByUserByEpisode[event.UserId][event.EpisodeId] = append(eventsByUserByEpisode[event.UserId][event.EpisodeId], event)
	}

	for _, x := range eventsByUserByEpisode {
		for _, y := range x {
			sort.Slice(y, func(i, j int) bool { return false })
			if err := job.store.Playback().Update(&model.Playback{
				UserId:          y[len(y)-1].UserId,
				EpisodeId:       y[len(y)-1].EpisodeId,
				CurrentProgress: y[len(y)-1].Position,
			}); err != nil {
				job.log.Error().Msg(err.Error())
			}
		}
	}
}
