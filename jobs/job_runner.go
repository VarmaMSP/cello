package jobs

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/jobs/job"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/elasticsearch"
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
	createThumbnailP  *rabbitmq.Producer

	// Message Consumers
	scheduledJobCallC *rabbitmq.Consumer
	importPodcastC    *rabbitmq.Consumer
	refreshPodcastC   *rabbitmq.Consumer
	createThumbnailC  *rabbitmq.Consumer

	// Jobs
	scrapeItunesJob    model.Job
	importPodcastJob   model.Job
	refreshPodcastJob  model.Job
	createThumbnailJob model.Job
}

func NewJobRunner(store store.Store, producerConn, consumerConn *amqp.Connection, qConfig *model.RabbitmqQueuesConfig, esConfig *model.ElasticsearchConfig) (*JobRunner, error) {
	// producers
	scheduledJobCallP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_SCHEDULED_JOB_CALL,
		DeliveryMode: qConfig.ScheduledJobCallQueue.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}
	importPodcastP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_IMPORT_PODCAST,
		DeliveryMode: qConfig.ImportPodcastQueue.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}
	refreshPodcastP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_REFRESH_PODCAST,
		DeliveryMode: qConfig.RefreshPodcastQueue.DeliveryMode,
	})
	if err != nil {
		return nil, err
	}
	createThumbnailP, err := rabbitmq.NewProducer(producerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_CREATE_THUMBNAIL,
		DeliveryMode: 0,
	})
	if err != nil {
		return nil, err
	}

	// consumers
	scheduledJobCallC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_SCHEDULED_JOB_CALL,
		ConsumerName:  qConfig.ScheduledJobCallQueue.ConsumerName,
		AutoAck:       qConfig.ScheduledJobCallQueue.ConsumerAutoAck,
		Exclusive:     qConfig.ScheduledJobCallQueue.ConsumerExclusive,
		PreFetchCount: qConfig.ScheduledJobCallQueue.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}
	importPodcastC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_IMPORT_PODCAST,
		ConsumerName:  qConfig.ImportPodcastQueue.ConsumerName,
		AutoAck:       qConfig.ImportPodcastQueue.ConsumerAutoAck,
		Exclusive:     qConfig.ImportPodcastQueue.ConsumerExclusive,
		PreFetchCount: qConfig.ImportPodcastQueue.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}
	refreshPodcastC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_REFRESH_PODCAST,
		ConsumerName:  qConfig.RefreshPodcastQueue.ConsumerName,
		AutoAck:       qConfig.RefreshPodcastQueue.ConsumerAutoAck,
		Exclusive:     qConfig.RefreshPodcastQueue.ConsumerExclusive,
		PreFetchCount: qConfig.RefreshPodcastQueue.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}
	createThumbnailC, err := rabbitmq.NewConsumer(consumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     model.QUEUE_NAME_CREATE_THUMBNAIL,
		ConsumerName:  qConfig.CreateThumbnailQueue.ConsumerName,
		AutoAck:       qConfig.CreateThumbnailQueue.ConsumerAutoAck,
		Exclusive:     qConfig.CreateThumbnailQueue.ConsumerExclusive,
		PreFetchCount: qConfig.CreateThumbnailQueue.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	// Elasticsearch
	esClient, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, err
	}

	// Jobs
	scrapeItunesJob, err := job.NewScrapeItunesJob(store, importPodcastP, 10)
	if err != nil {
		return nil, err
	}
	importPodcastJob, err := job.NewImportPodcastJob(store, esClient, 10)
	if err != nil {
		return nil, err
	}
	refreshPodcastJob, err := job.NewRefreshPodcastJob(store, 10)
	if err != nil {
		return nil, err
	}
	createThumbnailJob, err := job.NewCreateThumbnailJob(10)
	if err != nil {
		return nil, err
	}

	return &JobRunner{
		store:        store,
		producerConn: producerConn,
		consumerConn: consumerConn,
		scheduler:    NewScheduler(store, scheduledJobCallP),

		scheduledJobCallP: scheduledJobCallP,
		importPodcastP:    importPodcastP,
		refreshPodcastP:   refreshPodcastP,
		createThumbnailP:  createThumbnailP,

		scheduledJobCallC: scheduledJobCallC,
		importPodcastC:    importPodcastC,
		refreshPodcastC:   refreshPodcastC,
		createThumbnailC:  createThumbnailC,

		scrapeItunesJob:    scrapeItunesJob,
		importPodcastJob:   importPodcastJob,
		refreshPodcastJob:  refreshPodcastJob,
		createThumbnailJob: createThumbnailJob,
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
				r.scrapeItunesJob.Call(d)
			}
		}
	}()

	go func() {
		for d := range r.importPodcastC.D {
			r.importPodcastJob.Call(d)
		}
	}()

	go func() {
		for d := range r.refreshPodcastC.D {
			r.refreshPodcastJob.Call(d)
		}
	}()

	go func() {
		for d := range r.createThumbnailC.D {
			r.createThumbnailJob.Call(d)
		}
	}()
}
