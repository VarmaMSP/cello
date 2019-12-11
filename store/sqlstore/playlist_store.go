package sqlstore

import (
	"fmt"
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPlaylistStore struct {
	SqlStore
}

func NewSqlPlaylistStore(store SqlStore) store.PlaylistStore {
	return &SqlPlaylistStore{store}
}

func (s *SqlPlaylistStore) Save(playlist *model.Playlist) *model.AppError {
	playlist.PreSave()

	id, err := s.InsertWithoutPK("playlist", playlist)
	if err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.save", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"title": playlist.Title, "user_id": playlist.UserId},
		)
	}
	playlist.Id = id
	return nil
}

func (s *SqlPlaylistStore) Get(playlistId int64) (*model.Playlist, *model.AppError) {
	playlist := &model.Playlist{}
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE id = %d",
		joinStrings(playlist.DbColumns(), ","), playlistId,
	)

	if err := s.GetMaster().QueryRow(sql).Scan(playlist.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_playlist_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return playlist, nil
}

func (s *SqlPlaylistStore) GetByUserPaginated(userId int64, offset, limit int) (res []*model.Playlist, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE user_id = %d LIMIT %d, %d",
		joinStrings((&model.Playlist{}).DbColumns(), ","), userId, offset, limit,
	)

	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_by_user_paginated", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId, "offset": offset, "limit": limit},
		)
	}
	return
}

func (s *SqlPlaylistStore) SaveMember(member *model.PlaylistMember) *model.AppError {
	member.PreSave()

	if _, err := s.Insert("playlist_member", []model.DbModel{member}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.save_member", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": member.PlaylistId, "episode_id": member.EpisodeId},
		)
	}
	return nil
}

func (s *SqlPlaylistStore) DeleteMember(playlistId, episodeId int64) *model.AppError {
	sql := fmt.Sprintf(
		"UPDATE playlist_member SET active = 0, updated_at = %d WHERE playlist_id = %d AND episode_id = %d",
		model.Now(), playlistId, episodeId,
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.delete_member", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId, "episode_id": episodeId},
		)
	}
	return nil
}
