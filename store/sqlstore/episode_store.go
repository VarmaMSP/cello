package sqlstore

import (
	"net/http"
	"strings"

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
	sql := "SELECT " + Cols(episode) + " FROM episode WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, episodeId).Scan(episode.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": episodeId},
		)
	}
	return episode, nil
}

func (s *SqlEpisodeStore) GetByIds(episodeIds []int64) (res []*model.Episode, appE *model.AppError) {
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

func (s *SqlEpisodeStore) GetByPodcast(podcastId int64) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + " FROM episode WHERE podcast_id = ?"

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, podcastId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": podcastId},
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPodcastPaginated(podcastId int64, order string, offset, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE podcast_id = ?
		ORDER BY pub_date`

	if order == "pub_date_asc" {
		sql = sql + " ASC LIMIT ?, ?"
	} else {
		sql = sql + " DESC LIMIT ?, ?"
	}

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, podcastId, offset, limit); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": podcastId},
		)
	}
	return
}

func (s *SqlEpisodeStore) GetByPodcastIdsPaginated(podcastIds []int64, offset, limit int) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}) + ` FROM episode
		WHERE podcast_id IN (` + strings.Join(Replicate("?", len(podcastIds)), ",") + `)
		ORDER BY pub_date DESC
		LIMIT ?, ?`

	var values []interface{}
	for _, podcastId := range podcastIds {
		values = append(values, podcastId)
	}
	values = append(values, offset, limit)

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

func (s *SqlEpisodeStore) GetByPlaylist(playlistId int64) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}, "episode") + ` FROM episode
		INNER JOIN playlist_member ON playlist_member.episode_id = episode.id
		WHERE playlist_member.playlist_id = ?
		ORDER BY playlist_member.updated_at DESC`

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, playlistId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_episode_store.get_by_playlist", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return
}

func (s *SqlEpisodeStore) Block(episodeId int64) *model.AppError {
	return nil
}
