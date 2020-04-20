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
