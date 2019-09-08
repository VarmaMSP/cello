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

	if _, err := s.Insert("episode", []DbModel{episode}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"podcast_id": episode.PodcastId, "title": episode.Title},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) GetInfo(id string) (*model.EpisodeInfo, *model.AppError) {
	info := &model.EpisodeInfo{}
	sql := "SELECT " + strings.Join(info.DbColumns(), ",") + " FROM episode WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(info.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get_info", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return info, nil
}

func (s *SqlEpisodeStore) GetAllByPodcast(podcastId string, limit, offset int) ([]*model.EpisodeInfo, *model.AppError) {
	sql := "SELECT " + strings.Join((&model.EpisodeInfo{}).DbColumns(), ",") + ` FROM episode
		WHERE podcast_id = ?
		ORDER BY pub_date DESC`

	var res []*model.EpisodeInfo
	newItemFields := func() []interface{} {
		tmp := &model.EpisodeInfo{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.QueryRows(newItemFields, sql, podcastId); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]string{"podcast_id": podcastId},
		)
	}
	return res, nil
}

func (s *SqlEpisodeStore) GetAllGuidsByPodcast(podcastId string) ([]string, *model.AppError) {
	sql := `SELECT guid FROM episode WHERE podcast_id = ?`

	var res []string
	newItemFields := func() []interface{} {
		tmp := ""
		res = append(res, tmp)
		return []interface{}{&tmp}
	}

	if err := s.QueryRows(newItemFields, sql); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get_all_guids_by_podcast", err.Error(), http.StatusInternalServerError,
			map[string]string{"podcast_id": podcastId},
		)
	}
	return res, nil
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
