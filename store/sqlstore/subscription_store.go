package sqlstore

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlSubscriptionStore struct {
	SqlStore
}

func NewSqlSubscriptionStore(store SqlStore) store.SubscriptionStore {
	return &SqlSubscriptionStore{store}
}

func (s *SqlSubscriptionStore) Save(subscription *model.Subscription) *model.AppError {
	subscription.PreSave()

	if _, err := s.InsertOrUpdate("subscription", subscription, "active = 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save_subscription", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": subscription.UserId, "podcast_id": subscription.PodcastId},
		)
	}
	return nil
}

func (s *SqlSubscriptionStore) Delete(userId, podcastId int64) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE subscription SET active = 0 AND updated_at = %d WHERE podcast_id = %d AND user_id = %d`,
		model.Now(), podcastId, userId,
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"sqlstore.sql_podcast_store.delete_subscription", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"userId": userId, "podcastId": podcastId},
		)
	}
	return nil
}
