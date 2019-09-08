package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/store"
)

type DbModel interface {
	DbColumns() []string
	FieldAddrs() []interface{}
}

type SqlStore interface {
	GetMaster() *sql.DB

	Insert(tableName string, models []DbModel) (sql.Result, error)
	UpdateChanges(tableName string, old, new DbModel, where string, values ...interface{}) (sql.Result, error)
	QueryRows(newItemFields func() []interface{}, sql string, values ...interface{}) error

	Podcast() store.PodcastStore
	Episode() store.EpisodeStore
	Category() store.CategoryStore
	ItunesMeta() store.ItunesMetaStore
	JobSchedule() store.JobScheduleStore
	PodcastCuration() store.PodcastCurationStore
}
