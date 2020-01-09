package store

import (
	"github.com/varmamsp/cello/model"
)

type Store interface {
	User() UserStore
	Feed() FeedStore
	Podcast() PodcastStore
	Subscription() SubscriptionStore
	Episode() EpisodeStore
	Playback() PlaybackStore
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
	GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError)
	GetBySourcePaginated(source string, offset, limit int) ([]*model.Feed, *model.AppError)
	GetForRefreshPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError)
	Update(old, new *model.Feed) *model.AppError
}

type PodcastStore interface {
	Save(podcast *model.Podcast) *model.AppError
	Get(podcastId int64) (*model.Podcast, *model.AppError)
	GetByIds(podcastIds []int64) ([]*model.Podcast, *model.AppError)
	GetSubscriptions(userId int64) ([]*model.Podcast, *model.AppError)
	Update(old, new *model.Podcast) *model.AppError
}

type SubscriptionStore interface {
	Save(subscription *model.Subscription) *model.AppError
	Delete(userId, podcastId int64) *model.AppError
}

type EpisodeStore interface {
	Save(episode *model.Episode) *model.AppError
	Get(episodeId int64) (*model.Episode, *model.AppError)
	GetByIds(episodeIds []int64) ([]*model.Episode, *model.AppError)
	GetByPodcast(podcastId int64) ([]*model.Episode, *model.AppError)
	GetByPodcastPaginated(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError)
	GetByPodcastIdsPaginated(podcastIds []int64, offset, limit int) ([]*model.Episode, *model.AppError)
	GetByPlaylistPaginated(playlistId int64, offset, limit int) ([]*model.Episode, *model.AppError)
	Block(episodeIds []int64) *model.AppError
}

type PlaybackStore interface {
	Save(playback *model.Playback) *model.AppError
	GetByUserPaginated(userId int64, offset, limit int) ([]*model.Playback, *model.AppError)
	GetByUserByEpisodes(userId int64, episodeIds []int64) ([]*model.Playback, *model.AppError)
	Update(progress *model.PlaybackProgress) *model.AppError
}

type PlaylistStore interface {
	Save(playlist *model.Playlist) *model.AppError
	Get(playlistId int64) (*model.Playlist, *model.AppError)
	GetByUser(userId int64) ([]*model.Playlist, *model.AppError)
	GetByUserPaginated(userId int64, offset, limit int) ([]*model.Playlist, *model.AppError)
	Update(old, new *model.Playlist) *model.AppError
	UpdateMemberStats(playlistId int64) *model.AppError
	Delete(playlistId int64) *model.AppError
	SaveMember(member *model.PlaylistMember) *model.AppError
	GetMembers(playlistIds, episodeIds []int64) ([]*model.PlaylistMember, *model.AppError)
	GetMembersByPlaylist(playlist int64) ([]*model.PlaylistMember, *model.AppError)
	ChangeMemberPosition(playlistId, episodeId int64, from, to int) *model.AppError
	DeleteMember(playlistId, episodeId int64) *model.AppError
}

type CategoryStore interface {
	SavePodcastCategory(category *model.PodcastCategory) *model.AppError
}

type TaskStore interface {
	GetAll() ([]*model.Task, *model.AppError)
	Update(old, new *model.Task) *model.AppError
}
