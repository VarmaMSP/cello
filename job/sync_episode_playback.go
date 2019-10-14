package job

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type SyncEpisodePlaybackJob struct {
	*app.App
	input          <-chan amqp.Delivery
	inputBatchSize int
}

func NewSyncEpisodePlaybackJob(app *app.App, config *model.Config) (model.Job, error) {
	syncEpisodePlaybackC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_SYNC_EPISODE_PLAYBACK,
		ConsumerName:  config.Queues.SyncEpisodePlayback.ConsumerName,
		AutoAck:       config.Queues.SyncEpisodePlayback.ConsumerAutoAck,
		Exclusive:     config.Queues.SyncEpisodePlayback.ConsumerExclusive,
		PreFetchCount: config.Queues.SyncEpisodePlayback.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	return &SyncEpisodePlaybackJob{
		App:            app,
		input:          syncEpisodePlaybackC.D,
		inputBatchSize: config.Queues.SyncEpisodePlayback.ConsumerPreFetchCount,
	}, nil
}

func (job *SyncEpisodePlaybackJob) Run() {
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

func (job *SyncEpisodePlaybackJob) Call(deliveries []amqp.Delivery) {
	var progress map[string]*model.EpisodePlayback
	for _, delivery := range deliveries {
		var playback model.EpisodePlayback
		if err := json.Unmarshal(delivery.Body, &playback); err != nil {
			continue
		}

		p := progress[playback.EpisodeId]
		if p == nil || (p.PlayedBy == playback.PlayedBy && p.UpdatedAt < playback.UpdatedAt) {
			progress[playback.EpisodeId] = &playback
		}
	}

	for _, playback := range progress {
		job.Store.Episode().SetPlaybackCurrentTime(playback.EpisodeId, playback.PlayedBy, playback.CurrentTime)
	}
}
