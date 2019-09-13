package app

import (
	"fmt"

	"github.com/olivere/elastic/v7"
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
}

func NewApp(config model.Config) (*App, error) {
	store, err := sqlstore.NewSqlStore(&config.Mysql)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connecting to Mysql ...")
	}

	elasticSearch, err := elasticsearch.NewClient(&config.Elasticsearch)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connecting to Elasticsearch ...")
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
	}, nil
}
