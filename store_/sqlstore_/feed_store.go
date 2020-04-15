package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlFeedStore struct {
	sqlStore
}

func newSqlFeedStore(store sqlStore) store_.FeedStore {
	return &sqlFeedStore{store}
}

func (s *sqlFeedStore) Save(feed *model.Feed) *model.AppError {
	return nil
}

func (s *sqlFeedStore) Get(feedId int64) (*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) GetAllPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) GetBySourceId(source, sourceId string) (*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) GetBySourcePaginated(source string, offset, limit int) ([]*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) GetForRefreshPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) GetFailedToImportPaginated(lastId int64, limit int) ([]*model.Feed, *model.AppError) {
	return nil, nil
}

func (s *sqlFeedStore) Update(old, new *model.Feed) *model.AppError {
	return nil
}
