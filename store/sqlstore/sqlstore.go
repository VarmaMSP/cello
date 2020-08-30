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
	task         store.TaskStore
	user         store.UserStore
	playback     store.PlaybackStore
	subscription store.SubscriptionStore
	playlist     store.PlaylistStore
}

func NewSqlStore(broker sqldb.Broker) (store.Store, error) {
	s := &sqlStore{}
	s.feed = &sqlFeedStore{Broker: broker}
	s.podcast = &sqlPodcastStore{Broker: broker}
	s.episode = newSqlEpisodeStore(broker)
	s.category = newSqlCategoryStore(broker)
	s.task = &sqlTaskStore{Broker: broker}
	s.user = &sqlUserStore{Broker: broker}
	s.subscription = &sqlSubscriptionStore{Broker: broker}
	s.playback = &sqlPlaybackStore{Broker: broker}
	s.playlist = &sqlPlaylistStore{Store: s, Broker: broker}

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

func (s *sqlStore) Task() store.TaskStore {
	return s.task
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
