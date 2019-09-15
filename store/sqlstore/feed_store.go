package sqlstore

import (
	"net/http"
	"strings"

	"github.com/varmamsp/cello/model"
)

type SqlFeedStore struct {
	SqlStore
}

func NewSqlFeedStore(store SqlStore) *SqlFeedStore {
	return &SqlFeedStore{store}
}

func (s *SqlFeedStore) Save(feed *model.Feed) *model.AppError {
	feed.PreSave()

	if _, err := s.Insert("feed", []DbModel{feed}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_feed_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"source": feed.Source, "source_id": feed.SourceId},
		)
	}
	return nil
}

func (s *SqlFeedStore) Get(id string) (*model.Feed, *model.AppError) {
	feed := &model.Feed{}
	sql := "SELECT " + strings.Join(feed.DbColumns(), ",") + " FROM feed WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(feed.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_feed_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return feed, nil
}

func (s *SqlFeedStore) GetAllBySource(source string, offset, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := "SELECT " + strings.Join((&model.Feed{}).DbColumns(), ",") + " FROM feed WHERE source = ? LIMIT ?, ?"

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, source, offset, limit); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_feed_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"source": source},
		)
	}
	return
}

func (s *SqlFeedStore) GetAllToBeRefreshed(createdAfter int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := "SELECT " + strings.Join((&model.Feed{}).DbColumns(), ",") + ` FROM feed
		WHERE refresh_enabled = 1 AND
			  next_refresh_at > ? AND
			  created_at > ?
		ORDER BY created_at LIMIT ?`

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, model.Now(), createdAfter, limit); err != nil {
		appE = model.NewAppError("store.sqlstore.sql_feed_store.get_all_to_be_refreshed", err.Error(), http.StatusInternalServerError, nil)
	}
	return
}

func (s *SqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	if _, err := s.UpdateChanges("feed", old, new, "id = ?", new.Id); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_feed_store.update", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": new.Id},
		)
	}
	return nil
}
