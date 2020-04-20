package sqlstore

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlFeedStore struct {
	sqldb.Broker
}

func (s *sqlFeedStore) Save(feed *model.Feed) *model.AppError {
	feed.PreSave()

	res, err := s.Insert_("feed", feed)
	if err != nil {
		return model.New500Error("sql_store.sql_feed_store.save", err.Error(), nil)
	}
	feed.Id, _ = res.LastInsertId()
	return nil
}

func (s *sqlFeedStore) Get(feedId int64) (*model.Feed, *model.AppError) {
	res := &model.Feed{}
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE id = %d`,
		cols(res), feedId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlFeedStore) GetAllPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE id > %d ORDER BY id LIMIT %d`,
		cols(&model.Feed{}), lastId, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_feed_store.get_all_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlFeedStore) GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError) {
	res := &model.Feed{}
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE source = '%s' AND source_id = '%s'`,
		cols(res), source, sourceId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_by_source_id", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlFeedStore) GetBySourcePaginated(source string, offset, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE source = '%s' LIMIT %d, %d`,
		cols(&model.Feed{}), source, offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_feed_store.get_by_source_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlFeedStore) GetForRefreshPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed
		WHERE refresh_enabled = 1 AND
			  last_refresh_comment <> 'PENDING' AND
			  next_refresh_at < %d AND
			  id > %d 
		ORDER BY id
		LIMIT %d`,
		cols(&model.Feed{}), model.Now(), lastId, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_feed_store.get_for_refresh_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlFeedStore) GetFailedToImportPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM feed WHERE last_refresh_comment <> '' AND id > %d ORDER BY id LIMIT %d`,
		cols(&model.Feed{}), lastId, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Feed{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_feed_store.get_failed_to_import_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	return nil
}
