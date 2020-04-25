package app

import (
	"github.com/varmamsp/cello/model"
)

func (a *App) GetTypeaheadSuggestions(query string) ([]*model.SearchSuggestion, *model.AppError) {
	return a.Store.Podcast().GetTypeaheadSuggestions(query)
}

func (a *App) SearchEpisodes(query, sortBy string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return a.Store.Episode().Search(query, sortBy, offset, limit)
}

func (a *App) SearchPodcasts(query string, offset, limit int) ([]*model.Podcast, *model.AppError) {
	return a.Store.Podcast().Search(query, offset, limit)
}

func (a *App) SearchEpisodesInPodcast(podcastId int64, query string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return a.Store.Episode().SearchByPodcast(podcastId, query, offset, limit)
}
