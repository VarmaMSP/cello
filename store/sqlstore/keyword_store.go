package sqlstore

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
)

type SqlKeywordStore struct {
	SqlStore
}

func NewSqlKeywordStore(store SqlStore) *SqlKeywordStore {
	return &SqlKeywordStore{store}
}

func (s *SqlKeywordStore) Upsert(keyword *model.Keyword) (*model.Keyword, *model.AppError) {
	if k, err := s.GetByText(keyword.Text); err != nil {
		return nil, err
	} else if len(k) > 1 {
		return k[0], nil
	}

	keyword.PreSave()

	id, err := s.InsertWithoutPK("keyword", keyword)
	if err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_keyword_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"text": keyword.Text},
		)
	}

	keyword.Id = id
	return keyword, nil
}

func (s *SqlKeywordStore) SavePodcastKeyword(podcastKeyword *model.PodcastKeyword) (*model.PodcastKeyword, *model.AppError) {
	if _, err := s.Insert("podcast_keyword", []model.DbModel{podcastKeyword}); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_keyword_store.save_podcast_keyword", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"podcast_id": podcastKeyword.PodcastId},
		)
	}
	return podcastKeyword, nil
}

func (s *SqlKeywordStore) SaveEpisodeKeyword(episodeKeyword *model.EpisodeKeyword) (*model.EpisodeKeyword, *model.AppError) {
	if _, err := s.Insert("episode_keyword", []model.DbModel{episodeKeyword}); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_keyword_store.save_episode_keyword", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": episodeKeyword.EpisodeId},
		)
	}
	return episodeKeyword, nil
}

func (s *SqlKeywordStore) GetByText(text string) (res []*model.Keyword, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM keyword WHERE text = '%s'",
		joinStrings((&model.Keyword{}).DbColumns(), ","), text,
	)

	copyTo := func() []interface{} {
		tmp := &model.Keyword{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_keyword_store.get_by_text", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"text": text},
		)
	}
	return
}

func (s *SqlKeywordStore) GetAllPaginated(lastId int64, limit int) (res []*model.Keyword, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM keyword WHERE id > %d ORDER BY id LIMIT %d`,
		joinStrings((&model.Keyword{}).DbColumns(), ","), lastId, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Keyword{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_keyword_store.get_all_paginated", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlKeywordStore) GetDuplicates() (res []string, appE *model.AppError) {
	sql := "SELECT text FROM keyword GROUP BY text HAVING count(id) > 1"

	copyTo := func() []interface{} {
		tmp := ""
		res = append(res, tmp)
		return []interface{}{&tmp}
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_keyword_store.get_duplicates", err.Error(), http.StatusInternalServerError, nil,
		)
	}
	return
}

func (s *SqlKeywordStore) SetText(keywordId int64, text string) *model.AppError {
	sql := fmt.Sprintf("UPDATE keyword SET text = '%s' WHERE id = %d", text, keywordId)
	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_keyword_store.set_text", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"keyword_id": keywordId, "text": text},
		)
	}

	return nil
}

func (s *SqlKeywordStore) Delete(keywordId int64) *model.AppError {
	appErrorC := model.NewAppErrorC(
		"store.sqlstore.sql_keyword_store.delete", http.StatusInternalServerError, map[string]interface{}{"keyword_id": keywordId},
	)

	sql := fmt.Sprintf("DELETE FROM podcast_keyword WHERE keyword_id = %d", keywordId)
	if _, err := s.GetMaster().Exec(sql); err != nil {
		return appErrorC(err.Error())
	}

	sql = fmt.Sprintf("DELETE FROM episode_keyword WHERE keyword_id = %d", keywordId)
	if _, err := s.GetMaster().Exec(sql); err != nil {
		return appErrorC(err.Error())
	}

	sql = fmt.Sprintf("DELETE FROM keyword WHERE id = %d", keywordId)
	if _, err := s.GetMaster().Exec(sql); err != nil {
		return appErrorC(err.Error())
	}

	return nil
}
