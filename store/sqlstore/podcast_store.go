package sqlstore

import (
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlPodcastStore struct {
	sqldb.Broker
}

func (s *sqlPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	podcast.PreSave()

	if _, err := s.Insert("podcast", podcast); err != nil {
		return model.New500Error("sql_store.sql_podcast_store.save", err.Error(), nil)
	}
	return nil
}

func (s *sqlPodcastStore) Get(podcastId int64) (*model.Podcast, *model.AppError) {
	query := sqlf.Select("*").
		From("podcast").
		Where("id = ?", podcastId)

	var podcast model.Podcast
	if err := s.QueryRow(&podcast, query); err != nil {
		return nil, model.New500Error("sql_store.sql_podcast_store.get", err.Error(), nil)
	}
	return &podcast, nil
}

func (s *sqlPodcastStore) GetAllPaginated(lastId int64, limit int) ([]*model.Podcast, *model.AppError) {
	query := sqlf.Select("*").
		From("podcast").
		Where("id > ?", lastId).
		Limit(limit)

	var podcasts []*model.Podcast
	if err := s.Query(&podcasts, query); err != nil {
		return nil, model.New500Error("sql_store.sql_podcast_store.get_all_paginated", err.Error(), nil)
	}
	return podcasts, nil
}

func (s *sqlPodcastStore) GetByIds(podcastIds []int64) ([]*model.Podcast, *model.AppError) {
	query := sqlf.Select("*").
		From("podcast").
		Where("id IN ?", podcastIds)

	var podcasts []*model.Podcast
	if err := s.Query(&podcasts, query); err != nil {
		return nil, model.New500Error("sqlstore.sql_podcast_store.get_by_ids", err.Error(), nil)
	}
	return podcasts, nil
}

func (s *sqlPodcastStore) GetSubscriptions(userId int64) ([]*model.Podcast, *model.AppError) {
	query := sqlf.Select("podcast.*").
		From("podcast").
		Join("subscription AS sub", "sub.podcast_id = podcast.id").
		Where("sub.active = 1 AND sub.user_id = ?", userId).
		OrderBy("sub.updated_at DESC")

	var podcasts []*model.Podcast
	if err := s.Query(&podcasts, query); err != nil {
		return nil, model.New500Error("sqlstore.sql_podcast_store.get_subscriptions", err.Error(), nil)
	}
	return podcasts, nil
}

func (s *sqlPodcastStore) Update(old *model.Podcast, new *model.Podcast) *model.AppError {
	if _, err := s.Patch("podcast", old, new); err != nil {
		return model.New500Error("sqlstore.sql_podcast_store.update", err.Error(), nil)
	}
	return nil
}

func (s *sqlPodcastStore) GetTypeaheadSuggestions(query string) ([]*model.SearchSuggestion, *model.AppError) {
	panic("method not implemented by search layer")
}

func (s *sqlPodcastStore) Search(query string, offset, limit int) ([]*model.Podcast, *model.AppError) {
	panic("method not implemented by search layer")
}
