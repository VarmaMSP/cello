package sqlstore

import (
	"net/http"
	"strings"

	"github.com/varmamsp/cello/model"
)

type SqlPodcastCurationStore struct {
	SqlStore
}

func NewSqlPodcastCurationStore(store SqlStore) *SqlPodcastCurationStore {
	return &SqlPodcastCurationStore{store}
}

func (s *SqlPodcastCurationStore) Save(curation *model.PodcastCuration) *model.AppError {
	curation.PreSave()

	_, err := s.Insert([]DbModel{curation}, "podcast_curation")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_curation_store.save",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"title": curation.Title},
		)
	}
	return nil
}

func (s *SqlPodcastCurationStore) SaveItem(item *model.PodcastCurationItem) *model.AppError {
	item.PreSave()

	_, err := s.Insert([]DbModel{item}, "podcast_curation_item")
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_podcast_curation_store.save_item",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"curation_id": item.CurationId, "podcast_id": item.PodcastId},
		)
	}
	return nil
}

func (s *SqlPodcastCurationStore) GetAll() ([]*model.PodcastCuration, *model.AppError) {
	m := &model.PodcastCuration{}
	sql := "SELECT " + strings.Join(m.DbColumns(), ",") + " FROM podcast_curation ORDER BY created_at DESC"

	appErrorC := model.NewAppErrorC(
		"sqlstore.sql_podcast_curation_store.get_all",
		http.StatusInternalServerError,
		nil,
	)

	rows, err := s.GetMaster().Query(sql)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []*model.PodcastCuration
	for rows.Next() {
		tmp := &model.PodcastCuration{}
		if err := rows.Scan(tmp.FieldAddrs()...); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}

func (s *SqlPodcastCurationStore) GetPodcastsByCuration(curationId string, offset, limit int) ([]*model.PodcastInfo, *model.AppError) {
	m := &model.PodcastInfo{}
	sql := "SELECT " + strings.Join(DbColumnsWithPrefix(m, "podcast"), ",") + ` FROM podcast_curation_item
		INNER JOIN podcast ON podcast.id = podcast_curation_item.podcast_id
		WHERE podcast_curation_item.curation_id = ? 
		LIMIT ?, ?`

	appErrorC := model.NewAppErrorC(
		"sqlstore.sql_podcast_curation_store.get_all",
		http.StatusInternalServerError,
		map[string]string{"curation_id": curationId},
	)

	rows, err := s.GetMaster().Query(sql, curationId, offset, limit)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []*model.PodcastInfo
	for rows.Next() {
		tmp := &model.PodcastInfo{}
		if err := rows.Scan(tmp.FieldAddrs()...); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}
