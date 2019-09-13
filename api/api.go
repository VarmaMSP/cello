package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/store"
	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"
)

type Api struct {
	store    store.Store
	esClient *elastic.Client

	googleOAuthConfig *oauth2.Config

	enableCors bool

	server *http.Server
	router *httprouter.Router
}

func NewApi(store store.Store, client *elastic.Client) *Api {
	api := &Api{
		store:    store,
		esClient: client,

		googleOAuthConfig: &oauth2.Config{
			ClientID:     "",
			ClientSecret: "",
			RedirectURL:  "http://localhost:8080/api/google/callback",
			Endpoint:     googleOAuth2.Endpoint,
			Scopes:       []string{"profile", "email"},
		},

		router: httprouter.New(),
	}

	api.server = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: api.router,
	}

	api.RegisterPodcastHandlers()
	api.RegisterCurationRoutes()
	api.RegisterLoginHandlers()

	return api
}

func (api *Api) ListenAndServe() {
	err := api.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
