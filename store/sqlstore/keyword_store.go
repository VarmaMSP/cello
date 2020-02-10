package sqlstore

import (
	"database/sql"
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
	} else if k != nil {
		return k, nil
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

func (s *SqlKeywordStore) GetByText(text string) (*model.Keyword, *model.AppError) {
	keyword := &model.Keyword{}
	sql_ := fmt.Sprintf(
		"SELECT %s FROM keyword WHERE text = '%s'",
		joinStrings(keyword.DbColumns(), ","), text,
	)

	if err := s.GetMaster().QueryRow(sql_).Scan(keyword.FieldAddrs()...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, model.NewAppError(
			"store.sqlstore.sql_keyword_store.get_by_text", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"text": text},
		)
	}

	return keyword, nil
}
