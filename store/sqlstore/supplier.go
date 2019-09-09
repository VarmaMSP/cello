package sqlstore

import (
	"database/sql"
	"log"
	"os"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlSupplier struct {
	db          *sql.DB
	podcast     store.PodcastStore
	episode     store.EpisodeStore
	category    store.CategoryStore
	curation    store.CurationStore
	itunesMeta  store.ItunesMetaStore
	jobSchedule store.JobScheduleStore
}

func NewSqlStore(mysqlConfig *model.MysqlConfig) SqlStore {
	db, err := sql.Open("mysql", MakeMysqlDSN(mysqlConfig))
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	supplier := &SqlSupplier{}

	supplier.db = db
	supplier.podcast = NewSqlPodcastStore(supplier)
	supplier.episode = NewSqlEpisodeStore(supplier)
	supplier.category = NewSqlCategoryStore(supplier)
	supplier.curation = NewSqlCurationStore(supplier)
	supplier.itunesMeta = NewSqlItunesMetaStore(supplier)
	supplier.jobSchedule = NewSqlJobScheduleStore(supplier)

	return supplier
}

func (s *SqlSupplier) GetMaster() *sql.DB {
	return s.db
}

func (s *SqlSupplier) Insert(tableName string, models []DbModel) (sql.Result, error) {
	query, insertValues, noValues := InsertQuery(tableName, models)
	if noValues {
		return nil, nil
	}

	return s.db.Exec(query, insertValues...)
}

func (s *SqlSupplier) UpdateChanges(tableName string, old, new DbModel, where string, values ...interface{}) (sql.Result, error) {
	query, updateValues, noChanges := UpdateQuery("itunes_meta", old, new, " WHERE itunes_id = ?", values)
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

func (s *SqlSupplier) Podcast() store.PodcastStore {
	return s.podcast
}

func (s *SqlSupplier) Episode() store.EpisodeStore {
	return s.episode
}

func (s *SqlSupplier) Category() store.CategoryStore {
	return s.category
}

func (s *SqlSupplier) Curation() store.CurationStore {
	return s.curation
}

func (s *SqlSupplier) ItunesMeta() store.ItunesMetaStore {
	return s.itunesMeta
}

func (s *SqlSupplier) JobSchedule() store.JobScheduleStore {
	return s.jobSchedule
}
