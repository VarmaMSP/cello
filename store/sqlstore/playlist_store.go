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

func (s *SqlPlaylistStore) GetByUser(userId int64) (res []*model.Playlist, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE user_id = %d",
		joinStrings((&model.Playlist{}).DbColumns(), ","), userId,
	)

	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_by_user", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId},
		)
	}
	return
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

func (s *SqlPlaylistStore) UpdateMemberStats(stats *model.PlaylistMemberStats) *model.AppError {
	sql := fmt.Sprintf(
		"UPDATE playlist SET preview_image = %s, episode_count = %d, updated_at = %d WHERE id = %d",
		stats.PreviewImage, stats.EpisodeCount, model.Now(), stats.Id,
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.update_member_stats", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": stats.Id},
		)
	}
	return nil
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

func (s *SqlPlaylistStore) GetMembers(episodeId, playlistIds []int64) (res []*model.PlaylistMember, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist_member WHERE episode_id IN (%s) AND playlist_id IN (%s)",
		joinStrings((&model.PlaylistMember{}).DbColumns(), ","),
		joinInt64s(episodeId, ","),
		joinInt64s(playlistIds, ","),
	)

	copyTo := func() []interface{} {
		tmp := &model.PlaylistMember{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_memberships", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": episodeId},
		)
	}
	return
}

func (s *SqlPlaylistStore) GetMembersByPlaylist(playlistId int64) (res []*model.PlaylistMember, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist_member WHERE playlist_id = %d",
		joinStrings((&model.PlaylistMember{}).DbColumns(), ","), playlistId,
	)

	copyTo := func() []interface{} {
		tmp := &model.PlaylistMember{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_members_by_playlist", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return
}

func (s *SqlPlaylistStore) GetMembersCount(playlistId int64) (count int, appE *model.AppError) {
	sql := fmt.Sprintf("SELECT COUNT(*) FROM playlist_member WHERE playlist_id = %d", playlistId)

	if err := s.GetMaster().QueryRow(sql).Scan(&count); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_members_count", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return
}

func (s *SqlPlaylistStore) ChangeMemberPosition(playlistId, episodeId int64, from, to int) *model.AppError {
	var sql string

	// Modify positions for other members
	if to < from {
		sql = fmt.Sprintf(
			"UPDATE playlist_member SET position = position + 1, updated_at = %d WHERE position < %d AND position >= %d",
			model.Now(), from, to,
		)
	} else if to > from {
		sql = fmt.Sprintf(
			"UPDATE playlist_member SET position = position - 1, updated_at = %d WHERE position > %d AND position <= %d",
			model.Now(), from, to,
		)
	} else {
		return nil
	}

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.change_member_position", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId, "episode_id": episodeId, "from": from, "to": to},
		)
	}

	// set position of member
	sql = fmt.Sprintf(
		"UPDATE playlist_member SET position = %d, updated_at = %d WHERE playlist_id = %d AND episode_id = %d",
		to, model.Now(), playlistId, episodeId,
	)

	if _, err := s.GetMaster().Exec(sql); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.change_member_position", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId, "episode_id": episodeId, "from": from, "to": to},
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
