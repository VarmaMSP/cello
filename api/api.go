package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/job"
	"github.com/varmamsp/cello/model"
)

type Api struct {
	app *app.App

	server *http.Server
	router *httprouter.Router

	SchedulerJob           model.Job
	ImportPodcastJob       model.Job
	RefreshPodcastJob      model.Job
	CreateThumbnailJob     model.Job
	SyncEpisodePlaybackJob model.Job
}

func NewApi(config model.Config) (*Api, error) {
	api := &Api{}

	api.router = httprouter.New()
	api.server = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: api.router,
	}

	app, err := app.NewApp(config)
	if err != nil {
		return nil, err
	}
	api.app = app

	if config.Jobs.Scheduler.Enable {
		api.app.Log.Info().Msg("Starting scheduler job...")
		job, err := job.NewSchedulerJob(api.app, &config)
		if err != nil {
			return nil, err
		}
		api.SchedulerJob = job
		go api.SchedulerJob.Run()
	}

	if config.Jobs.ImportPodcast.Enable {
		api.app.Log.Info().Msg("Starting import podcast job...")
		job, err := job.NewImportPodcastJob(api.app, &config)
		if err != nil {
			return nil, err
		}
		api.ImportPodcastJob = job
		go api.ImportPodcastJob.Run()
	}

	if config.Jobs.RefreshPodcast.Enable {
		api.app.Log.Info().Msg("Starting refresh podcast job...")
		job, err := job.NewRefreshPodcastJob(api.app, &config)
		if err != nil {
			return nil, err
		}
		api.RefreshPodcastJob = job
		go api.RefreshPodcastJob.Run()
	}

	if config.Jobs.CreateThumbnail.Enable {
		api.app.Log.Info().Msg("Starting create thumbnail job...")
		job, err := job.NewCreateThumbnailJob(api.app, &config)
		if err != nil {
			return nil, err
		}
		api.CreateThumbnailJob = job
		go api.CreateThumbnailJob.Run()
	}

	if config.Jobs.SyncPlayback.Enable {
		api.app.Log.Info().Msg("Starting sync playback job...")
		job, err := job.NewSyncPlaybackJob(api.app, &config)
		if err != nil {
			return nil, err
		}
		api.SyncEpisodePlaybackJob = job
		go api.SyncEpisodePlaybackJob.Run()
	}

	api.RegisterSearchHandlers()
	api.RegisterPodcastHandlers()
	api.RegisterSubscriptionHandlers()
	api.RegisterEpisodeHandlers()
	api.RegisterPlaybackHandlers()
	api.RegisterHistoryHandlers()
	api.RegisterPlaylistHandlers()
	api.RegisterUserHandlers()

	return api, nil
}

func (api *Api) ListenAndServe() {
	api.app.Log.Info().Msg("Server listening on port: 8081")
	err := api.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
