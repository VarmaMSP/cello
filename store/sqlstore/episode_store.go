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

func (s *SqlEpisodeStore) GetAllPublishedBetween(from, to *time.Time, podcastIds []string) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE podcast_id IN (` + strings.Join(Replicate("?", len(podcastIds)), ",") + `) AND
			  pub_date BETWEEN ? AND ?
		ORDER BY pub_date DESC`

	var values []interface{}
	for _, podcastId := range podcastIds {
		values = append(values, podcastId)
	}
	values = append(values, from.Format(model.MYSQL_DATETIME), to.Format(model.MYSQL_DATETIME))

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

func (s *SqlEpisodeStore) SetPlaybackCurrentTime(episodeId, playedBy string, currentTime int) *model.AppError {
	sql := `UPDATE episode_playback current_time = ?, updated_at = ? WHERE episode_id = ? AND played_by = ?`

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
