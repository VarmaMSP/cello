package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlEpisodeStore struct {
	SqlStore
}

func (s *SqlEpisodeStore) SaveAll(episodes []*model.Episode) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		models := make([]DbModel, len(episodes))
		for i := range models {
			models[i] = episodes[i]
		}

		res, err := s.Insert(models, "episode")
		if err != nil {
			r.Err = model.NewAppError(
				"EpisodeStore.SaveAll",
				"store.sqlstore.sql_episode_store.save_all",
				map[string]interface{}{
					"podcast_id": episodes[0].PodcastId,
				},
				"Cannot save episodes",
				http.StatusInternalServerError,
			)
			return
		}
		r.Data = res
	})
}

func NewSqlEpisodeStore(store SqlStore) store.EpisodeStore {
	return &SqlEpisodeStore{store}
}
