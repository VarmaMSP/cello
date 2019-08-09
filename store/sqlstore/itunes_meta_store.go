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

	_, err := s.Insert([]DbModel{meta}, "itunes_meta")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"itunes_id": meta.ItunesId, "feed_url": meta.FeedUrl},
		)
	}
	return nil
}

func (s *SqlItunesMetaStore) GetStatus(itunesId string) (string, *model.AppError) {
	sql := `SELECT added_to_db FROM itunes_meta WHERE itunes_id = ?`

	var status string
	err := s.GetMaster().QueryRow(sql, itunesId).Scan(&status)
	if err != nil {
		return "", model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.get_status",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"itunes_id": itunesId},
		)
	}

	return status, nil
}

func (s *SqlItunesMetaStore) GetItunesIdList(offset, limit int) ([]string, *model.AppError) {
	sql := `SELECT itunes_id FROM itunes_meta LIMIT ?, ?`

	appErrorC := model.NewAppErrorC(
		"store.sqlstore.sql_itunes_meta_store.get_itunes_id_list",
		http.StatusInternalServerError,
		map[string]string{"offset": strconv.Itoa(offset), "limit": strconv.Itoa(limit)},
	)

	rows, err := s.GetMaster().Query(sql, offset, limit)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}

func (s *SqlItunesMetaStore) SetStatus(itunesId, status string) *model.AppError {
	sql := `UPDATE itunes_meta SET added_to_db = ?, updated_at = ? WHERE itunes_id = ?`

	_, err := s.GetMaster().Exec(sql, status, itunesId)
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_itunes_meta_store.set_status",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"itunes_id": itunesId, "status": status},
		)
	}
	return nil
}
