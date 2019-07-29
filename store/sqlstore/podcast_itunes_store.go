package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastItunesStore struct {
	SqlStore
}

func NewSqlPodcastItunesStore(store SqlStore) *SqlPodcastItunesStore {
	return &SqlPodcastItunesStore{store}
}

func (s *SqlPodcastItunesStore) SaveAll(podcastItunesMeta []*model.PodcastItunes) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		models := make([]DbModel, len(podcastItunesMeta))
		for i := range models {
			models[i] = podcastItunesMeta[i]
		}

		res, err := s.Insert(models, "podcast_itunes")
		if err != nil {
			r.Err = model.NewAppError(
				"store.sqlstore.sql_podcast_itunes_store.save_all",
				err.Error(),
				http.StatusInternalServerError,
				nil,
			)
		}
		r.Data = res
	})
}

func (s *SqlPodcastItunesStore) GetItunesIdsAfter(afterId string, limit int) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		sql := `SELECT itunes_id FROM podcast_itunes 
			WHERE itunes_id > ? 
			ORDER BY itunes_id ASC 
			LIMIT ?`

		rows, err := s.GetMaster().Query(sql, afterId, limit)
		if err != nil {
			r.Err = model.NewAppError(
				"store.sqlstore.sql_podcast_itunes_store.get_itunes_ids_After",
				err.Error(),
				http.StatusInternalServerError,
				nil,
			)
			return
		}
		defer rows.Close()

		var ids []string
		for rows.Next() {
			var tmp string
			err := rows.Scan(&tmp)
			if err != nil {
				r.Err = model.NewAppError(
					"store.sqlstore.sql_podcast_itunes_store.get_itunes_ids_After",
					err.Error(),
					http.StatusInternalServerError,
					nil,
				)
				return
			}
			ids = append(ids, tmp)
		}
		if err := rows.Err(); err != nil {
			r.Err = model.NewAppError(
				"store.sqlstore.sql_podcast_itunes_store.get_itunes_ids_After",
				err.Error(),
				http.StatusInternalServerError,
				nil,
			)
			return
		}
		r.Data = ids
	})
}
