package sqlstore

import (
	"net/http"
	"strconv"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlEpisodeStore struct {
	SqlStore
}

func NewSqlEpisodeStore(store SqlStore) store.EpisodeStore {
	return &SqlEpisodeStore{store}
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
				"store.sqlstore.sql_episode_store.save_all",
				err.Error(),
				http.StatusInternalServerError,
				map[string]string{
					"podcast_id": strconv.FormatInt(episodes[0].PodcastId, 10),
				},
			)
			return
		}
		r.Data = res
	})
}
