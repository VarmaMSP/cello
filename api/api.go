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
	App *app.App

	Server *http.Server
	Router *httprouter.Router

	SchedulerJob           model.Job
	ImportPodcastJob       model.Job
	RefreshPodcastJob      model.Job
	CreateThumbnailJob     model.Job
	SyncEpisodePlaybackJob model.Job
}

func NewApi(config model.Config) (*Api, error) {
	api := &Api{}

	api.Router = httprouter.New()
	api.Server = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: api.Router,
	}

	app, err := app.NewApp(config)
	if err != nil {
		return nil, err
	}
	api.App = app

	if config.Jobs.Scheduler.Enable {
		api.App.Log.Info().Msg("Starting scheduler job...")
		job, err := job.NewSchedulerJob(api.App, &config)
		if err != nil {
			return nil, err
		}
		api.SchedulerJob = job
		go api.SchedulerJob.Run()
	}

	if config.Jobs.ImportPodcast.Enable {
		api.App.Log.Info().Msg("Starting import podcast job...")
		job, err := job.NewImportPodcastJob(api.App, &config)
		if err != nil {
			return nil, err
		}
		api.ImportPodcastJob = job
		go api.ImportPodcastJob.Run()
	}

	if config.Jobs.RefreshPodcast.Enable {
		api.App.Log.Info().Msg("Starting refresh podcast job...")
		job, err := job.NewRefreshPodcastJob(api.App, &config)
		if err != nil {
			return nil, err
		}
		api.RefreshPodcastJob = job
		go api.RefreshPodcastJob.Run()
	}

	if config.Jobs.CreateThumbnail.Enable {
		api.App.Log.Info().Msg("Starting create thumbnail job...")
		job, err := job.NewCreateThumbnailJob(api.App, &config)
		if err != nil {
			return nil, err
		}
		api.CreateThumbnailJob = job
		go api.CreateThumbnailJob.Run()
	}

	if config.Jobs.SyncPlayback.Enable {
		api.App.Log.Info().Msg("Starting sync playback job...")
		job, err := job.NewSyncPlaybackJob(api.App, &config)
		if err != nil {
			return nil, err
		}
		api.SyncEpisodePlaybackJob = job
		go api.SyncEpisodePlaybackJob.Run()
	}

	api.RegisterHandlers()

	return api, nil
}

func (api *Api) ListenAndServe() {
	api.App.Log.Info().Msg("Server listening on port: 8081")
	err := api.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
