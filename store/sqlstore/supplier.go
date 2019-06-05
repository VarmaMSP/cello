package sqlstore

import (
	"database/sql"

	"github.com/varmamsp/cello/store"
)

type SqlSupplier struct {
	db      *sql.DB
	podcast store.PodcastStore
	episode store.EpisodeStore
}

func (s *SqlSupplier) GetMaster() *sql.DB {
	return s.db
}

func (s *SqlSupplier) Insert(models []DbModel, tableName string) (sql.Result, error) {
	query := InsertQuery(tableName, models[0], len(models))
	values := make([]interface{}, 0)
	for i := range models {
		values = append(
			values,
			ValuesFromAddrs(models[i].FieldAddrs()),
		)
	}
	return s.db.Exec(query, values...)
}

func (s *SqlSupplier) Podcast() store.PodcastStore {
	return s.podcast
}

func (s *SqlSupplier) Episode() store.EpisodeStore {
	return s.episode
}
