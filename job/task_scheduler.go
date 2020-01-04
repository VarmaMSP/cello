package job

import (
	"fmt"
	"time"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/job/task"
	"github.com/varmamsp/cello/model"
)

type TaskSchedulerJob struct {
	*app.App
	scrapeTrending         *task.ScrapeTrending
	scrapeCategories       *task.ScrapeCategories
	scrapeItunesDirectory  *task.ScrapeItunesDirectory
	schedulePodcastRefresh *task.SchedulePodcastRefresh
}

func NewTaskSchedulerJob(app *app.App, config *model.Config) (model.Job, error) {
	var err error

	s := &TaskSchedulerJob{App: app}

	s.scrapeTrending, err = task.NewScrapeTrending(app)
	if err != nil {
		return nil, err
	}

	s.scrapeCategories, err = task.NewScrapeCategories(app)
	if err != nil {
		return nil, err
	}

	s.scrapeItunesDirectory, err = task.NewScrapeItunesDirectory(app, config)
	if err != nil {
		return nil, err
	}

	s.schedulePodcastRefresh, err = task.NewSchedulePodcastRefresh(app, config)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (job *TaskSchedulerJob) Run() {
	ticker := time.NewTicker(10 * time.Second)

	for _ = range ticker.C {
		tasks, err := job.Store.Task().GetAll()
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

func (job *TaskSchedulerJob) periodic(task *model.Task) {
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

func (job *TaskSchedulerJob) oneoff(task *model.Task) {
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

func (job *TaskSchedulerJob) immediate(task *model.Task) {
	now := model.Now()

	taskU := *task
	taskU.Active = 0
	taskU.UpdatedAt = now
	if err := job.Store.Task().Update(task, &taskU); err != nil {
		return
	}
	job.callTask(task)
}

func (job *TaskSchedulerJob) callTask(task *model.Task) {
	switch task.Name {
	case model.TASK_NAME_SCRAPE_TRENDING:
		job.scrapeTrending.Call()

	case model.TASK_NAME_SCRAPE_CATEGORIES:
		job.scrapeCategories.Call()

	case model.TASK_NAME_SCRAPE_ITUNES_DIRECTORY:
		job.scrapeItunesDirectory.Call()

	case model.TASK_NAME_SCHEDULE_PODCAST_REFRESH:
		job.schedulePodcastRefresh.Call()
	}
}
