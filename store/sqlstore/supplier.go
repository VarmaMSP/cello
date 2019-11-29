package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlSupplier struct {
	db           *sql.DB
	user         store.UserStore
	feed         store.FeedStore
	podcast      store.PodcastStore
	subscription store.SubscriptionStore
	episode      store.EpisodeStore
	playback     store.PlaybackStore
	playlist     store.PlaylistStore
	category     store.CategoryStore
	task         store.TaskStore
}

func NewSqlStore(config *model.Config) (SqlStore, error) {
	db, err := sql.Open("mysql", MakeMysqlDSN(config))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	supplier := &SqlSupplier{}

	supplier.db = db
	supplier.user = NewSqlUserStore(supplier)
	supplier.feed = NewSqlFeedStore(supplier)
	supplier.podcast = NewSqlPodcastStore(supplier)
	supplier.subscription = NewSqlSubscriptionStore(supplier)
	supplier.episode = NewSqlEpisodeStore(supplier)
	supplier.playback = NewSqlPlaybackStore(supplier)
	supplier.playlist = NewSqlPlaylistStore(supplier)
	supplier.category = NewSqlCategoryStore(supplier)
	supplier.task = NewSqlTaskStore(supplier)

	return supplier, nil
}

func (s *SqlSupplier) GetMaster() *sql.DB {
	return s.db
}

func (s *SqlSupplier) Insert(tableName string, items []model.DbModel) (sql.Result, error) {
	query, insertValues, noValues := InsertQuery(tableName, items)
	if noValues {
		return nil, nil
	}

	return s.db.Exec(query, insertValues...)
}

func (s *SqlSupplier) InsertWithoutPK(tableName string, item model.DbModel) (int64, error) {
	query, insertValues, noValues := InsertQueryWithoutPK(tableName, item)
	if noValues {
		return 0, errors.New("no values in provided item")
	}

	res, err := s.db.Exec(query, insertValues...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *SqlSupplier) InsertOrUpdate(tableName string, m model.DbModel, updateSql string, updateValues ...interface{}) (sql.Result, error) {
	query, insertValues, _ := InsertQuery(tableName, []model.DbModel{m})

	return s.db.Exec(query+" ON DUPLICATE KEY UPDATE "+updateSql, append(insertValues, updateValues...)...)
}

func (s *SqlSupplier) UpdateChanges(tableName string, old, new model.DbModel, where string, values ...interface{}) (sql.Result, error) {
	query, updateValues, noChanges := UpdateQuery(tableName, old, new, where, values)
	if noChanges {
		return nil, nil
	}

	return s.db.Exec(query, updateValues...)
}

func (s *SqlSupplier) Query(copyTo func() []interface{}, sql string, values ...interface{}) error {
	rows, err := s.db.Query(sql, values...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(copyTo()...); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *SqlSupplier) User() store.UserStore {
	return s.user
}

func (s *SqlSupplier) Feed() store.FeedStore {
	return s.feed
}

func (s *SqlSupplier) Podcast() store.PodcastStore {
	return s.podcast
}

func (s *SqlSupplier) Subscription() store.SubscriptionStore {
	return s.subscription
}

func (s *SqlSupplier) Episode() store.EpisodeStore {
	return s.episode
}

func (s *SqlSupplier) Playback() store.PlaybackStore {
	return s.playback
}

func (s *SqlSupplier) Playlist() store.PlaylistStore {
	return s.playlist
}

func (s *SqlSupplier) Category() store.CategoryStore {
	return s.category
}

func (s *SqlSupplier) Task() store.TaskStore {
	return s.task
}
