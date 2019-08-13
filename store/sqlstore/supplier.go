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
	supplier.itunesMeta = NewSqlItunesMetaStore(supplier)
	supplier.jobSchedule = NewSqlJobScheduleStore(supplier)

	return supplier
}

func (s *SqlSupplier) GetMaster() *sql.DB {
	return s.db
}

func (s *SqlSupplier) Insert(models []DbModel, tableName string) (sql.Result, error) {
	l := len(models)
	if l == 0 {
		return nil, nil
	}

	query := InsertQuery(tableName, models[0], len(models))
	values := make([]interface{}, 0)
	for i := range models {
		values = append(
			values,
			ValuesFromAddrs(models[i].FieldAddrs())...,
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

func (s *SqlSupplier) Category() store.CategoryStore {
	return s.category
}

func (s *SqlSupplier) ItunesMeta() store.ItunesMetaStore {
	return s.itunesMeta
}

func (s *SqlSupplier) JobSchedule() store.JobScheduleStore {
	return s.jobSchedule
}
