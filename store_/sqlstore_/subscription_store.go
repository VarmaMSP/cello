package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlSubscriptionStore struct {
	sqldb.Broker
}

func (s *sqlSubscriptionStore) Save(subscription *model.Subscription) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlSubscriptionStore) Delete(userId int64, podcastId int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}
