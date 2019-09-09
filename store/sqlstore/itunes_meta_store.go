package sqlstore

import (
	"net/http"
	"strconv"

	"github.com/varmamsp/cello/model"
)

type SqlItunesMetaStore struct {
	SqlStore
}

func NewSqlItunesMetaStore(store SqlStore) *SqlItunesMetaStore {
	return &SqlItunesMetaStore{store}
}

func (s *SqlItunesMetaStore) Save(meta *model.ItunesMeta) *model.AppError {
	meta.PreSave()

	if _, err := s.Insert("itunes_meta", []DbModel{meta}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"itunes_id": meta.ItunesId, "feed_url": meta.FeedUrl},
		)
	}
	return nil
}

func (s *SqlItunesMetaStore) Update(old, new *model.ItunesMeta) *model.AppError {
	if _, err := s.UpdateChanges("itunes_meta", old, new, "itunes_id = ?", new.ItunesId); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.update", err.Error(), http.StatusInternalServerError,
			map[string]string{"itunes_id": new.ItunesId},
		)
	}
	return nil
}

func (s *SqlItunesMetaStore) GetItunesIdList(offset, limit int) (res []string, appE *model.AppError) {
	sql := `SELECT itunes_id FROM itunes_meta LIMIT ?, ?`

	copyTo := func() []interface{} {
		tmp := ""
		res = append(res, tmp)
		return []interface{}{tmp}
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.get_itunes_id_list", err.Error(), http.StatusInternalServerError,
			map[string]string{"offset": strconv.Itoa(offset), "limit": strconv.Itoa(limit)},
		)
	}
	return
}
