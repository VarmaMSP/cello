package job

import (
	"fmt"
	"time"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/job/task"
	"github.com/varmamsp/cello/model"
)

type SchedulerJob struct {
	*app.App
	scrapeItunes           *task.ScrapeItunes
	schedulePodcastrefresh *task.SchedulePodcastRefresh
}

func NewSchedulerJob(app *app.App, config *model.Config) (model.Job, error) {
	scrapeItunes, err := task.NewScrapeItunes(app, config)
	if err != nil {
		return nil, err
	}

	schedulePodcastRefresh, err := task.NewSchedulePodcastRefresh(app, config)
	if err != nil {
		return nil, err
	}

	return &SchedulerJob{
		App:                    app,
		scrapeItunes:           scrapeItunes,
		schedulePodcastrefresh: schedulePodcastRefresh,
	}, nil
}

func (job *SchedulerJob) Run() {
	ticker := time.NewTicker(10 * time.Second)

	for _ = range ticker.C {
		tasks, err := job.Store.Task().GetAllActive()
		if err != nil {
			job.Log.Error().
				Str("at", "Scheduler Job").
				Str("from", err.Id).
				Msg(err.DetailedError)
			continue
		}

		for _, task := range tasks {
			switch task.Type {
			case model.TASK_TYPE_PERIODIC:
				job.periodic(task)

			case model.TASK_TYPE_ONEOFF:
				job.oneoff(task)

			case model.TASK_TYPE_IMMEDIATE:
				job.immediate(task)
			}
		}
	}
}

func (job *SchedulerJob) periodic(task *model.Task) {
	now := model.Now()
	if task.NextRunAt > now {
		return
	}

	taskU := *task
	taskU.NextRunAt = now + int64(task.Interval)
	taskU.UpdatedAt = now
	if err := job.Store.Task().Update(task, &taskU); err != nil {
		fmt.Println(err)
		return
	}
	job.callTask(task)
}

func (job *SchedulerJob) oneoff(task *model.Task) {
	now := model.Now()
	if task.NextRunAt > now {
		return
	}

	taskU := *task
	taskU.Active = 0
	taskU.UpdatedAt = now
	if err := job.Store.Task().Update(task, &taskU); err != nil {
		return
	}
	job.callTask(task)
}

func (job *SchedulerJob) immediate(task *model.Task) {
	now := model.Now()
	taskU := *task
	taskU.Active = 0
	taskU.UpdatedAt = now
	if err := job.Store.Task().Update(task, &taskU); err != nil {
		return
	}
	job.callTask(task)
}

func (job *SchedulerJob) callTask(task *model.Task) {
	switch task.Name {
	case model.TASK_NAME_SCRAPE_ITUNES:
		job.scrapeItunes.Call()
	case model.TASK_NAME_SCHEDULE_PODCAST_REFRESH:
		job.schedulePodcastrefresh.Call()
	}
}
