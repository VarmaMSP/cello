package sqlstore

import (
	"fmt"

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
	res := &model.Podcast{}
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast WHERE id = %d`,
		cols(res), podcastId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_podcast_store.get", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlPodcastStore) GetAllPaginated(lastId int64, limit int) (res []*model.Podcast, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast WHERE id > %d ORDER BY id LIMIT %d`,
		cols(&model.Podcast{}), lastId, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_podcast_store.get_all_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlPodcastStore) GetByIds(podcastIds []int64) (res []*model.Podcast, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast WHERE id IN (%s)`,
		cols(&model.Podcast{}), joinInt64s(podcastIds),
	)
	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sqlstore.sql_podcast_store.get_by_ids", err.Error(), nil)
	}
	return
}

func (s *sqlPodcastStore) GetSubscriptions(userId int64) (res []*model.Podcast, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM podcast
			INNER JOIN subscription ON subscription.podcast_id = podcast.id
			WHERE subscription.active = 1 AND subscription.user_id = %d
			ORDER BY subscription.updated_at DESC`,
		cols(&model.Podcast{}, "podcast"),
		userId,
	)
	copyTo := func() []interface{} {
		tmp := &model.Podcast{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sqlstore.sql_podcast_store.get_subscriptions", err.Error(), nil)
	}
	return
}

func (s *sqlPodcastStore) GetTypeaheadSuggestions(query string) ([]*model.SearchSuggestion, *model.AppError) {
	panic("method not implemented by search layer")
}

func (s *sqlPodcastStore) Search(query string, offset, limit int) ([]*model.Podcast, *model.AppError) {
	panic("method not implemented by search layer")
}

func (s *sqlPodcastStore) Update(old *model.Podcast, new *model.Podcast) *model.AppError {
	if _, err := s.Patch("podcast", old, new); err != nil {
		return model.New500Error("sqlstore.sql_podcast_store.update", err.Error(), nil)
	}
	return nil
}
