package sqlstore

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlPlaybackStore struct {
	sqldb.Broker
}

func (s *sqlPlaybackStore) Save(playback *model.Playback) *model.AppError {
	playback.PreSave()

	sql := fmt.Sprintf(
		`INSERT INTO playback (%s) VALUES (%s)
			ON DUPLICATE KEY UPDATE
				play_count = play_count + 1,
				last_played_at = '%s',
				updated_at = %d`,
		cols(playback), vals(playback), model.NowDateTime(), model.Now(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.playback_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaybackStore) Upsert(playback *model.Playback) *model.AppError {
	playback.PreSave()

	sql := fmt.Sprintf(
		`INSERT INTO playback (%s) VALUES (%s)
			ON DUPLICATE KEY UPDATE
				play_count = play_count + 1,
				last_played_at = '%s',
				updated_at = %d`,
		cols(playback), vals(playback), model.NowDateTime(), model.Now(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.playback_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaybackStore) GetByUserPaginated(userId int64, offset int, limit int) (res []*model.Playback, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playback WHERE user_id = %d
			ORDER by last_played_at
			DESC LIMIT %d, %d`,
		cols(&model.Playback{}), userId, offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playback_store.get_by_user_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlPlaybackStore) GetByUserByEpisodes(userId int64, episodeIds []int64) (res []*model.Playback, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playback WHERE episode_id IN (%s) AND user_id = %d`,
		cols(&model.Playback{}), joinInt64s(episodeIds), userId,
	)
	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playback_store.get_by_user_by_episodes", err.Error(), nil)
	}
	return
}

func (s *sqlPlaybackStore) Update(playback *model.Playback) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE playback
			SET current_progress = %f, updated_at = %d
			WHERE user_id = %d AND episode_id = %d`,
		playback.CurrentProgress, model.Now(), playback.UserId, playback.EpisodeId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playback_store.update", err.Error(), nil)
	}
	return nil
}
