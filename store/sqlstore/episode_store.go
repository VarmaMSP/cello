package sqlstore

import (
	"net/http"

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

	_, err := s.Insert([]DbModel{episode}, "episode")
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
