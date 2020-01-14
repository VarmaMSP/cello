package task

import (
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
)

type ReimportPodcasts struct {
	*app.App
	importPodcastP *rabbitmq.Producer
}

func NewReimportPodcasts(app *app.App, config *model.Config) (*ReimportPodcasts, error) {
	importPodcastP, err := rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.EXCHANGE_NAME_PHENOPOD_DIRECT,
		RoutingKey:   rabbitmq.ROUTING_KEY_IMPORT_PODCAST,
		DeliveryMode: config.Queues.ImportPodcast.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}

	return &ReimportPodcasts{
		App:            app,
		importPodcastP: importPodcastP,
	}, nil
}

func (s *ReimportPodcasts) Call() {
	go func() {
		limit := 5000
		lastId := int64(0)

		for {
			feeds, err := s.Store.Feed().GetFailedToImportPaginated(lastId, limit)
			if err != nil {
				break
			}

			for _, feed := range feeds {
				s.importPodcastP.D <- feed
			}

			if len(feeds) < limit {
				break
			}
			lastId = feeds[len(feeds)-1].Id
		}
	}()
}
