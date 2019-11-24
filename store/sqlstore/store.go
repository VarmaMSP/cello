package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlStore interface {
	GetMaster() *sql.DB

	Insert(tableName string, models []model.DbModel) (sql.Result, error)
	InsertWithoutPK(tableName string, item model.DbModel) (int64, error)
	InsertOrUpdate(tableName string, m model.DbModel, updateSql string, updateValues ...interface{}) (sql.Result, error)

	UpdateChanges(tableName string, old, new model.DbModel, where string, values ...interface{}) (sql.Result, error)
	Query(copyTo func() []interface{}, sql string, values ...interface{}) error

	User() store.UserStore
	Feed() store.FeedStore
	Podcast() store.PodcastStore
	Episode() store.EpisodeStore
	Playlist() store.PlaylistStore
	Category() store.CategoryStore
	Task() store.TaskStore
}
