package app

import (
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/rs/zerolog"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/cache"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type App struct {
	Dev    bool
	Config *model.Config

	Store        store.Store
	SearchEngine searchengine.Broker
	MessageQueue messagequeue.Broker
	FileStorage  filestorage.Broker
	Cache        cache.Broker

	SessionManager *scs.SessionManager

	SyncPlaybackP messagequeue.Producer

	Log zerolog.Logger
}

func NewApp(store store.Store, se searchengine.Broker, mq messagequeue.Broker, fs filestorage.Broker, cache cache.Broker, log zerolog.Logger, config *model.Config) (*App, error) {
	app := &App{
		Dev:          config.Env == "dev",
		Config:       config,
		Store:        store,
		SearchEngine: se,
		MessageQueue: mq,
		FileStorage:  fs,
		Cache:        cache,
		Log:          log,
	}

	if p, err := mq.NewProducer(
		messagequeue.EXCHANGE_PHENOPOD_DIRECT,
		messagequeue.ROUTING_KEY_SYNC_PLAYBACK,
		config.Queues.RefreshPodcast.DeliveryMode,
	); err != nil {
		return nil, err
	} else {
		app.SyncPlaybackP = p
	}

	app.SessionManager = scs.New()
	app.SessionManager.Store = redisstore.New(cache.C())
	app.SessionManager.Lifetime = 30 * 24 * time.Hour

	return app, nil
}
