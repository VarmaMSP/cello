package store

import (
	"github.com/varmamsp/cello/model"
)

type Store interface {
	User() UserStore
	Feed() FeedStore
	Podcast() PodcastStore
	Episode() EpisodeStore
	Playlist() PlaylistStore
	Category() CategoryStore
	Task() TaskStore
}

type UserStore interface {
	Save(user *model.User) *model.AppError
	SaveSocialAccount(accountType string, account model.DbModel) *model.AppError
	Get(userId int64) (*model.User, *model.AppError)
	GetSocialAccount(accountType, id string) (model.DbModel, *model.AppError)
}

type FeedStore interface {
	Save(feed *model.Feed) *model.AppError
	Get(feedId int64) (*model.Feed, *model.AppError)
	GetBySource(source, sourceId string) (*model.Feed, *model.AppError)
	GetAllBySource(source string, offset, limit int) ([]*model.Feed, *model.AppError)
	GetAllToBeRefreshed(createdAfter int64, limit int) ([]*model.Feed, *model.AppError)
	Update(old, new *model.Feed) *model.AppError
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
	SaveSubscription(subscription *model.PodcastSubscription) *model.AppError
	Get(podcastId int64) (*model.Podcast, *model.AppError)
	GetAllSubscribedBy(userId int64) ([]*model.Podcast, *model.AppError)
	DeleteSubscription(userId, podcastId int64) *model.AppError
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
	SavePlayback(playback *model.EpisodePlayback) *model.AppError
	Get(episodeId int64) (*model.Episode, *model.AppError)
	GetAllByIds(episodeIds []int64) ([]*model.Episode, *model.AppError)
	GetAllByPodcast(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError)
	GetAllPublishedBefore(podcastIds []int64, offset, limit int) ([]*model.Episode, *model.AppError)
	GetAllPlaybacks(episodeIds []int64, userId int64) ([]*model.EpisodePlayback, *model.AppError)
	GetAllPlaybacksByUser(userId int64, offset, limit int) ([]*model.EpisodePlayback, *model.AppError)
	SetPlaybackCurrentTime(episodeId, userId int64, currentTime int) *model.AppError
	Block(podcastId int64, episodeGuid string) *model.AppError
}

type PlaylistStore interface {
	Save(playlist *model.Playlist) *model.AppError
	SaveItem(playlistItem *model.PlaylistItem) *model.AppError
	Get(playlistId int64) (*model.Playlist, *model.AppError)
	GetAllByUser(userId int64) ([]*model.Playlist, *model.AppError)
	GetAllEpisodesInPlaylist(playlistId int64) ([]*model.Episode, *model.AppError)
}

type CategoryStore interface {
	SavePodcastCategory(category *model.PodcastCategory) *model.AppError
}

type TaskStore interface {
	GetAll() ([]*model.Task, *model.AppError)
	Update(old, new *model.Task) *model.AppError
}
