package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

type SqlKeywordStore struct {
	SqlStore
}

func NewSqlKeywordStore(store SqlStore) *SqlKeywordStore {
	return &SqlKeywordStore{store}
}

func (s *SqlKeywordStore) Save(keyword *model.Keyword) *model.AppError {
	keyword.PreSave()

	id, err := s.InsertWithoutPK("keyword", keyword)
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_keyword_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"source": keyword.Source, "source_id": keyword.SourceId, "text": keyword.Text},
		)
	}
	keyword.Id = id
	return nil
}
