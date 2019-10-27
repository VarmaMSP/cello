package sqlstore

import (
	"net/http"
	"strings"
	"time"

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
			map[string]string{"podcast_id": episode.PodcastId, "title": episode.Title},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) SavePlayback(playback *model.EpisodePlayback) *model.AppError {
	playback.PreSave()

	if _, err := s.InsertOrUpdate("episode_playback", playback, "count_ = count_ + 1, updated_at = ?", model.Now()); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.save_playback", err.Error(), http.StatusInternalServerError,
			map[string]string{"episode_id": playback.EpisodeId, "played_by": playback.PlayedBy},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) Get(id string) (*model.Episode, *model.AppError) {
	episode := &model.Episode{}
	sql := "SELECT " + Cols(episode) + " FROM episode WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(episode.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return episode, nil
}

func (s *SqlEpisodeStore) GetAllByIds(episodeIds []string) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE id IN (` + strings.Join(Replicate("?", len(episodeIds)), ",") + `)`

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	var values []interface{}
	for _, episodeId := range episodeIds {
		values = append(values, episodeId)
	}

	if err := s.Query(copyTo, sql, values...); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_by_played", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlEpisodeStore) GetAllByPodcast(podcastId string, limit, offset int) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE podcast_id = ?
		ORDER BY pub_date DESC`

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, podcastId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]string{"podcast_id": podcastId},
		)
	}
	return
}

func (s *SqlEpisodeStore) GetAllPublishedBefore(podcastIds []string, before *time.Time, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE podcast_id IN (` + strings.Join(Replicate("?", len(podcastIds)), ",") + `) AND pub_date < ?
		ORDER BY pub_date DESC
		LIMIT ?`

	var values []interface{}
	for _, podcastId := range podcastIds {
		values = append(values, podcastId)
	}
	values = append(values, model.FormatDateTime(before), limit)

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, values...); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_published_between", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlEpisodeStore) GetAllPlaybacks(episodeIds []string, userId string) (res []*model.EpisodePlayback, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.EpisodePlayback{}, "episode_playback") + ` FROM episode_playback
		WHERE episode_id IN (` + strings.Join(Replicate("?", len(episodeIds)), ",") + `) AND played_by = ?`

	var values []interface{}
	for _, episodeId := range episodeIds {
		values = append(values, episodeId)
	}
	values = append(values, userId)

	copyTo := func() []interface{} {
		tmp := &model.EpisodePlayback{}
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

func (s *SqlEpisodeStore) GetAllPlaybacksByUser(userId string) (res []*model.EpisodePlayback, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.EpisodePlayback{}) + ` FROM episode_playback WHERE played_by = ? ORDER by updated_at DESC`

	copyTo := func() []interface{} {
		tmp := &model.EpisodePlayback{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_playbacks_by_user", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": userId},
		)
	}
	return
}

func (s *SqlEpisodeStore) SetPlaybackCurrentTime(episodeId, playedBy string, currentTime int) *model.AppError {
	sql := `UPDATE episode_playback SET current_time_ = ?, updated_at = ? WHERE episode_id = ? AND played_by = ?`

	if _, err := s.GetMaster().Exec(sql, currentTime, model.Now(), episodeId, playedBy); err != nil {
		return model.NewAppError(
			"store.sql_store.sql_episode_store.set_playback_current_time", err.Error(), http.StatusInternalServerError,
			map[string]string{"episode_id": episodeId, "played_by": playedBy},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) Block(podcastId, episodeGuid string) *model.AppError {
	sql := `UPDATE episode SET block = 1 WHERE podcast_id = ? AND guid = ?`

	if _, err := s.GetMaster().Exec(sql, podcastId, episodeGuid); err != nil {
		return model.NewAppError(
			"store.sql_store.sql_episode_store.block", err.Error(), http.StatusInternalServerError,
			map[string]string{"podcast_id": podcastId, "episode_guid": episodeGuid},
		)
	}
	return nil
}
