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
	scheduledJobCallP *rabbitmq.Producer
}

func NewScheduler(store store.Store, scheduledJobCallP *rabbitmq.Producer) *Scheduler {
	return &Scheduler{
		store:             store,
		scheduledJobCallP: scheduledJobCallP,
	}
}

func (s *Scheduler) Start() {
	go s.scheduleJobRun()
}

func (s *Scheduler) scheduleJobRun() {
	ticker := time.NewTicker(10 * time.Second)

	for _ = range ticker.C {
		schedules, err := s.store.JobSchedule().GetAllActive()
		if err != nil {
			continue
		}

		fmt.Println("scheduling")

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

	fmt.Printf("Scheduled Periodic Job: %s\n", jobName)
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
