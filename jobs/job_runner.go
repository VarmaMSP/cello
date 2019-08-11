package jobs

import (
	"encoding/json"
	"fmt"

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
	// Scheduler
	scheduler *Scheduler
	// Message Producers
	scheduledWorkP  *rabbitmq.Producer
	importPodcastP  *rabbitmq.Producer
	refreshPodcastP *rabbitmq.Producer
	// Message Consumers
	scheduledWorkC  *rabbitmq.Consumer
	importPodcastC  *rabbitmq.Consumer
	refreshPodcastC *rabbitmq.Consumer
	// Jobs
	scrapeItunesJob   model.Job
	importPodcastJob  model.Job
	refreshPodcastJob model.Job
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
	refreshPodcastP, err := rabbitmq.NewProducer(producerConn, model.QUEUE_NAME_REFRESH_PODCAST, amqp.Persistent)
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
	refreshPodcastC, err := rabbitmq.NewConsumer(consumerConn, model.QUEUE_NAME_REFRESH_PODCAST)
	if err != nil {
		return nil, err
	}

	return &JobRunner{
		store:        store,
		producerConn: producerConn,
		consumerConn: consumerConn,
		scheduler:    NewScheduler(store, refreshPodcastP, scheduledWorkP),

		scheduledWorkP:  scheduledWorkP,
		importPodcastP:  importPodcastP,
		refreshPodcastP: refreshPodcastP,

		scheduledWorkC:  scheduledWorkC,
		importPodcastC:  importPodcastC,
		refreshPodcastC: refreshPodcastC,

		scrapeItunesJob:   job.NewScrapeItunesJob(store, importPodcastP, 10),
		importPodcastJob:  job.NewImportPodcastJob(store, 10),
		refreshPodcastJob: job.NewRefreshPodcastJob(store, 10),
	}, nil
}

func (r *JobRunner) Start() {
	r.scheduler.Start()
	r.scrapeItunesJob.Start()
	r.importPodcastJob.Start()
	r.refreshPodcastJob.Start()

	go func() {
		for i := range r.scheduledWorkC.D {
			fmt.Printf("Input scheduled work: %s\n", i.Body)
			var input model.ScheduledWorkInput
			if err := json.Unmarshal(i.Body, &input); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(input)
			if input.JobName == model.JOB_NAME_SCRAPE_ITUNES {
				fmt.Println("starting itunes crawl")
				r.scrapeItunesJob.InputChan() <- nil
			}
		}
	}()

	go func() {
		for i := range r.importPodcastC.D {
			var input model.ItunesMeta
			if err := json.Unmarshal(i.Body, &input); err != nil {
				continue
			}

			fmt.Println(input)
			r.importPodcastJob.InputChan() <- &input
		}
	}()

	go func() {
		for i := range r.refreshPodcastC.D {
			fmt.Printf("Input refresh podcast: %s\n", i.Body)
			var input model.PodcastFeedDetails
			if err := json.Unmarshal(i.Body, &input); err != nil {
				continue
			}

			r.refreshPodcastJob.InputChan() <- &input
		}
	}()
}
