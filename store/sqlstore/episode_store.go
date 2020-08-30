package sqlstore

import (
	"fmt"

	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlEpisodeStore struct {
	sqldb.Broker
	episodeQuery *sqlf.Stmt
}

func newSqlEpisodeStore(broker sqldb.Broker) *sqlEpisodeStore {
	return &sqlEpisodeStore{
		Broker: broker,
		episodeQuery: sqlf.NoDialect.
			Select("episode.*").
			From("episode"),
	}
}

func (s *sqlEpisodeStore) Save(episode *model.Episode) *model.AppError {
	episode.PreSave()

	res, err := s.Insert_("episode", episode)
	if err != nil {
		return model.New500Error("sql_store.sql_episode_store.save", err.Error(), nil)
	}
	episode.Id, _ = res.LastInsertId()
	return nil
}

func (s *sqlEpisodeStore) Get(episodeId int64) (*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("id = ?", episodeId)

	var episode model.Episode
	if err := s.QueryRow_(&episode, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get", err.Error(), nil)
	}

	return &episode, nil
}

func (s *sqlEpisodeStore) GetAllPaginated(lastId int64, limit int) ([]*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("id > ?", lastId).
		Limit(limit)

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_all_panigated", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) GetByIds(episodeIds []int64) ([]*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("id IN (?)", episodeIds)

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query, sqldb.ExpandVars); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_by_ids", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) GetByPodcast(podcastId int64) ([]*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("podcast_id = ?", podcastId)

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_by_podcast", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) GetByPodcastPaginated(podcastId int64, order string, offset int, limit int) ([]*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("podcast_id = ?", podcastId).
		Offset(offset).
		Limit(limit)

	if order == "pub_date_asc" {
		query = query.OrderBy("pub_date ASC")
	} else {
		query = query.OrderBy("pub_date DESC")
	}

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_by_podcast_paginated", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) GetByPodcastIdsPaginated(podcastIds []int64, offset int, limit int) ([]*model.Episode, *model.AppError) {
	query := s.episodeQuery.
		Where("podcast_id IN ?", podcastIds).
		Offset(offset).
		Limit(limit)

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_by_podcast_ids_paginated", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) GetByPlaylistPaginated(playlistId int64, offset int, limit int) (res []*model.Episode, appE *model.AppError) {
	query := s.episodeQuery.
		Join("playlist_member AS pm", "pm.episode_id = episode.id").
		Where("pm.playlist_id = ?", playlistId).
		OrderBy("pm.update_at DESC").
		Offset(offset).
		Limit(limit)

	var episodes []*model.Episode
	if err := s.Query_(&episodes, query); err != nil {
		return nil, model.New500Error("sql_store.sql_episode_store.get_by_playlist_paginated", err.Error(), nil)
	}

	return episodes, nil
}

func (s *sqlEpisodeStore) Block(episodeIds []int64) *model.AppError {
	sql := fmt.Sprintf(
		`UPDATE episode SET block = 1 WHERE id IN (%s)`,
		joinInt64s(episodeIds),
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_episode_store.block", err.Error(), nil)
	}
	return nil
}

// NOT IMPLEMENTED WITH MYSQL
func (s *sqlEpisodeStore) Search(query, sortBy string, offset, limit int) ([]*model.Episode, *model.AppError) {
	panic("episode.search method not implemented by search layer")
}

func (s *sqlEpisodeStore) SearchByPodcast(podcastId int64, query string, offset, limit int) ([]*model.Episode, *model.AppError) {
	panic("episode.search method not implemented by search layer")
}
