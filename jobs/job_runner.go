package jobs

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/jobs/job"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
)

type JobRunner struct {
	store        store.Store
	producerConn *amqp.Connection
	consumerConn *amqp.Connection

	// Schedules periodic jobs
	scheduler *Scheduler

	// Message Consumers
	scheduledJobsC     *rabbitmq.Consumer
	importPodcastJobsC *rabbitmq.Consumer

	// Jobs
	scrapeItunesJob  model.Job
	importPodcastJob model.Job
}

func NewJobRunner(store store.Store, producerConn, consumerConn *amqp.Connection) (*JobRunner, error) {
	// producers
	scheduledJobsP, err := rabbitmq.NewProducer(producerConn, model.QUEUE_NAME_SCHEDULED_JOBS, amqp.Persistent)
	if err != nil {
		return nil, err
	}
	importPodcastJobsP, err := rabbitmq.NewProducer(producerConn, model.QUEUE_NAME_IMPORT_PODCAST, amqp.Persistent)
	if err != nil {
		return nil, err
	}

	// consumers
	scheduledJobsC, err := rabbitmq.NewConsumer(consumerConn, model.QUEUE_NAME_SCHEDULED_JOBS)
	if err != nil {
		return nil, err
	}
	importPodcastJobsC, err := rabbitmq.NewConsumer(consumerConn, model.QUEUE_NAME_IMPORT_PODCAST)
	if err != nil {
		return nil, err
	}

	return &JobRunner{
		store:              store,
		producerConn:       producerConn,
		consumerConn:       consumerConn,
		scheduler:          NewScheduler(store, scheduledJobsP),
		scheduledJobsC:     scheduledJobsC,
		importPodcastJobsC: importPodcastJobsC,
		scrapeItunesJob:    job.NewScrapeItunesJob(store, importPodcastJobsP, 10),
		importPodcastJob:   job.NewImportPodcastJob(store, 10),
	}, nil
}

func (r *JobRunner) Start() {
	go r.scheduler.Start()

	go func() {
		for i := range r.scheduledJobsC.D {
			msg := model.MapFromJson(i.Body)
			if msg["job_name"] == model.JOB_NAME_CRAWL_ITUNES {
				r.scrapeItunesJob.InputChan() <- nil
			}
		}
	}()

	go func() {
		for i := range r.importPodcastJobsC.D {
			var input model.ImportPodcastInput
			if err := json.Unmarshal(i.Body, &input); err != nil {
				continue
			}

			r.importPodcastJob.InputChan() <- input
		}
	}()
}
