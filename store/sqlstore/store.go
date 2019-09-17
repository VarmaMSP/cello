package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlStore interface {
	GetMaster() *sql.DB

	Insert(tableName string, models []model.DbModel) (sql.Result, error)
	UpdateChanges(tableName string, old, new model.DbModel, where string, values ...interface{}) (sql.Result, error)
	Query(copyTo func() []interface{}, sql string, values ...interface{}) error

	User() store.UserStore
	Feed() store.FeedStore
	Podcast() store.PodcastStore
	Episode() store.EpisodeStore
	Category() store.CategoryStore
	Curation() store.CurationStore
	Task() store.TaskStore
}
