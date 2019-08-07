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
