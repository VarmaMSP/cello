package sqlstore

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlEpisodeStore struct {
	SqlStore
}

func NewSqlEpisodeStore(store SqlStore) store.EpisodeStore {
	return &SqlEpisodeStore{store}
}

func (s *SqlEpisodeStore) Save(episode *model.Episode) *model.AppError {
	episode.PreSave()

	if _, err := s.Insert("episode", []model.DbModel{episode}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": episode.PodcastId, "title": episode.Title},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) Get(episodeId int64) (*model.Episode, *model.AppError) {
	episode := &model.Episode{}
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE id = %d`,
		joinStrings(episode.DbColumns(), ","), episodeId,
	)

	if err := s.GetMaster().QueryRow(sql).Scan(episode.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": episodeId},
		)
	}
	return episode, nil
}

func (s *SqlEpisodeStore) GetByIds(episodeIds []int64) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE id IN (%s)`,
		joinStrings((&model.Episode{}).DbColumns(), ","), joinInt64s(episodeIds, ","),
	)

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_ids", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPodcast(podcastId int64) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id = %d`,
		joinStrings((&model.Episode{}).DbColumns(), ","), podcastId,
	)

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": podcastId},
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPodcastPaginated(podcastId int64, order string, offset, limit int) (res []*model.Episode, appE *model.AppError) {
	sqlOrder := "DESC"
	if order == "pub_date_asc" {
		sqlOrder = "ASC"
	}

	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id = %d ORDER by pub_date %s LIMIT %d, %d`,
		joinStrings((&model.Episode{}).DbColumns(), ","), podcastId, sqlOrder, offset, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_podcast_paginated", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": podcastId},
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPodcastIdsPaginated(podcastIds []int64, offset, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id IN (%s) ORDER BY pub_date DESC LIMIT %d, %d`,
		joinStrings((&model.Episode{}).DbColumns(), ","), joinInt64s(podcastIds, ","), offset, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_podcast_ids_paginated", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPlaylistPaginated(playlistId int64, offset, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}, "episode") + ` FROM episode
		INNER JOIN playlist_member ON playlist_member.episode_id = episode.id
		WHERE playlist_member.playlist_id = ?
		ORDER BY playlist_member.updated_at DESC
		LIMIT ?, ?`

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, playlistId, offset, limit); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_playlist_paginated", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return
}

func (s *SqlEpisodeStore) Block(episodeId int64) *model.AppError {
	sql := fmt.Sprintf(`UPDATE episode SET block = 1 WHERE id = %d`, episodeId)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.block", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": episodeId},
		)
	}
	return nil
}
