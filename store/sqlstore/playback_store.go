package sqlstore

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPlaybackStore struct {
	SqlStore
}

func NewSqlPlaybackStore(store SqlStore) store.PlaybackStore {
	return &SqlPlaybackStore{store}
}

func (s *SqlPlaybackStore) Save(playback *model.Playback) *model.AppError {
	playback.PreSave()

	if _, err := s.InsertOrUpdate("playback", playback, "play_count = play_count + 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playback_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": playback.EpisodeId, "user_id": playback.UserId},
		)
	}
	return nil
}

func (s *SqlPlaybackStore) Upsert(playback *model.Playback) *model.AppError {
	playback.PreSave()

	sql := fmt.Sprintf(
		`INSERT INTO playback (%s) VALUES (%s)
			ON DUPLICATE KEY UPDATE
				play_count = play_count + 1,
				last_played_at = '%s',
				updated_at = %d`,
		joinStrings(playback.DbColumns(), ","),
		joinValues(playback.FieldAddrs(), ","),
		model.NowDateTime(),
		model.Now(),
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playback_store.upsert", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": playback.EpisodeId, "user_id": playback.UserId},
		)
	}
	return nil
}

func (s *SqlPlaybackStore) GetByUserPaginated(userId int64, offset, limit int) (res []*model.Playback, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playback WHERE user_id = %d ORDER by last_played_at DESC LIMIT %d, %d`,
		joinStrings((&model.Playback{}).DbColumns(), ","), userId, offset, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playback_store.get_by_user_paginated", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId},
		)
	}
	return
}

func (s *SqlPlaybackStore) GetByUserByEpisodes(userId int64, episodeIds []int64) (res []*model.Playback, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playback WHERE episode_id IN (%s) AND user_id = %d`,
		joinStrings((&model.Playback{}).DbColumns(), ","), joinInt64s(episodeIds, ","), userId,
	)

	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playback_store.get_by_user_by_episodes", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlPlaybackStore) Update(playback *model.Playback) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE playback SET
			current_progress = %f,
			updated_at = %d
		WHERE user_id = %d AND episode_id = %d`,
		playback.CurrentProgress, model.Now(), playback.UserId, playback.EpisodeId,
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playback_store.update", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": playback.UserId, "episode_id": playback.EpisodeId},
		)
	}
	return nil
}
