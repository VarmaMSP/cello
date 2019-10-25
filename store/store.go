package store

import (
	"time"

	"github.com/varmamsp/cello/model"
)

type Store interface {
	User() UserStore
	Feed() FeedStore
	Podcast() PodcastStore
	Episode() EpisodeStore
	Category() CategoryStore
	Curation() CurationStore
	Task() TaskStore
}

type UserStore interface {
	Save(user *model.User) *model.AppError
	SaveSocialAccount(accountType string, account model.DbModel) *model.AppError
	Get(userId string) (*model.User, *model.AppError)
	GetSocialAccount(accountType, id string) (model.DbModel, *model.AppError)
}

type FeedStore interface {
	Save(feed *model.Feed) *model.AppError
	Get(id string) (*model.Feed, *model.AppError)
	GetBySource(source, sourceId string) (*model.Feed, *model.AppError)
	GetAllBySource(source string, offset, limit int) ([]*model.Feed, *model.AppError)
	GetAllToBeRefreshed(createdAfter int64, limit int) ([]*model.Feed, *model.AppError)
	Update(old, new *model.Feed) *model.AppError
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
	SaveSubscription(subscription *model.PodcastSubscription) *model.AppError
	Get(podcastId string) (*model.Podcast, *model.AppError)
	GetAllByCuration(curationId string, offset, limit int) ([]*model.Podcast, *model.AppError)
	GetAllSubscribedBy(userId string) ([]*model.Podcast, *model.AppError)
	DeleteSubscription(userId, podcastId string) *model.AppError
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
	SavePlayback(playback *model.EpisodePlayback) *model.AppError
	Get(episodeId string) (*model.Episode, *model.AppError)
	GetAllByIds(episodeIds []string) ([]*model.Episode, *model.AppError)
	GetAllByPodcast(podcastId string, limit, offset int) ([]*model.Episode, *model.AppError)
	GetAllPublishedBefore(podcastIds []string, before *time.Time, limit int) ([]*model.Episode, *model.AppError)
	GetAllPlaybacks(episodeIds []string, userId string) ([]*model.EpisodePlayback, *model.AppError)
	GetAllPlaybacksByUser(userId string) ([]*model.EpisodePlayback, *model.AppError)
	SetPlaybackCurrentTime(episodeId, userId string, currentTime int) *model.AppError
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

type TaskStore interface {
	GetAllActive() ([]*model.Task, *model.AppError)
	Update(old, new *model.Task) *model.AppError
}
