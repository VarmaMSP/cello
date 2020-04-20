package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type searchPodcastStore struct {
	store.PodcastStore
	search searchengine.Broker
}

func (s *searchPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	if err := s.PodcastStore.Save(podcast); err != nil {
		return err
	}
	if err := s.search.Index(searchengine.PODCAST_INDEX, podcast.ForIndexing()); err != nil {
		return model.New500Error("search_layer.search_podcast_store.save", err.Error(), nil)
	}
	return nil
}

// func (s *searchPodcastStore) SyncIndex() *model.AppError {
// 	if err := s.es.DeleteIndex(PODCAST_INDEX); err != nil {
// 		return err
// 	}
// 	if err := s.es.CreateIndex(PODCAST_INDEX, PODCAST_INDEX_MAPPING); err != nil {
// 		return err
// 	}

// 	lastId, limit := int64(0), 10000
// 	for {
// 		podcasts, err := s.PodcastStore.GetAllPaginated(lastId, limit)
// 		if err != nil {
// 			return err
// 		}

// 		m := make([]model.EsModel, len(podcasts))
// 		for i, podcast := range podcasts {
// 			m[i] = podcast.ForIndexing()
// 		}

// 		if err := s.es.BulkIndexDoc(PODCAST_INDEX, m); err != nil {
// 			return err
// 		}

// 		if len(podcasts) < limit {
// 			break
// 		}
// 		lastId = podcasts[len(podcasts)-1].Id
// 	}

// 	return nil
// }
