package sqlstore_

import (
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store_"
)

type sqlStore struct {
	feed         store_.FeedStore
	podcast      store_.PodcastStore
	episode      store_.EpisodeStore
	category     store_.CategoryStore
	task         store_.TaskStore
	user         store_.UserStore
	playback     store_.PlaybackStore
	subscription store_.SubscriptionStore
	playlist     store_.PlaylistStore
}

func NewSqlStore(broker sqldb.Broker) (store_.Store, error) {
	s := &sqlStore{}
	s.feed = newSqlFeedStore(broker)
	s.podcast = newSqlPodcastStore(broker)
	s.episode = newSqlEpisodeStore(broker)
	s.category = newSqlCategoryStore(broker)
	s.task = newSqlTaskStore(broker)
	s.user = newSqlUserStore(broker)
	s.subscription = newSqlSubscriptionStore(broker)
	s.playback = newSqlPlaybackStore(broker)
	s.playlist = newSqlPlaylistStore(broker)

	return s, nil
}

func (s *sqlStore) Feed() store_.FeedStore {
	return s.feed
}

func (s *sqlStore) Podcast() store_.PodcastStore {
	return s.podcast
}

func (s *sqlStore) Episode() store_.EpisodeStore {
	return s.episode
}

func (s *sqlStore) Category() store_.CategoryStore {
	return s.category
}

func (s *sqlStore) Task() store_.TaskStore {
	return s.task
}

func (s *sqlStore) User() store_.UserStore {
	return s.user
}

func (s *sqlStore) Playback() store_.PlaybackStore {
	return s.playback
}

func (s *sqlStore) Subscription() store_.SubscriptionStore {
	return s.subscription
}

func (s *sqlStore) Playlist() store_.PlaylistStore {
	return s.playlist
}