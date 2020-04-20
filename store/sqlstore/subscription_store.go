package sqlstore

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/util/datetime"
)

type sqlSubscriptionStore struct {
	sqldb.Broker
}

func (s *sqlSubscriptionStore) Save(subscription *model.Subscription) *model.AppError {
	subscription.PreSave()

	sql := fmt.Sprintf(
		`INSERT INTO subscription (%s) VALUES (%s)
			ON DUPLICATE UPDATE active = 1, updated_at = %d`,
		cols(subscription), vals(subscription), datetime.Unix(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_subscription_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlSubscriptionStore) Delete(userId int64, podcastId int64) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE subscription SET active = 0 AND updated_at = %d
			WHERE podcast_id = %d AND user_id = %d`,
		datetime.Unix(), podcastId, userId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_subscription_store.delete", err.Error(), nil)
	}
	return nil
}
