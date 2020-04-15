package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlSubscriptionStore struct {
	sqlStore
}

func newSqlSubscriptionStore(store sqlStore) store_.SubscriptionStore {
	return &sqlSubscriptionStore{store}
}

func (s *sqlSubscriptionStore) Save(subscription *model.Subscription) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlSubscriptionStore) Delete(userId int64, podcastId int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}
