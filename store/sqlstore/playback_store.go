package sqlstore

import (
	"fmt"

	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/util/datetime"
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
		cols(playback), vals(playback), datetime.Now(), datetime.Unix(),
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
		cols(playback), vals(playback), datetime.Now(), datetime.Unix(),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.playback_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaybackStore) GetByUserPaginated(userId int64, offset int, limit int) ([]*model.Playback, *model.AppError) {
	query := sqlf.Select("playback.*").
		From("playback").
		Where("user_id = ?", userId).
		OrderBy("last_played_at DESC").
		Offset(offset).
		Limit(limit)

	var playbacks []*model.Playback
	if err := s.Query_(&playbacks, query); err != nil {
		return nil, model.New500Error("sql_store.sql_playback_store.get_by_user_paginated", err.Error(), nil)
	}
	return playbacks, nil
}

func (s *sqlPlaybackStore) GetByUserByEpisodes(userId int64, episodeIds []int64) (res []*model.Playback, appE *model.AppError) {
	query := sqlf.Select("playback.*").
		From("playback").
		Where("user_id = ? AND episode_id IN (?)", userId, episodeIds)

	var playbacks []*model.Playback
	if err := s.Query_(&playbacks, query, sqldb.ExpandVars); err != nil {
		return nil, model.New500Error("sql_store.sql_playback_store.get_by_user_by_episodes", err.Error(), nil)
	}
	return playbacks, nil
}

func (s *sqlPlaybackStore) Update(playback *model.Playback) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE playback
			SET current_progress = %f, updated_at = %d
			WHERE user_id = %d AND episode_id = %d`,
		playback.CurrentProgress, datetime.Unix(), playback.UserId, playback.EpisodeId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playback_store.update", err.Error(), nil)
	}
	return nil
}
