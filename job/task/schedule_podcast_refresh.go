package task

import (
	"time"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
	"github.com/varmamsp/cello/store"
)

type SchedulePodcastRefresh struct {
	*app.App
	store           store.Store
	refreshPodcastP *rabbitmq.Producer
}

func NewSchedulePodcastRefresh(app *app.App, config *model.Config) (*SchedulePodcastRefresh, error) {
	refreshPodcastP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_REFRESH_PODCAST,
		DeliveryMode: config.Queues.RefreshPodcast.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}

	return &SchedulePodcastRefresh{
		App:             app,
		refreshPodcastP: refreshPodcastP,
	}, nil
}

func (s *SchedulePodcastRefresh) Call() {
	s.Log.Info().Msg("Schedule podcast refresh started")
	limit := 10000
	ticker := time.NewTicker(time.Minute)

	for _ = range ticker.C {
		for createdAfter := int64(0); ; {
			feeds, err := s.store.Feed().GetAllToBeRefreshed(createdAfter, limit)
			if err != nil {
				break
			}

			for _, feed := range feeds {
				feedU := feed
				feedU.LastRefreshAt = model.Now()
				feedU.LastRefreshComment = "PENDING"
				if err := s.store.Feed().Update(feed, feedU); err != nil {
					continue
				}

				s.refreshPodcastP.D <- feedU
			}

			if len(feeds) < limit {
				break
			}
			createdAfter = feeds[len(feeds)-1].CreatedAt
		}
	}
}
