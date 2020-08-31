package sqlstore

import (
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/util/datetime"
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
	query := sqlf.Select("feed.*").
		From("feed").
		Where("id = ?", feedId)

	var feed model.Feed
	if err := s.QueryRow_(&feed, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get", err.Error(), nil)
	}
	return &feed, nil
}

func (s *sqlFeedStore) GetAllPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError) {
	query := sqlf.Select("feed.*").
		From("feed").
		Where("id > ?", lastId).
		OrderBy("id").
		Limit(limit)

	var feeds []*model.Feed
	if err := s.Query_(&feeds, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_all_paginated", err.Error(), nil)
	}
	return feeds, nil
}

func (s *sqlFeedStore) GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError) {
	query := sqlf.Select("feed.*").
		From("feed").
		Where("source = ? AND source_id = ?", source, sourceId)

	var feed model.Feed
	if err := s.QueryRow_(&feed, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_by_source_id", err.Error(), nil)
	}
	return &feed, nil
}

func (s *sqlFeedStore) GetBySourcePaginated(source string, offset, limit int) (res []*model.Feed, appE *model.AppError) {
	query := sqlf.Select("feed.*").
		From("feed").
		Where("source = ?", source).
		Offset(offset).
		Limit(limit)

	var feeds []*model.Feed
	if err := s.Query_(&feeds, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_by_source_paginated", err.Error(), nil)
	}
	return feeds, nil
}

func (s *sqlFeedStore) GetForRefreshPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError) {
	query := sqlf.Select("feed.*").
		From("feed").
		Where(`
			refresh_enabled = 1 AND
			last_refresh_comment <> 'PENDING' AND
			next_refresh_at < ? AND
			id > ?
		`, datetime.Unix(), lastId, limit).
		OrderBy("id").
		Limit(limit)

	var feeds []*model.Feed
	if err := s.Query_(&feeds, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_for_refresh_paginated", err.Error(), nil)
	}
	return feeds, nil
}

func (s *sqlFeedStore) GetFailedToImportPaginated(lastId int64, limit int) (res []*model.Feed, appE *model.AppError) {
	query := sqlf.Select("feed.*").
		From("feed").
		Where("last_refresh_comment <> '' AND id > ?", lastId).
		OrderBy("id").
		Limit(limit)

	var feeds []*model.Feed
	if err := s.Query_(&feeds, query); err != nil {
		return nil, model.New500Error("sql_store.sql_feed_store.get_failed_to_import_paginated", err.Error(), nil)
	}
	return feeds, nil
}

func (s *sqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	if _, err := s.Patch("feed", old, new); err != nil {
		return model.New500Error("sql_store.sql_feed_store.update", err.Error(), nil)
	}
	return nil
}
