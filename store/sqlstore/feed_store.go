package sqlstore

import (
	"net/http"

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

	id, err := s.InsertWithoutPK("feed", feed)
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_feed_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": feed.Source, "source_id": feed.SourceId},
		)
	}
	feed.Id = id
	return nil
}

func (s *SqlFeedStore) Get(feedId int64) (*model.Feed, *model.AppError) {
	feed := &model.Feed{}
	sql := "SELECT " + Cols(feed) + " FROM feed WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, feedId).Scan(feed.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_feed_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"id": feed},
		)
	}
	return feed, nil
}

func (s *SqlFeedStore) GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError) {
	feed := &model.Feed{}
	sql := "SELECT " + Cols(feed) + " FROM feed WHERE source = ? AND source_id = ?"

	if err := s.GetMaster().QueryRow(sql, source, sourceId).Scan(feed.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_feed_store.get_by_source", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": source, "source_id": sourceId},
		)
	}
	return feed, nil
}

func (s *SqlFeedStore) GetBySourcePaginated(source string, offset, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Feed{}) + " FROM feed WHERE source = ? LIMIT ?, ?"

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, source, offset, limit); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_feed_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": source},
		)
	}
	return
}

func (s *SqlFeedStore) GetForRefreshPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Feed{}) + ` FROM feed
		WHERE refresh_enabled = 1 AND
			  last_refresh_comment <> 'PENDING' AND
			  next_refresh_at < ? AND
			  id > ? 
		ORDER BY id
		LIMIT ?`

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, model.Now(), lastId, limit); err != nil {
		appE = model.NewAppError("store.sqlstore.sql_feed_store.get_all_to_be_refreshed", err.Error(), http.StatusInternalServerError, nil)
	}
	return
}

func (s *SqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	if _, err := s.UpdateChanges("feed", old, new, "id = ?", new.Id); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_feed_store.update", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"id": new.Id},
		)
	}
	return nil
}
