package sqlstore

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlEpisodeStore struct {
	sqldb.Broker
}

func (s *sqlEpisodeStore) Save(episode *model.Episode) *model.AppError {
	episode.PreSave()

	res, err := s.Insert_("episode", episode)
	if err != nil {
		return model.New500Error("sql_store.sql_episode_store.save", err.Error(), nil)
	}
	episode.Id, _ = res.LastInsertId()
	return nil
}

func (s *sqlEpisodeStore) Get(episodeId int64) (*model.Episode, *model.AppError) {
	res := &model.Episode{}
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE id = %d`,
		cols(res), episodeId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlEpisodeStore) GetAllPaginated(lastId int64, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE id > %d ORDER BY id LIMIT %d`,
		cols(&model.Episode{}), lastId, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_all_panigated", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) GetByIds(episodeIds []int64) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE id IN (%s)`,
		cols(&model.Episode{}), joinInt64s(episodeIds),
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_by_ids", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) GetByPodcast(podcastId int64) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id = %d`,
		cols(&model.Episode{}), podcastId,
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_by_podcast", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) GetByPodcastPaginated(podcastId int64, order string, offset int, limit int) (res []*model.Episode, appE *model.AppError) {
	sqlOrder := "DESC"
	if order == "pub_date_asc" {
		sqlOrder = "ASC"
	}

	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id = %d ORDER by pub_date %s LIMIT %d, %d`,
		cols(&model.Episode{}), podcastId, sqlOrder, offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_by_podcast_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) GetByPodcastIdsPaginated(podcastIds []int64, offset int, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode WHERE podcast_id IN (%s) ORDER BY pub_date DESC LIMIT %d, %d`,
		cols(&model.Episode{}), joinInt64s(podcastIds), offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_by_podcast_ids_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) GetByPlaylistPaginated(playlistId int64, offset int, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM episode INNER JOIN playlist_member ON playlist_member.episode_id = episode.id
			WHERE playlist_member.playlist_id = %ds
			ORDER BY playlist_member.update_at DESC
			LIMIT %d, %d`,
		cols(&model.Episode{}), playlistId, offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_episode_store.get_by_playlist_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlEpisodeStore) Block(episodeIds []int64) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE episode SET block = 1 WHERE id IN (%s)`,
		joinInt64s(episodeIds),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_episode_store.block", err.Error(), nil)
	}
	return nil
}
