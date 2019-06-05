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

	Insert(models []DbModel, tableName string) (sql.Result, error)

	Podcast() store.PodcastStore
	Episode() store.EpisodeStore
}
