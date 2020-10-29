package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/job"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/cache"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/message_queue"
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

func New(config *model.Config) (*Server, error) {
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

	var mqBroker message_queue.Broker
	if mq, err := message_queue.NewBroker(config); err != nil {
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

	svr := &Server{}

	if app, err := app.NewApp(store, seBroker, mqBroker, fsBroker, cacheBroker, logger, config); err != nil {
		return nil, err
	} else {
		svr.httpSvr = &http.Server{
			Addr:    config.Adress,
			Handler: newRouter(app),
		}
	}

	if jobSvr, err := job.NewJobServer(store, seBroker, mqBroker, fsBroker, logger, config); err != nil {
		return nil, err
	} else {
		svr.jobSvr = jobSvr
	}

	return svr, nil
}

func (svr *Server) Start() {
	svr.jobSvr.Start()

	fmt.Println("Server Running on PORT :8080")
	err := svr.httpSvr.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
