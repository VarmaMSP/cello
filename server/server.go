package server

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/varmamsp/cello/app_"
	"github.com/varmamsp/cello/job"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/cache"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/store/searchlayer"
	"github.com/varmamsp/cello/store/sqlstore"
)

type Server struct {
	httpSvr *http.Server
	jobSvr  *job.Server
}

func NewServer(config *model.Config) (*Server, error) {
	var store store.Store
	if db, err := sqldb.NewBroker(config); err != nil {
		return nil, err
	} else if s, err := sqlstore.NewSqlStore(db); err != nil {
		return nil, err
	} else {
		store = s
	}

	var seBroker searchengine.Broker
	if se, err := searchengine.NewBroker(config); err != nil {
		return nil, err
	} else {
		store = searchlayer.NewSearchLayer(store, se)
		seBroker = se
	}

	var mqBroker messagequeue.Broker
	if mq, err := messagequeue.NewBroker(config); err != nil {
		return nil, err
	} else {
		mqBroker = mq
	}

	var fsBroker filestorage.Broker
	if fs, err := filestorage.NewBroker(config); err != nil {
		return nil, err
	} else {
		fsBroker = fs
	}

	var cacheBroker cache.Broker
	if cache, err := cache.NewBroker(config); err != nil {
		return nil, err
	} else {
		cacheBroker = cache
	}

	var logger zerolog.Logger
	if config.Env == "dev" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	app := app_.NewApp(store, seBroker, mqBroker, fsBroker, cacheBroker, logger, config)
	job, _ := job.NewJobServer(store, seBroker, mqBroker, fsBroker, config)

	return &Server{
		httpSvr: &http.Server{
			Addr:    "127.0.0.1:8081",
			Handler: newRouter(app),
		},
		jobSvr: job,
	}, nil
}

func (svr *Server) ListenAndServe() {

	// err := api.Server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
