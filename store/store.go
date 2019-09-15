package store

import "github.com/varmamsp/cello/model"

type Store interface {
	Feed() FeedStore
	Podcast() PodcastStore
	Episode() EpisodeStore
	Category() CategoryStore
	Curation() CurationStore
	ItunesMeta() ItunesMetaStore
	JobSchedule() JobScheduleStore
}

type FeedStore interface {
	Save(feed *model.Feed) *model.AppError
	Get(id string) (*model.Feed, *model.AppError)
	GetAllBySource(source string, offset, limit int) ([]*model.Feed, *model.AppError)
	GetAllToBeRefreshed(createdAfter int64, limit int) ([]*model.Feed, *model.AppError)
	Update(old, new *model.Feed) *model.AppError
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
	Get(podcastId string) (*model.Podcast, *model.AppError)
	GetAllByCuration(curationId string, offset, limit int) ([]*model.Podcast, *model.AppError)
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
	Get(episodeId string) (*model.Episode, *model.AppError)
	GetAllByPodcast(podcastId string, limit, offset int) ([]*model.Episode, *model.AppError)
	Block(podcastId, episodeGuid string) *model.AppError
}

type CategoryStore interface {
	SavePodcastCategory(category *model.PodcastCategory) *model.AppError
}

type CurationStore interface {
	Save(curation *model.Curation) *model.AppError
	SavePodcastCuration(item *model.PodcastCuration) *model.AppError
	Get(curationId string) (*model.Curation, *model.AppError)
	GetAll() ([]*model.Curation, *model.AppError)
	Delete(curationId string) *model.AppError
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
