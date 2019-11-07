package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlSupplier struct {
	db       *sql.DB
	user     store.UserStore
	feed     store.FeedStore
	podcast  store.PodcastStore
	episode  store.EpisodeStore
	playlist store.PlaylistStore
	category store.CategoryStore
	curation store.CurationStore
	task     store.TaskStore
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
	supplier.episode = NewSqlEpisodeStore(supplier)
	supplier.playlist = NewSqlPlaylistStore(supplier)
	supplier.category = NewSqlCategoryStore(supplier)
	supplier.curation = NewSqlCurationStore(supplier)
	supplier.task = NewSqlTaskStore(supplier)

	return supplier, nil
}

func (s *SqlSupplier) GetMaster() *sql.DB {
	return s.db
}

func (s *SqlSupplier) Insert(tableName string, models []model.DbModel) (sql.Result, error) {
	query, insertValues, noValues := InsertQuery(tableName, models)
	if noValues {
		return nil, nil
	}

	return s.db.Exec(query, insertValues...)
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

func (s *SqlSupplier) Episode() store.EpisodeStore {
	return s.episode
}

func (s *SqlSupplier) Playlist() store.PlaylistStore {
	return s.playlist
}

func (s *SqlSupplier) Category() store.CategoryStore {
	return s.category
}

func (s *SqlSupplier) Curation() store.CurationStore {
	return s.curation
}

func (s *SqlSupplier) Task() store.TaskStore {
	return s.task
}
