package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/store"
)

type DbModel interface {
	GetDbColumns() []string
	GetFieldAddrs() []interface{}
}

type SqlStore interface {
	GetMaster() *sql.DB

	Insert(model DbModel, tableName string)
	BulkInsert(models []DbModel, tableName string)

	Podcast() store.PodcastStore
	Episode() store.EpisodeStore
}
