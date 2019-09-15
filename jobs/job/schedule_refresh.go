package job

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
)

type ScheduleRefreshJob struct {
	store           store.Store
	refreshPodcastP *rabbitmq.Producer
}

func NewScheduleRefreshJob(store store.Store, refreshPodcastP *rabbitmq.Producer) (model.Job, error) {
	return &ScheduleRefreshJob{store, refreshPodcastP}, nil
}

func (job *ScheduleRefreshJob) Call(delivery amqp.Delivery) {
	delivery.Ack(false)

	limit := 10000
	ticker := time.NewTicker(time.Minute)

	for _ = range ticker.C {
		for createdAfter := int64(0); ; {
			feeds, err := job.store.Feed().GetAllToBeRefreshed(createdAfter, limit)
			if err != nil {
				break
			}

			for _, feed := range feeds {
				feedU := feed
				feedU.LastRefreshAt = model.Now()
				feedU.LastRefreshComment = "PENDING"
				if err := job.store.Feed().Update(feed, feedU); err != nil {
					continue
				}

				fmt.Printf("Scheduled podcast refresh: %s\n", feed.Id)
				job.refreshPodcastP.D <- feedU
			}

			if len(feeds) < limit {
				break
			}
			createdAfter = feeds[len(feeds)-1].CreatedAt
		}
	}
}
