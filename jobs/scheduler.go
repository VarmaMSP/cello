package jobs

import (
	"fmt"
	"time"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
)

// Schedules job calls at certain time
// 1 - Query Job schedule table
// 2 - Push jobs to queue that needs to be run now
// 3 - Repeat the above for every 10 sec.

type Scheduler struct {
	store             store.Store
	refreshPodcastP   *rabbitmq.Producer
	scheduledJobCallP *rabbitmq.Producer
}

func NewScheduler(store store.Store, refreshPodcastP, scheduledJobCallP *rabbitmq.Producer) *Scheduler {
	return &Scheduler{
		store:             store,
		refreshPodcastP:   refreshPodcastP,
		scheduledJobCallP: scheduledJobCallP,
	}
}

func (s *Scheduler) Start() {
	go s.scheduleJobRun()
	go s.schedulePodcastRefresh()
}

func (s *Scheduler) schedulePodcastRefresh() {
	limit := 10000
	ticker := time.NewTicker(time.Minute)

	for _ = range ticker.C {
		for createdAfter := int64(0); ; {
			detailsList, err := s.store.Podcast().GetAllToBeRefreshed(createdAfter, limit)
			if err != nil {
				break
			}

			for _, details := range detailsList {
				detailsU := details
				detailsU.LastRefreshAt = model.Now()
				detailsU.LastRefreshStatus = model.StatusPending
				if err := s.store.Podcast().UpdateFeedDetails(details, detailsU); err != nil {
					continue
				}

				fmt.Printf("Scheduled podcast refresh: %s\n", details.Id)
				s.refreshPodcastP.D <- detailsU
			}

			if len(detailsList) < limit {
				break
			}
			createdAfter = detailsList[len(detailsList)-1].CreatedAt
		}
	}
}

func (s *Scheduler) scheduleJobRun() {
	ticker := time.NewTicker(10 * time.Second)

	for _ = range ticker.C {
		schedules, err := s.store.JobSchedule().GetAllActive()
		if err != nil {
			continue
		}

		for _, schedule := range schedules {
			switch schedule.Type {
			case model.JOB_SCHEDULE_TYPE_PERIODIC:
				s.periodic(schedule.JobName, schedule.RunAt, schedule.RunAfter)

			case model.JOB_SCHEDULE_TYPE_ONEOFF:
				s.oneoff(schedule.JobName, schedule.RunAt)

			case model.JOB_SCHEDULE_TYPE_IMMEDIATE:
				s.immediate(schedule.JobName)
			}
		}
	}
}

func (s *Scheduler) periodic(jobName string, lastRunAt, runAfter int64) {
	now := model.Now()
	if lastRunAt+runAfter > now {
		return
	}
	if err := s.store.JobSchedule().SetRunAt(jobName, now); err != nil {
		return
	}

	fmt.Printf("Scheduled Job: %s\n", jobName)
	s.scheduledJobCallP.D <- map[string]string{"job_name": jobName}
}

func (s *Scheduler) oneoff(jobName string, runAt int64) {
	now := model.Now()
	if runAt > now {
		return
	}
	if err := s.store.JobSchedule().Disable(jobName); err != nil {
		return
	}

	s.scheduledJobCallP.D <- map[string]string{"job_name": jobName}
}

func (s *Scheduler) immediate(jobName string) {
	if err := s.store.JobSchedule().Disable(jobName); err != nil {
		return
	}

	s.scheduledJobCallP.D <- map[string]string{"job_name": jobName}
}
