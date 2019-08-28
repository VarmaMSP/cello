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

func NewScheduleRefreshJob(store store.Store, refreshPodcastP *rabbitmq.Producer) (model.Job, *model.AppError) {
	return &ScheduleRefreshJob{store, refreshPodcastP}, nil
}

func (job *ScheduleRefreshJob) Call(delivery amqp.Delivery) {
	delivery.Ack(false)

	limit := 10000
	ticker := time.NewTicker(time.Minute)

	for _ = range ticker.C {
		for createdAfter := int64(0); ; {
			detailsList, err := job.store.Podcast().GetAllToBeRefreshed(createdAfter, limit)
			if err != nil {
				break
			}

			for _, details := range detailsList {
				detailsU := details
				detailsU.LastRefreshAt = model.Now()
				detailsU.LastRefreshStatus = model.StatusPending
				if err := job.store.Podcast().UpdateFeedDetails(details, detailsU); err != nil {
					continue
				}

				fmt.Printf("Scheduled podcast refresh: %s\n", details.Id)
				job.refreshPodcastP.D <- detailsU
			}

			if len(detailsList) < limit {
				break
			}
			createdAfter = detailsList[len(detailsList)-1].CreatedAt
		}
	}
}
