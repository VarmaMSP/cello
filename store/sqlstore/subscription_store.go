package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/leporo/sqlf"
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
			ON DUPLICATE KEY UPDATE active = 1, updated_at = %d`,
		cols(subscription), vals(subscription), datetime.Unix(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_subscription_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlSubscriptionStore) GetByUser(userId int64) ([]*model.Subscription, *model.AppError) {
	query := sqlf.
		Select("*").
		From("subscription").
		Where("user_id = ?", userId)

	var subscriptions []*model.Subscription
	if err := s.Query(&subscriptions, query); err != nil {
		return nil, model.New500Error("sqlstore.sql_subscription_store.get_by_user", err.Error(), nil)
	}
	return subscriptions, nil
}

func (s *sqlSubscriptionStore) IsUserSubscribed(userId int64, podcastId int64) (bool, *model.AppError) {
	query := sqlf.
		Select("*").
		From("subscription").
		Where("user_id = ? AND podcast_id = ?", userId, podcastId)

	var subscription model.Subscription
	if err := s.QueryRow(&subscription, query); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, model.New500Error("sqlstore.sql_subscription_store.is_user_subscribed", err.Error(), nil)
	}
	return true, nil
}

func (s *sqlSubscriptionStore) Delete(userId int64, podcastId int64) *model.AppError {
	sql := fmt.Sprintf(
		`DELETE FROM subscription WHERE podcast_id = %d AND user_id = %d`,
		podcastId, userId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_subscription_store.delete", err.Error(), nil)
	}
	return nil
}
