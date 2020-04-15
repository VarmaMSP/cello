package sqlstore_

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlStore interface {
	getMaster() *sql.DB
}

type sqlSupplier struct {
	db *sql.DB

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

func NewSqlStore(config *model.Config) (store_.Store, error) {
	db, err := sql.Open("mysql", makeMysqlDSN(config))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	supplier := &sqlSupplier{}

	supplier.db = db
	supplier.feed = newSqlFeedStore(supplier)
	supplier.podcast = newSqlPodcastStore(supplier)
	supplier.episode = newSqlEpisodeStore(supplier)
	supplier.category = newSqlCategoryStore(supplier)
	supplier.task = newSqlTaskStore(supplier)
	supplier.user = newSqlUserStore(supplier)
	supplier.subscription = newSqlSubscriptionStore(supplier)
	supplier.playback = newSqlPlaybackStore(supplier)
	supplier.playlist = newSqlPlaylistStore(supplier)

	return supplier, nil
}

// sqlstore implementation
func (s *sqlSupplier) getMaster() *sql.DB {
	return s.db
}

// Store implementation
func (s *sqlSupplier) Feed() store_.FeedStore {
	return s.feed
}

func (s *sqlSupplier) Podcast() store_.PodcastStore {
	return s.podcast
}

func (s *sqlSupplier) Episode() store_.EpisodeStore {
	return s.episode
}

func (s *sqlSupplier) Category() store_.CategoryStore {
	return s.category
}

func (s *sqlSupplier) Task() store_.TaskStore {
	return s.task
}

func (s *sqlSupplier) User() store_.UserStore {
	return s.user
}

func (s *sqlSupplier) Playback() store_.PlaybackStore {
	return s.playback
}

func (s *sqlSupplier) Subscription() store_.SubscriptionStore {
	return s.subscription
}

func (s *sqlSupplier) Playlist() store_.PlaylistStore {
	return s.playlist
}

func makeMysqlDSN(config *model.Config) string {
	c := mysql.NewConfig()
	c.Net = "tcp"
	c.Addr = config.Mysql.Address
	c.DBName = config.Mysql.Database
	c.User = config.Mysql.User
	c.Passwd = config.Mysql.Password
	c.AllowNativePasswords = true
	c.Params = map[string]string{"collation": "utf8mb4_unicode_ci"}
	c.ReadTimeout = 2 * time.Minute
	c.WriteTimeout = 2 * time.Minute

	return c.FormatDSN()
}
