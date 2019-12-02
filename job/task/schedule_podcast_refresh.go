package task

import (
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type SchedulePodcastRefresh struct {
	*app.App
	refreshPodcastP *rabbitmq.Producer
}

func NewSchedulePodcastRefresh(app *app.App, config *model.Config) (*SchedulePodcastRefresh, error) {
	refreshPodcastP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.EXCHANGE_NAME_PHENOPOD_DIRECT,
		RoutingKey:   rabbitmq.ROUTING_KEY_REFRESH_PODCAST,
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
	go func() {
		limit := 10000
		lastId := int64(0)

		for {
			feeds, err := s.Store.Feed().GetForRefreshPaginated(lastId, limit)
			if err != nil {
				break
			}

			for _, feed := range feeds {
				feedU := feed
				feedU.LastRefreshAt = model.Now()
				feedU.LastRefreshComment = "PENDING"
				if err := s.Store.Feed().Update(feed, feedU); err != nil {
					continue
				}
				s.refreshPodcastP.D <- feedU
			}

			if len(feeds) < limit {
				break
			}
			lastId = feeds[len(feeds)-1].Id
		}
	}()
}
