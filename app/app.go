package app

import (
	"os"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/dghubble/oauth1"
	twitterOAuth "github.com/dghubble/oauth1/twitter"
	"github.com/gomodule/redigo/redis"
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
	facebookOAuth "golang.org/x/oauth2/facebook"
	googleOAuth "golang.org/x/oauth2/google"
)

type App struct {
	HostName string

	Store                store.Store
	Redis                *redis.Pool
	ElasticSearch        *elastic.Client
	RabbitmqProducerConn *amqp.Connection
	RabbitmqConsumerConn *amqp.Connection

	SessionManager      *scs.SessionManager
	GoogleOAuthConfig   *oauth2.Config
	FacebookOAuthConfig *oauth2.Config
	TwitterOAuthConfig  *oauth1.Config

	Log zerolog.Logger
}

func NewApp(config model.Config) (*App, error) {
	app := &App{}

	if config.Env == "dev" {
		app.Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
		app.HostName = "http://localhost:8080"
	} else {
		app.Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
		app.HostName = "https://phenopod.com"
	}

	app.Log.Info().Msg("Connecting to Mysql ...")
	store, err := sqlstore.NewSqlStore(&config)
	if err != nil {
		return nil, err
	}
	app.Store = store

	app.Log.Info().Msg("Connecting to redis ...")
	redisConnPool, err := NewRedisConnPool(&config)
	if err != nil {
		return nil, err
	}
	app.Redis = redisConnPool

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
	rabbitmqConsumerConn, err := rabbitmq.NewConnection(&config)
	if err != nil {
		return nil, err
	}
	app.RabbitmqProducerConn = rabbitmqProducerConn
	app.RabbitmqConsumerConn = rabbitmqConsumerConn

	app.SessionManager = scs.New()
	app.SessionManager.Store = redisstore.New(app.Redis)

	app.GoogleOAuthConfig = &oauth2.Config{
		ClientID:     config.OAuth.Google.ClientId,
		ClientSecret: config.OAuth.Google.ClientSecret,
		RedirectURL:  config.OAuth.Google.RedirectUrl,
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       config.OAuth.Google.Scopes,
	}

	app.FacebookOAuthConfig = &oauth2.Config{
		ClientID:     config.OAuth.Facebook.ClientId,
		ClientSecret: config.OAuth.Facebook.ClientSecret,
		RedirectURL:  config.OAuth.Facebook.RedirectUrl,
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       config.OAuth.Facebook.Scopes,
	}

	app.TwitterOAuthConfig = &oauth1.Config{
		ConsumerKey:    config.OAuth.Twitter.ClientId,
		ConsumerSecret: config.OAuth.Twitter.ClientSecret,
		CallbackURL:    config.OAuth.Twitter.RedirectUrl,
		Endpoint:       twitterOAuth.AuthorizeEndpoint,
	}

	return app, nil
}

func NewRedisConnPool(config *model.Config) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle: config.Redis.MaxIdleConn,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Redis.Address)
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return nil, err
	}
	return pool, nil
}
