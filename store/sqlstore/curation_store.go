package sqlstore

import (
	"net/http"

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

func (s *SqlCurationStore) Get(curationId string) (*model.Curation, *model.AppError) {
	curation := &model.Curation{}
	sql := "SELECT " + Cols(curation) + " FROM curation WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, curationId).Scan(curation.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_curation_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"curation_id": curationId},
		)
	}
	return curation, nil
}

func (s *SqlCurationStore) GetAll() (res []*model.Curation, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Curation{}) + " FROM curation ORDER BY created_at DESC"

	copyTo := func() []interface{} {
		tmp := &model.Curation{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_curation_store.get_all", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return
}

func (s *SqlCurationStore) Delete(curationId string) *model.AppError {
	sql := "DELETE FROM curation WHERE id = ?"

	if _, err := s.GetMaster().Exec(sql, curationId); err != nil {
		return model.NewAppError(
			"sqlstore.sql_curation_store.delete", err.Error(), http.StatusInternalServerError,
			map[string]string{"curation_id": curationId},
		)
	}
	return nil
}
