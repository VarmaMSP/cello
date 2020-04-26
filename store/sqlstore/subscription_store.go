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
			ON DUPLICATE KEY UPDATE active = 1, updated_at = %d`,
		cols(subscription), vals(subscription), datetime.Unix(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_subscription_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlSubscriptionStore) GetByUser(userId int64) (res []*model.Subscription, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM subscription WHERE user_id = %d`,
		cols(&model.Subscription{}), userId,
	)
	copyTo := func() []interface{} {
		tmp := &model.Subscription{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sqlstore.sql_subscription_store.get_by_user", err.Error(), nil)
	}
	return
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
