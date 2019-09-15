package app

import (
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/jobs/job"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/elasticsearch"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/store/sqlstore"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type App struct {
	Store                store.Store
	ElasticSearch        *elastic.Client
	RabbitmqProducerConn *amqp.Connection
	RabbitmqConsumerConn *amqp.Connection

	GoogleOAuthConfig *oauth2.Config

	ScrapeItunesJob    model.Job
	ImportPodcastJob   model.Job
	RefreshPodcastJob  model.Job
	CreateThumbnailJob model.Job

	Log zerolog.Logger
}

func NewApp(config model.Config) (*App, error) {
	app := &App{}

	dev := true
	if dev {
		app.Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	} else {
		app.Log = zerolog.New(os.Stdout).With().Timestamp().Logger()

	}

	log.Info().Msg("Connecting to Mysql ...")
	store, err := sqlstore.NewSqlStore(&config.Mysql)
	if err != nil {
		return nil, err
	}
	app.Store = store

	log.Info().Msg("Connecting to ElasticSearch ...")
	elasticSearch, err := elasticsearch.NewClient(&config.Elasticsearch)
	if err != nil {
		return nil, err
	}
	app.ElasticSearch = elasticSearch

	log.Info().Msg("Connecting to Rabbitmq ...")
	rabbitmqProducerConn, err := rabbitmq.NewConnection(&config.Rabbitmq)
	if err != nil {
		return nil, err
	}
	app.RabbitmqProducerConn = rabbitmqProducerConn

	rabbitmqConsumerConn, err := rabbitmq.NewConnection(&config.Rabbitmq)
	if err != nil {
		return nil, err
	}
	app.RabbitmqConsumerConn = rabbitmqConsumerConn

	if config.Jobs.ScrapeItunes.Enable {
		job, err := job.NewScrapeItunesJob(app, &config)
		if err != nil {
			return nil, err
		}
		app.ScrapeItunesJob = job
	}

	if config.Jobs.ImportPodcast.Enable {
		job, err := job.NewImportPodcastJob(app, &config)
		if err != nil {
			return nil, err
		}
		app.ImportPodcastJob = job
	}

	if config.Jobs.RefreshPodcast.Enable {
		job, err := job.NewRefreshPodcastJob(app, &config)
		if err != nil {
			return nil, err
		}
		app.RefreshPodcastJob = job
	}

	if config.Jobs.CreateThumbnail.Enable {
		job, err := job.NewCreateThumbnailJob(app, &config)
		if err != nil {
			return nil, err
		}
		app.CreateThumbnailJob = job
	}

	app.GoogleOAuthConfig = &oauth2.Config{
		ClientID:     config.OAuth.Google.ClientId,
		ClientSecret: config.OAuth.Google.ClientSecret,
		RedirectURL:  config.OAuth.Google.RedirectUrl,
		Endpoint:     google.Endpoint,
		Scopes:       config.OAuth.Google.Scopes,
	}

	return app, nil
}
