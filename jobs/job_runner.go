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
	// Message Producers
	scheduledWorkP *rabbitmq.Producer
	importPodcastP *rabbitmq.Producer
	// Message Consumers
	scheduledWorkC *rabbitmq.Consumer
	importPodcastC *rabbitmq.Consumer
	// Jobs
	scrapeItunesJob  model.Job
	importPodcastJob model.Job
}

func NewJobRunner(store store.Store, producerConn, consumerConn *amqp.Connection) (*JobRunner, error) {
	// producers
	scheduledWorkP, err := rabbitmq.NewProducer(producerConn, model.QUEUE_NAME_SCHEDULED_WORK, amqp.Persistent)
	if err != nil {
		return nil, err
	}
	importPodcastP, err := rabbitmq.NewProducer(producerConn, model.QUEUE_NAME_IMPORT_PODCAST, amqp.Persistent)
	if err != nil {
		return nil, err
	}

	// consumers
	scheduledWorkC, err := rabbitmq.NewConsumer(consumerConn, model.QUEUE_NAME_SCHEDULED_WORK)
	if err != nil {
		return nil, err
	}
	importPodcastC, err := rabbitmq.NewConsumer(consumerConn, model.QUEUE_NAME_IMPORT_PODCAST)
	if err != nil {
		return nil, err
	}

	return &JobRunner{
		store:        store,
		producerConn: producerConn,
		consumerConn: consumerConn,
		scheduler:    NewScheduler(store, scheduledWorkP),

		scheduledWorkP: scheduledWorkP,
		importPodcastP: importPodcastP,

		scheduledWorkC: scheduledWorkC,
		importPodcastC: importPodcastC,

		scrapeItunesJob:  job.NewScrapeItunesJob(store, importPodcastP, 10),
		importPodcastJob: job.NewImportPodcastJob(store, 10),
	}, nil
}

func (r *JobRunner) Start() {
	go r.scheduler.Start()

	go func() {
		for i := range r.scheduledWorkC.D {
			var input model.ScheduledWorkInput
			if err := json.Unmarshal(i.Body, &input); err != nil {
				continue
			}

			if input.JobName == model.JOB_NAME_SCRAPE_ITUNES {
				r.scrapeItunesJob.InputChan() <- nil
			}
		}
	}()

	go func() {
		for i := range r.importPodcastC.D {
			var input model.ImportPodcastInput
			if err := json.Unmarshal(i.Body, &input); err != nil {
				continue
			}

			r.importPodcastJob.InputChan() <- input
		}
	}()
}
