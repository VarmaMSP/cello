package sqlstore

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
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

func NewSqlStore() SqlStore {
	supplier := &SqlSupplier{}
	supplier.db = InitDb()
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

func InitDb() *sql.DB {
	config := mysql.Config{}
	config.User = "root"
	config.Passwd = "tiracapeta"
	config.Addr = "localhost:3306"
	config.DBName = "phenopod"
	config.AllowNativePasswords = true
	config.Params = map[string]string{"collation": "utf8mb4_unicode_ci"}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	return db
}
