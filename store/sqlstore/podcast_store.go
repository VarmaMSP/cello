package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastStore struct {
	SqlStore
}

func NewSqlPodcastStore(store SqlStore) store.PodcastStore {
	return &SqlPodcastStore{store}
}

func (s *SqlPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	podcast.PreSave()

	if _, err := s.Insert("podcast", []model.DbModel{podcast}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"title": podcast.Title},
		)
	}
	return nil
}

func (s *SqlPodcastStore) SaveSubscription(subscription *model.PodcastSubscription) *model.AppError {
	subscription.PreSave()

	if _, err := s.InsertOrUpdate("podcast_subscription", subscription, "active = 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save_subscription", err.Error(), http.StatusInternalServerError,
			map[string]string{"subscribed_by": subscription.SubscribedBy, "podcast_id": subscription.PodcastId},
		)
	}
	return nil
}

func (s *SqlPodcastStore) Get(podcastId string) (*model.Podcast, *model.AppError) {
	podcast := &model.Podcast{}
	sql := "SELECT " + Cols(podcast) + " FROM podcast WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, podcastId).Scan(podcast.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_podcast_store.get_podcast", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": podcastId},
		)
	}
	return podcast, nil
}

func (s *SqlPodcastStore) GetAllByCuration(curationId string, offset, limit int) (res []*model.Podcast, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Podcast{}, "podcast") + ` FROM podcast
		INNER JOIN podcast_curation ON podcast_curation.podcast_id = podcast.id
		WHERE podcast_curation.curation_id = ? 
		LIMIT ?, ?`

	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, curationId, offset, limit); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_curation_store.get_podcast_by_curation", err.Error(), http.StatusInternalServerError,
			map[string]string{"curation_id": curationId},
		)
	}
	return
}

func (s *SqlPodcastStore) GetAllSubscribedBy(userId string) (res []*model.Podcast, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Podcast{}, "podcast") + ` FROM podcast
		INNER JOIN podcast_subscription ON podcast_subscription.podcast_id = podcast.id
		WHERE podcast_subscription.subscribed_by = ?
		ORDER BY podcast_subscription.updated_at DESC`

	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_curation_store.get_podcast_by_curation", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": userId},
		)
	}
	return
}

func (s *SqlPodcastStore) DeleteSubscription(userId, podcastId string) *model.AppError {
	sql := "UPDATE podcast_subscription SET active = 0 AND updated_at = ? WHERE podcast_id = ? AND subscribed_by = ?"

	if _, err := s.GetMaster().Exec(sql, model.Now(), podcastId, userId); err != nil {
		return model.NewAppError(
			"sqlstore.sql_podcast_store.delete_subscription", err.Error(), http.StatusInternalServerError,
			map[string]string{"userId": userId, "podcastId": podcastId},
		)
	}
	return nil
}
