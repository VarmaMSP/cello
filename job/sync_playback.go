package job

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type SyncPlaybackJob struct {
	*app.App
	input          <-chan amqp.Delivery
	inputBatchSize int
}

func NewSyncPlaybackJob(app *app.App, config *model.Config) (model.Job, error) {
	syncPlaybackC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     rabbitmq.QUEUE_NAME_SYNC_PLAYBACK,
		ConsumerName:  config.Queues.SyncPlayback.ConsumerName,
		AutoAck:       config.Queues.SyncPlayback.ConsumerAutoAck,
		Exclusive:     config.Queues.SyncPlayback.ConsumerExclusive,
		PreFetchCount: config.Queues.SyncPlayback.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	return &SyncPlaybackJob{
		App:            app,
		input:          syncPlaybackC.D,
		inputBatchSize: config.Queues.SyncPlayback.ConsumerPreFetchCount,
	}, nil
}

func (job *SyncPlaybackJob) Run() {
	timeout := time.NewTimer(time.Minute)

	for {
		var deliveries []amqp.Delivery
	BATCH_LOOP:
		for i, _ := 0, timeout.Reset(30*time.Second); i < job.inputBatchSize; i++ {
			select {
			case delivery := <-job.input:
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
			progress := &model.PlaybackProgress{
				UserId:        y[len(y)-1].UserId,
				EpisodeId:     y[len(y)-1].EpisodeId,
				Progress:      y[len(y)-1].Position,
				ProgressDelta: y[len(y)-1].Position - y[0].Position,
			}
			job.Store.Playback().Update(progress)
		}
	}
}
