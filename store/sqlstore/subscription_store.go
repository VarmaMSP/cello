package sqlstore

import (
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

	if _, err := s.InsertOrUpdate("podcast_subscription", subscription, "active = 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save_subscription", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": subscription.UserId, "podcast_id": subscription.PodcastId},
		)
	}
	return nil
}

func (s *SqlSubscriptionStore) Delete(userId, podcastId int64) *model.AppError {
	sql := "UPDATE podcast_subscription SET active = 0 AND updated_at = ? WHERE podcast_id = ? AND subscribed_by = ?"

	if _, err := s.GetMaster().Exec(sql, model.Now(), podcastId, userId); err != nil {
		return model.NewAppError(
			"sqlstore.sql_podcast_store.delete_subscription", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"userId": userId, "podcastId": podcastId},
		)
	}
	return nil
}
