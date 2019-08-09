package store

import "github.com/varmamsp/cello/model"

type Store interface {
	Podcast() PodcastStore
	Episode() EpisodeStore
	Category() CategoryStore
	ItunesMeta() ItunesMetaStore
	JobSchedule() JobScheduleStore
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
}

type CategoryStore interface {
	SavePodcastCategory(category *model.PodcastCategory) *model.AppError
}

type ItunesMetaStore interface {
	Save(meta *model.ItunesMeta) *model.AppError
	GetStatus(itunesId string) (string, *model.AppError)
	GetItunesIdList(offset, limit int) ([]string, *model.AppError)
	SetStatus(itunesId, status string) *model.AppError
}

type JobScheduleStore interface {
	GetAllActive() ([]*model.JobSchedule, *model.AppError)
	Disable(jobName string) *model.AppError
	SetRunAt(jobName string, runAt int64) *model.AppError
}
