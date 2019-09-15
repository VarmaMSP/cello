package app

import (
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
	"github.com/varmamsp/cello/service/rabbitmq"
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

	app.Log.Info().Msg("Connecting to Mysql ...")
	store, err := sqlstore.NewSqlStore(&config)
	if err != nil {
		return nil, err
	}
	app.Store = store

	app.Log.Info().Msg("Connecting to ElasticSearch ...")
	elasticSearch, err := elasticsearch.NewClient(&config)
	if err != nil {
		return nil, err
	}
	app.ElasticSearch = elasticSearch

	app.Log.Info().Msg("Connecting to Rabbitmq ...")
	rabbitmqProducerConn, err := rabbitmq.NewConnection(&config)
	if err != nil {
		return nil, err
	}
	app.RabbitmqProducerConn = rabbitmqProducerConn

	rabbitmqConsumerConn, err := rabbitmq.NewConnection(&config)
	if err != nil {
		return nil, err
	}
	app.RabbitmqConsumerConn = rabbitmqConsumerConn

	app.GoogleOAuthConfig = &oauth2.Config{
		ClientID:     config.OAuth.Google.ClientId,
		ClientSecret: config.OAuth.Google.ClientSecret,
		RedirectURL:  config.OAuth.Google.RedirectUrl,
		Endpoint:     google.Endpoint,
		Scopes:       config.OAuth.Google.Scopes,
	}

	return app, nil
}
