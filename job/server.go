package job

import (
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/message_queue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type Job interface {
	Start()
}

type Server struct {
	log    zerolog.Logger
	config *model.Config

	importPodcast   Job
	refreshPodcast  Job
	createThumbnail Job
	syncPlayback    Job
}

func NewJobServer(store store.Store, se searchengine.Broker, mq message_queue.Broker, fs filestorage.Broker, log zerolog.Logger, config *model.Config) (*Server, error) {
	svr := &Server{
		log:    log,
		config: config,
	}

	if svr.config.Jobs.ImportPodcast.Enable {
		if job, err := NewImportPodcastJob(store, mq, svr.log, svr.config); err != nil {
			return nil, err
		} else {
			svr.importPodcast = job
		}
	}

	if svr.config.Jobs.RefreshPodcast.Enable {
		if job, err := NewRefreshPodcastJob(store, mq, svr.log, svr.config); err != nil {
			return nil, err
		} else {
			svr.refreshPodcast = job
		}
	}

	if svr.config.Jobs.SyncPlayback.Enable {
		if job, err := NewSyncPlaybackJob(store, mq, svr.log, svr.config); err != nil {
			return nil, err
		} else {
			svr.syncPlayback = job
		}
	}

	if svr.config.Jobs.CreateThumbnail.Enable {
		if job, err := NewCreateThumbnailJob(store, mq, fs, svr.log, svr.config); err != nil {
			return nil, err
		} else {
			svr.createThumbnail = job
		}
	}

	// if svr.config.Jobs.TaskScheduler.Enable {
	// 	if job, err := (store, se, mq, svr.log, svr.config); err != nil {
	// 		return nil, err
	// 	} else {
	// 		svr.taskRunner = job
	// 	}
	// }

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

	if svr.syncPlayback != nil {
		svr.syncPlayback.Start()
	}
}
