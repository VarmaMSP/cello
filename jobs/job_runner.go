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
	// Scheduler
	scheduler *Scheduler
	// Message Producers
	scheduledJobCallP *rabbitmq.Producer
	importPodcastP    *rabbitmq.Producer
	refreshPodcastP   *rabbitmq.Producer
	// Message Consumers
	scheduledJobCallC *rabbitmq.Consumer
	importPodcastC    *rabbitmq.Consumer
	refreshPodcastC   *rabbitmq.Consumer
	// Jobs
	scrapeItunesJob   model.Job
	importPodcastJob  model.Job
	refreshPodcastJob model.Job
}

func NewJobRunner(store store.Store, producerConn, consumerConn *amqp.Connection) (*JobRunner, error) {
	// producers
	scheduledJobCallP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_SCHEDULED_JOB_CALL,
		DeliveryMode: amqp.Persistent,
	})
	if err != nil {
		return nil, err
	}
	importPodcastP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_IMPORT_PODCAST,
		DeliveryMode: amqp.Persistent,
	})
	if err != nil {
		return nil, err
	}
	refreshPodcastP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_REFRESH_PODCAST,
		DeliveryMode: amqp.Persistent,
	})
	if err != nil {
		return nil, err
	}

	// consumers
	scheduledJobCallC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_SCHEDULED_JOB_CALL,
		ConsumerName:  "scheduled_job_run",
		AutoAck:       false,
		Exclusive:     true,
		PreFetchCount: 100,
	})
	if err != nil {
		return nil, err
	}
	importPodcastC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_IMPORT_PODCAST,
		ConsumerName:  "import_podcast",
		AutoAck:       false,
		Exclusive:     true,
		PreFetchCount: 100,
	})
	if err != nil {
		return nil, err
	}
	refreshPodcastC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_REFRESH_PODCAST,
		ConsumerName:  "refresh_podcast",
		AutoAck:       false,
		Exclusive:     true,
		PreFetchCount: 100,
	})
	if err != nil {
		return nil, err
	}

	scrapeItunesJob, err := job.NewScrapeItunesJob(store, importPodcastP, 10)
	if err != nil {
		return nil, err
	}

	return &JobRunner{
		store:        store,
		producerConn: producerConn,
		consumerConn: consumerConn,
		scheduler:    NewScheduler(store, refreshPodcastP, scheduledJobCallP),

		scheduledJobCallP: scheduledJobCallP,
		importPodcastP:    importPodcastP,
		refreshPodcastP:   refreshPodcastP,

		scheduledJobCallC: scheduledJobCallC,
		importPodcastC:    importPodcastC,
		refreshPodcastC:   refreshPodcastC,

		scrapeItunesJob:   scrapeItunesJob,
		importPodcastJob:  job.NewImportPodcastJob(store, 10),
		refreshPodcastJob: job.NewRefreshPodcastJob(store, 10),
	}, nil
}

func (r *JobRunner) Start() {
	r.scheduler.Start()

	go func() {
		for d := range r.scheduledJobCallC.D {
			var input map[string]string
			if err := json.Unmarshal(d.Body, &input); err != nil {
				continue
			}

			if input["job_name"] == model.JOB_NAME_SCRAPE_ITUNES {
				r.scrapeItunesJob.Call(&d)
			}
		}
	}()

	go func() {
		for d := range r.importPodcastC.D {
			r.importPodcastJob.Call(&d)
		}
	}()

	go func() {
		for d := range r.refreshPodcastC.D {
			r.refreshPodcastJob.Call(&d)
		}
	}()
}
