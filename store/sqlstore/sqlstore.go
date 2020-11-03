package sqlstore

import (
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store"
)

type sqlStore struct {
	feed         store.FeedStore
	podcast      store.PodcastStore
	episode      store.EpisodeStore
	category     store.CategoryStore
	user         store.UserStore
	playback     store.PlaybackStore
	subscription store.SubscriptionStore
	playlist     store.PlaylistStore
}

func NewSqlStore(broker sqldb.Broker) (store.Store, error) {
	s := &sqlStore{}
	s.feed = newSqlFeedStore(broker)
	s.podcast = newSqlPodcastStore(broker)
	s.episode = newSqlEpisodeStore(broker)
	s.category = newSqlCategoryStore(broker)
	s.user = newSqlUserStore(broker)
	s.subscription = newSqlSubscriptionStore(broker)
	s.playback = newSqlPlaybackStore(broker)
	s.playlist = newSqlPlaylistStore(broker)

	return s, nil
}

func (s *sqlStore) Feed() store.FeedStore {
	return s.feed
}

func (s *sqlStore) Podcast() store.PodcastStore {
	return s.podcast
}

func (s *sqlStore) Episode() store.EpisodeStore {
	return s.episode
}

func (s *sqlStore) Category() store.CategoryStore {
	return s.category
}

func (s *sqlStore) User() store.UserStore {
	return s.user
}

func (s *sqlStore) Playback() store.PlaybackStore {
	return s.playback
}

func (s *sqlStore) Subscription() store.SubscriptionStore {
	return s.subscription
}

func (s *sqlStore) Playlist() store.PlaylistStore {
	return s.playlist
}

func (s *sqlStore) File() store.FileStore {
	panic("method not implemented by sql layer")
}
