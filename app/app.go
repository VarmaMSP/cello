package app

import (
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/elasticsearch"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/store/sqlstore"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type App struct {
	Store         store.Store
	ElasticSearch *elastic.Client

	GoogleOAuthConfig *oauth2.Config

	Log zerolog.Logger
}

func NewApp(config model.Config) (*App, error) {
	dev := true
	var log zerolog.Logger
	if dev {
		log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	} else {
		log = zerolog.New(os.Stdout).With().Timestamp().Logger()

	}

	log.Info().Msg("Connecting to Mysql ...")
	store, err := sqlstore.NewSqlStore(&config.Mysql)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Connecting to ElasticSearch ...")
	elasticSearch, err := elasticsearch.NewClient(&config.Elasticsearch)
	if err != nil {
		return nil, err
	}

	googleOAuthConfig := &oauth2.Config{
		ClientID:     config.OAuth.Google.ClientId,
		ClientSecret: config.OAuth.Google.ClientSecret,
		RedirectURL:  config.OAuth.Google.RedirectUrl,
		Endpoint:     google.Endpoint,
		Scopes:       config.OAuth.Google.Scopes,
	}

	return &App{
		Store:             store,
		ElasticSearch:     elasticSearch,
		GoogleOAuthConfig: googleOAuthConfig,
		Log:               log,
	}, nil
}
