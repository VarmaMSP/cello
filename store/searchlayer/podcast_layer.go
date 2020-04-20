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
