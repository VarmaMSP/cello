package sqlstore

import (
	"net/http"
	"strings"

	"github.com/varmamsp/cello/model"
)

type SqlCurationStore struct {
	SqlStore
}

func NewSqlCurationStore(store SqlStore) *SqlCurationStore {
	return &SqlCurationStore{store}
}

func (s *SqlCurationStore) Save(curation *model.Curation) *model.AppError {
	curation.PreSave()

	_, err := s.Insert("curation", []DbModel{curation})
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_curation_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"title": curation.Title},
		)
	}
	return nil
}

func (s *SqlCurationStore) SavePodcastCuration(item *model.PodcastCuration) *model.AppError {
	item.PreSave()

	_, err := s.Insert("podcast_curation", []DbModel{item})
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_curation_store.save_podcast_curation", err.Error(), http.StatusInternalServerError,
			map[string]string{"curation_id": item.CurationId, "podcast_id": item.PodcastId},
		)
	}
	return nil
}

func (s *SqlCurationStore) GetAll() ([]*model.Curation, *model.AppError) {
	cols := strings.Join((&model.Curation{}).DbColumns(), ",")
	sql := "SELECT " + cols + " FROM curation ORDER BY created_at DESC"

	var res []*model.Curation
	newItemFields := func() []interface{} {
		tmp := &model.Curation{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.QueryRows(newItemFields, sql); err != nil {
		return nil, model.NewAppError(
			"sqlstore.sql_curation_store.get_all", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return res, nil
}

func (s *SqlCurationStore) GetPodcastsByCuration(curationId string, offset, limit int) ([]*model.PodcastInfo, *model.AppError) {
	cols := strings.Join(DbColumnsWithPrefix(&model.PodcastInfo{}, "podcast"), ",")
	sql := "SELECT " + cols + ` FROM podcast_curation
		INNER JOIN podcast ON podcast.id = podcast_curation.podcast_id
		WHERE podcast_curation.curation_id = ? 
		LIMIT ?, ?`

	var res []*model.PodcastInfo
	newItemFields := func() []interface{} {
		tmp := &model.PodcastInfo{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.QueryRows(newItemFields, sql, curationId, offset, limit); err != nil {
		return nil, model.NewAppError(
			"sqlstore.sql_curation_store.get_podcast_by_curation", err.Error(), http.StatusInternalServerError,
			map[string]string{"curation_id": curationId},
		)
	}
	return res, nil
}
