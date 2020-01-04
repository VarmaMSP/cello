package sqlstore

import (
	"fmt"
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
			map[string]interface{}{"title": podcast.Title},
		)
	}
	return nil
}

func (s *SqlPodcastStore) Get(podcastId int64) (*model.Podcast, *model.AppError) {
	podcast := &model.Podcast{}
	sql := fmt.Sprintf(
		"SELECT %s FROM podcast WHERE id = %d",
		joinStrings(podcast.DbColumns(), ","), podcastId,
	)

	if err := s.GetMaster().QueryRow(sql).Scan(podcast.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_podcast_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"id": podcastId},
		)
	}
	return podcast, nil
}

func (s *SqlPodcastStore) GetByIds(podcastIds []int64) (res []*model.Podcast, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast WHERE id IN (%s)`,
		joinStrings((&model.Podcast{}).DbColumns(), ","), joinInt64s(podcastIds, ","),
	)

	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_podcast_store.get_by_ids", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_ids": podcastIds},
		)
	}
	return
}

func (s *SqlPodcastStore) GetSubscriptions(userId int64) (res []*model.Podcast, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Podcast{}, "podcast") + ` FROM podcast
		INNER JOIN subscription ON subscription.podcast_id = podcast.id
		WHERE subscription.active = 1 AND subscription.user_id = ?
		ORDER BY subscription.updated_at DESC`

	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_podcast_store.get_subscriptions", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId},
		)
	}
	return
}

func (s *SqlPodcastStore) Update(old, new *model.Podcast) *model.AppError {
	if _, err := s.Update_("podcast", old, new, fmt.Sprintf("id = %d", new.Id)); err != nil {
		return model.NewAppError(
			"sqlstore.sql_curation_store.update", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": new.Id},
		)
	}
	return nil
}
