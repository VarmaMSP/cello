package sqlstore

import (
	"fmt"
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
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE id = %d`,
		joinStrings(feed.DbColumns(), ","), feedId,
	)

	if err := s.GetMaster().QueryRow(sql).Scan(feed.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_feed_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"id": feed},
		)
	}
	return feed, nil
}

func (s *SqlFeedStore) GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError) {
	feed := &model.Feed{}
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE source = %s AND source_id = %s`,
		joinStrings(feed.DbColumns(), ","), source, sourceId,
	)

	if err := s.GetMaster().QueryRow(sql).Scan(feed.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_feed_store.get_by_source_id", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": source, "source_id": sourceId},
		)
	}
	return feed, nil
}

func (s *SqlFeedStore) GetBySourcePaginated(source string, offset, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE source = '%s' LIMIT %d, %d`,
		joinStrings((&model.Feed{}).DbColumns(), ","), source, offset, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_feed_store.get_by_source_paginated", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": source},
		)
	}
	return
}

func (s *SqlFeedStore) GetForRefreshPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed
		WHERE refresh_enabled = 1 AND
			  last_refresh_comment <> 'PENDING' AND
			  next_refresh_at < %d AND
			  id > %d 
		ORDER BY id
		LIMIT %d`,
		joinStrings((&model.Feed{}).DbColumns(), ","), model.Now(), lastId, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError("store.sqlstore.sql_feed_store.get_for_refresh_paginated", err.Error(), http.StatusInternalServerError, nil)
	}
	return
}

func (s *SqlFeedStore) GetFailedToImportPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE last_refresh_comment = '' AND id > %d ORDER BY id LIMIT %d`,
		joinStrings((&model.Feed{}).DbColumns(), ","), lastId, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError("store.sqlstore.sql_feed_store.get_failed_to_import_paginated", err.Error(), http.StatusInternalServerError, nil)
	}
	return
}

func (s *SqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	if _, err := s.Update_("feed", old, new, fmt.Sprintf("id = %d", new.Id)); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_feed_store.update", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"id": new.Id},
		)
	}
	return nil
}
