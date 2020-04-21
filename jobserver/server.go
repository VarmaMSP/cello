package jobserver

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type Job interface {
	Start()
}

type Server struct {
	store  store.Store
	se     searchengine.Broker
	mq     messagequeue.Broker
	fs     filestorage.Broker
	config *model.Config

	importPodcast   Job
	refreshPodcast  Job
	createThumbnail Job
	syncPlayback    Job
	taskRunner      Job
}

func NewJobServer(store store.Store, se searchengine.Broker, mq messagequeue.Broker, fs filestorage.Broker, config *model.Config) (*Server, error) {
	svr := &Server{
		store:  store,
		se:     se,
		mq:     mq,
		config: config,
	}

	if svr.config.Jobs.ImportPodcast.Enable {
		if job, err := NewImportPodcastJob(svr.store, svr.mq, svr.config); err != nil {
			return nil, err
		} else {
			svr.importPodcast = job
		}
	}

	if svr.config.Jobs.RefreshPodcast.Enable {
		if job, err := NewRefreshPodcastJob(svr.store, svr.mq, svr.config); err != nil {
			return nil, err
		} else {
			svr.refreshPodcast = job
		}
	}

	if svr.config.Jobs.SyncPlayback.Enable {
		if job, err := NewSyncPlaybackJob(svr.store, svr.mq, svr.config); err != nil {
			return nil, err
		} else {
			svr.syncPlayback = job
		}
	}

	if svr.config.Jobs.CreateThumbnail.Enable {
		if job, err := NewCreateThumbnailJob(svr.store, svr.mq, svr.fs, svr.config); err != nil {
			return nil, err
		} else {
			svr.createThumbnail = job
		}
	}

	if svr.config.Jobs.TaskScheduler.Enable {
		if job, err := NewTaskRunnerJob(svr.store, svr.se, svr.mq, svr.config); err != nil {
			return nil, err
		} else {
			svr.taskRunner = job
		}
	}

	return svr, nil
}

func (svr *Server) Start() {
	if svr.refreshPodcast != nil {
		svr.refreshPodcast.Start()
	}

	if svr.importPodcast != nil {
		svr.importPodcast.Start()
	}

	if svr.createThumbnail != nil {
		svr.createThumbnail.Start()
	}

	if svr.createThumbnail != nil {
		svr.createThumbnail.Start()
	}

	if svr.taskRunner != nil {
		svr.taskRunner.Start()
	}
}
