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

	_, err := s.Insert([]DbModel{podcast}, "podcast")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"title": podcast.Title, "feed_url": podcast.FeedUrl},
		)
	}
	return nil
}

func (s *SqlPodcastStore) GetInfo(podcastId string) (*model.PodcastInfo, *model.AppError) {
	info := &model.PodcastInfo{}
	sql := "SELECT " + strings.Join(info.DbColumns(), ",") + " FROM podcast WHERE id = ?"

	err := s.GetMaster().QueryRow(sql, podcastId).Scan(info.FieldAddrs()...)
	if err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_podcast_store.get_info",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"id": podcastId},
		)
	}
	return info, nil
}

func (s *SqlPodcastStore) GetAllToBeRefreshed(createdAfter int64, limit int) ([]*model.PodcastFeedDetails, *model.AppError) {
	m := &model.PodcastFeedDetails{}
	sql := "SELECT " + strings.Join(m.DbColumns(), ",") + ` FROM podcast
		WHERE refresh_enabled = 1 AND
		      last_refresh_status <> 'PENDING' AND 
			  next_refresh_at < ? AND
			  created_at > ?
		ORDER BY created_at LIMIT ?`

	appErrorC := model.NewAppErrorC(
		"sqlstore.sql_podcast_store.get_all_to_be_refreshed",
		http.StatusInternalServerError,
		nil,
	)

	rows, err := s.GetMaster().Query(sql, model.Now(), createdAfter, limit)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []*model.PodcastFeedDetails
	for rows.Next() {
		tmp := &model.PodcastFeedDetails{}
		if err := rows.Scan(tmp.FieldAddrs()...); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}

	return res, nil
}

func (s *SqlPodcastStore) UpdateFeedDetails(old, new *model.PodcastFeedDetails) *model.AppError {
	sql, values := UpdateQuery("podcast", old, new)
	if len(values) == 0 {
		return nil
	}
	sql = sql + " WHERE id = ?"
	values = append(values, new.Id)

	_, err := s.GetMaster().Exec(sql, values...)
	if err != nil {
		return model.NewAppError(
			"sqlstore.sql_podcast_store.update_feed_details",
			err.Error(),
			http.StatusInternalServerError,
			nil,
		)
	}
	return nil
}
