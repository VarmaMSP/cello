package sqlstore

import (
	"net/http"
	"strings"

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

	if _, err := s.Insert("podcast", []DbModel{podcast}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"title": podcast.Title, "feed_url": podcast.FeedUrl},
		)
	}
	return nil
}

func (s *SqlPodcastStore) GetInfo(podcastId string) (*model.PodcastInfo, *model.AppError) {
	info := &model.PodcastInfo{}
	sql := "SELECT " + strings.Join(info.DbColumns(), ",") + " FROM podcast WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, podcastId).Scan(info.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_podcast_store.get_info", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": podcastId},
		)
	}
	return info, nil
}

func (s *SqlPodcastStore) GetAllToBeRefreshed(createdAfter int64, limit int) (res []*model.PodcastFeedDetails, appE *model.AppError) {
	sql := "SELECT " + strings.Join((&model.PodcastFeedDetails{}).DbColumns(), ",") + ` FROM podcast
		WHERE refresh_enabled = 1 AND
		      last_refresh_status <> 'PENDING' AND 
			  next_refresh_at < ? AND
			  created_at > ?
		ORDER BY created_at LIMIT ?`

	copyTo := func() []interface{} {
		tmp := &model.PodcastFeedDetails{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, model.Now(), createdAfter, limit); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_podcast_store.get_all_to_be_refreshed", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return
}

func (s *SqlPodcastStore) UpdateFeedDetails(old, new *model.PodcastFeedDetails) *model.AppError {

	if _, err := s.UpdateChanges("podcast", old, new, "id = ?", new.Id); err != nil {
		return model.NewAppError(
			"sqlstore.sql_podcast_store.update_feed_details", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return nil
}
