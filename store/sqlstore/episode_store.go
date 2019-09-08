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

	_, err := s.Insert("episode", []DbModel{episode})
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_episode_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"podcast_id": episode.PodcastId, "title": episode.Title},
		)
	}
	return nil
}

func (s *SqlEpisodeStore) GetInfo(id string) (*model.EpisodeInfo, *model.AppError) {
	info := &model.EpisodeInfo{}
	sql := "SELECT " + strings.Join(info.DbColumns(), ",") + " FROM episode WHERE id = ?"

	err := s.GetMaster().QueryRow(sql, id).Scan(info.FieldAddrs()...)
	if err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_episode_store.get_info",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return info, nil
}

func (s *SqlEpisodeStore) GetAllByPodcast(podcastId string, limit, offset int) ([]*model.EpisodeInfo, *model.AppError) {
	m := &model.EpisodeInfo{}
	sql := "SELECT " + strings.Join(m.DbColumns(), ",") + ` FROM episode
		WHERE podcast_id = ?
		ORDER BY pub_date DESC`

	appErrorC := model.NewAppErrorC(
		"store.sqlstore.sql_episode_store.get_all_by_podcast",
		http.StatusInternalServerError,
		map[string]string{"podcast_id": podcastId},
	)

	rows, err := s.GetMaster().Query(sql, podcastId)
	if err != nil {
		return nil, appErrorC((err.Error()))
	}
	defer rows.Close()

	var res []*model.EpisodeInfo
	for rows.Next() {
		tmp := &model.EpisodeInfo{}
		if err := rows.Scan(tmp.FieldAddrs()...); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}

func (s *SqlEpisodeStore) GetAllGuidsByPodcast(podcastId string) ([]string, *model.AppError) {
	sql := `SELECT guid FROM episode WHERE podcast_id = ?`

	appErrorC := model.NewAppErrorC(
		"store.sqlstore.sql_episode_store.get_all_guids_by_podcast",
		http.StatusInternalServerError,
		map[string]string{"podcast_id": podcastId},
	)

	rows, err := s.GetMaster().Query(sql, podcastId)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}

func (s *SqlEpisodeStore) Block(podcastId, episodeGuid string) *model.AppError {
	sql := `UPDATE episode SET block = 1 WHERE podcast_id = ? AND guid = ?`

	_, err := s.GetMaster().Exec(sql, podcastId, episodeGuid)
	if err != nil {
		return model.NewAppError(
			"store.sql_store.sql_episode_store.block",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"podcast_id": podcastId, "episode_guid": episodeGuid},
		)
	}
	return nil
}
