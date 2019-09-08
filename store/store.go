package store

import "github.com/varmamsp/cello/model"

type Store interface {
	Podcast() PodcastStore
	Episode() EpisodeStore
	Category() CategoryStore
	Curation() CurationStore
	ItunesMeta() ItunesMetaStore
	JobSchedule() JobScheduleStore
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
	GetInfo(podcastId string) (*model.PodcastInfo, *model.AppError)
	GetAllToBeRefreshed(createdAfter int64, limit int) ([]*model.PodcastFeedDetails, *model.AppError)
	UpdateFeedDetails(old, new *model.PodcastFeedDetails) *model.AppError
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
	GetInfo(id string) (*model.EpisodeInfo, *model.AppError)
	GetAllByPodcast(podcastId string, limit, offset int) ([]*model.EpisodeInfo, *model.AppError)
	GetAllGuidsByPodcast(podcastId string) ([]string, *model.AppError)
	Block(podcastId, episodeGuid string) *model.AppError
}

type CategoryStore interface {
	SavePodcastCategory(category *model.PodcastCategory) *model.AppError
}

type CurationStore interface {
	Save(curation *model.Curation) *model.AppError
	SavePodcastCuration(item *model.PodcastCuration) *model.AppError
	GetAll() ([]*model.Curation, *model.AppError)
	GetPodcastsByCuration(curationId string, offset, limit int) ([]*model.PodcastInfo, *model.AppError)
}

type ItunesMetaStore interface {
	Save(meta *model.ItunesMeta) *model.AppError
	Update(old, new *model.ItunesMeta) *model.AppError
	GetItunesIdList(offset, limit int) ([]string, *model.AppError)
}

type JobScheduleStore interface {
	GetAllActive() ([]*model.JobSchedule, *model.AppError)
	Disable(jobName string) *model.AppError
	SetRunAt(jobName string, runAt int64) *model.AppError
}
