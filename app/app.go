package app

import (
	"os"
	"time"

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

	SyncEpisodePlaybackP *rabbitmq.Producer
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

	var err error

	app.Log.Info().Msg("Connecting to Mysql ...")
	app.Store, err = sqlstore.NewSqlStore(&config)
	if err != nil {
		return nil, err
	}

	app.Log.Info().Msg("Connecting to redis ...")
	app.Redis, err = NewRedisConnPool(&config)
	if err != nil {
		return nil, err
	}

	app.Log.Info().Msg("Connecting to ElasticSearch ...")
	app.ElasticSearch, err = elasticsearch.NewClient(&config)
	if err != nil {
		return nil, err
	}

	app.Log.Info().Msg("Connecting to Rabbitmq ...")
	app.RabbitmqProducerConn, err = rabbitmq.NewConnection(&config)
	if err != nil {
		return nil, err
	}
	app.RabbitmqConsumerConn, err = rabbitmq.NewConnection(&config)
	if err != nil {
		return nil, err
	}

	app.SessionManager = scs.New()
	app.SessionManager.Store = redisstore.New(app.Redis)
	app.SessionManager.Lifetime = 20 * 24 * time.Hour

	app.GoogleOAuthConfig = NewGoogleOAuthConfig(&config)
	app.FacebookOAuthConfig = NewFacebookOAuthConfig(&config)
	app.TwitterOAuthConfig = NewTwitterOAuthConfig(&config)

	app.SyncEpisodePlaybackP, err = rabbitmq.NewProducer(app.RabbitmqProducerConn, &rabbitmq.ProducerOpts{
		ExchangeName: rabbitmq.DefaultExchange,
		QueueName:    model.QUEUE_NAME_SYNC_EPISODE_PLAYBACK,
		DeliveryMode: config.Queues.SyncEpisodePlayback.DeliveryMode,
	})
	if err != nil {
		return nil, err
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

func NewGoogleOAuthConfig(config *model.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OAuth.Google.ClientId,
		ClientSecret: config.OAuth.Google.ClientSecret,
		RedirectURL:  config.OAuth.Google.RedirectUrl,
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       config.OAuth.Google.Scopes,
	}
}

func NewFacebookOAuthConfig(config *model.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OAuth.Facebook.ClientId,
		ClientSecret: config.OAuth.Facebook.ClientSecret,
		RedirectURL:  config.OAuth.Facebook.RedirectUrl,
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       config.OAuth.Facebook.Scopes,
	}
}

func NewTwitterOAuthConfig(config *model.Config) *oauth1.Config {
	return &oauth1.Config{
		ConsumerKey:    config.OAuth.Twitter.ClientId,
		ConsumerSecret: config.OAuth.Twitter.ClientSecret,
		CallbackURL:    config.OAuth.Twitter.RedirectUrl,
		Endpoint:       twitterOAuth.AuthorizeEndpoint,
	}
}
