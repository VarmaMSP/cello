package jobserver

import (
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
)

type Job interface {
	Start()
}

type JobServer struct {
	store store.Store
	mq    messagequeue.Broker

	importPodcast   Job
	refreshPodcast  Job
	createThumbnail Job
	syncPlayback    Job
}
