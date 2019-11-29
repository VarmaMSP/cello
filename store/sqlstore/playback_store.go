package sqlstore

import (
	"net/http"
	"strings"

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

	if _, err := s.InsertOrUpdate("playback", playback, "total_plays = total_plays + 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.save_playback", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": playback.EpisodeId, "user_id": playback.UserId},
		)
	}
	return nil
}

func (s *SqlPlaybackStore) GetByUserPaginated(userId int64, offset, limit int) (res []*model.Playback, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Playback{}) + ` FROM playback
		WHERE user_id = ? ORDER by last_played_at DESC LIMIT ?, ?`

	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId, offset, limit); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_playbacks_by_user", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId},
		)
	}
	return
}

func (s *SqlPlaybackStore) GetByUserByEpisodes(userId int64, episodeIds []int64) (res []*model.Playback, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Playback{}) + ` FROM playback
		WHERE episode_id IN (` + strings.Join(Replicate("?", len(episodeIds)), ",") + `) AND user_id = ?`

	var values []interface{}
	for _, episodeId := range episodeIds {
		values = append(values, episodeId)
	}
	values = append(values, userId)

	copyTo := func() []interface{} {
		tmp := &model.Playback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, values...); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_playbacks", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlPlaybackStore) Update(old, new *model.Playback) *model.AppError {
	return nil
}
