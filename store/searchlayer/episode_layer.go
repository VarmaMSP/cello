package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type searchEpisodeStore struct {
	store.EpisodeStore
	search searchengine.Broker
}

func (s *searchEpisodeStore) Save(episode *model.Episode) *model.AppError {
	if err := s.EpisodeStore.Save(episode); err != nil {
		return err
	}
	if err := s.search.Index(searchengine.EPISODE_INDEX, episode.ForIndexing()); err != nil {
		return model.New500Error("search_layer.search_episode_store.save", err.Error(), nil)
	}
	return nil
}

// func (s *searchEpisodeStore) SyncIndex() *model.AppError {
// 	if err := s.es.DeleteIndex(EPISODE_INDEX); err != nil {
// 		return err
// 	}
// 	if err := s.es.CreateIndex(EPISODE_INDEX, EPISODE_INDEX_MAPPING); err != nil {
// 		return err
// 	}

// 	lastId, limit := int64(0), 10000
// 	for {
// 		episodes, err := s.EpisodeStore.GetAllPaginated(lastId, limit)
// 		if err != nil {
// 			return err
// 		}

// 		m := make([]model.EsModel, len(episodes))
// 		for i, episode := range episodes {
// 			m[i] = episode.ForIndexing()
// 		}

// 		if err := s.es.BulkIndexDoc(EPISODE_INDEX, m); err != nil {
// 			return err
// 		}

// 		if len(episodes) < limit {
// 			break
// 		}
// 		lastId = episodes[len(episodes)-1].Id
// 	}

// 	return nil
// }
