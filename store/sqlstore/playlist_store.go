package sqlstore

import (
	"fmt"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type sqlPlaylistStore struct {
	store.Store
	sqldb.Broker
}

func (s *sqlPlaylistStore) Save(playlist *model.Playlist) *model.AppError {
	playlist.PreSave()

	res, err := s.Insert_("playlist", playlist)
	if err != nil {
		return model.New500Error("sql_store.sql_playlist_store.save", err.Error(), nil)
	}
	playlist.Id, _ = res.LastInsertId()
	return nil
}

func (s *sqlPlaylistStore) Get(playlistId int64) (*model.Playlist, *model.AppError) {
	res := &model.Playlist{}
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE id = %d",
		cols(res), playlistId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_playlist_store.get", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlPlaylistStore) GetByUser(userId int64) (res []*model.Playlist, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE user_id = %d",
		cols(&model.Playlist{}), userId,
	)
	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playlist_store.get_by_user", err.Error(), nil)
	}
	return
}

func (s *sqlPlaylistStore) GetByUserPaginated(userId int64, offset int, limit int) (res []*model.Playlist, appE *model.AppError) {
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist WHERE user_id = %d LIMIT %d, %d",
		cols(&model.Playlist{}), userId, offset, limit,
	)
	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playlist_store.get_by_user_paginated", err.Error(), nil)
	}
	return
}

func (s *sqlPlaylistStore) Update(old *model.Playlist, new *model.Playlist) *model.AppError {
	if _, err := s.Patch("playlist", old, new); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.upate", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaylistStore) UpdateMemberStats(playlistId int64) *model.AppError {
	count, err := s.GetMemberCount(playlistId)
	if err != nil {
		return err
	}

	var previewImage string
	if firstMember, err := s.GetMemberByPosition(playlistId, 1); err != nil {
		return err
	} else if firstMember == nil {
		previewImage = "placeholder"
	} else if episode, err := s.Episode().Get(firstMember.EpisodeId); err != nil {
		return err
	} else if podcast, err := s.Podcast().Get(episode.PodcastId); err != nil {
		return err
	} else {
		previewImage = hashid.UrlParam(podcast.Title, podcast.Id)
	}

	sql := fmt.Sprintf(
		`UPDATE playlist SET episode_count = %d, preview_image = "%s" WHERE id = %d`,
		count, previewImage, playlistId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.update_member_stats", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaylistStore) Delete(playlistId int64) *model.AppError {
	sql := fmt.Sprintf(`DELETE FROM playlist_member WHERE playlist_id = %d`, playlistId)
	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.delete", err.Error(), nil)
	}

	sql = fmt.Sprintf(`DELETE FROM playlist WHERE id = %d`, playlistId)
	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.delete", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaylistStore) SaveMember(member *model.PlaylistMember) *model.AppError {
	member.PreSave()

	if _, err := s.Insert("playlist_member", member); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.add_member", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaylistStore) GetMember(playlistId, episodeId int64) (*model.PlaylistMember, *model.AppError) {
	res := &model.PlaylistMember{}
	sql := fmt.Sprintf(
		`SELECT %s FROM playlist_member WHERE playlist_id = %d AND episode_id = %d`,
		cols(res), playlistId, episodeId,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_playlist_store.get_member", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlPlaylistStore) GetMembers(playlistIds, episodeIds []int64) (res []*model.PlaylistMember, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playlist_member WHERE playlist_id IN (%s) AND episode_id IN (%s)`,
		cols(&model.PlaylistMember{}), joinInt64s(playlistIds), joinInt64s(episodeIds),
	)
	copyTo := func() []interface{} {
		tmp := &model.PlaylistMember{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playlist_store.get_members", err.Error(), nil)
	}
	return
}

func (s *sqlPlaylistStore) GetMembersByPlaylist(playlistId int64) (res []*model.PlaylistMember, appE *model.AppError) {
	sql := fmt.Sprintf(
		`SELECT %s FROM playlist_member WHERE playlist_id = %d ORDER BY position ASC`,
		cols(&model.PlaylistMember{}), playlistId,
	)
	copyTo := func() []interface{} {
		tmp := &model.PlaylistMember{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.New500Error("sql_store.sql_playlist_store.get_members_by_playlist", err.Error(), nil)
	}
	return
}

func (s *sqlPlaylistStore) GetMemberByPosition(playlistId int64, position int) (*model.PlaylistMember, *model.AppError) {
	res := &model.PlaylistMember{}
	sql := fmt.Sprintf(
		"SELECT %s FROM playlist_member WHERE playlist_id = %d AND position = %d",
		cols(res), playlistId, position,
	)

	if err := s.QueryRow(res.FieldAddrs(), sql); err != nil {
		return nil, model.New500Error("sql_store.sql_playlist_store.get_member_by_position", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlPlaylistStore) GetMemberCount(playlistId int64) (int, *model.AppError) {
	res := 0
	sql := fmt.Sprintf(`SELECT COUNT(*) FROM playlist_member WHERE playlist_id = %d`, playlistId)

	if err := s.QueryRow([]interface{}{&res}, sql); err != nil {
		return 0, model.New500Error("sql_store.sql_playlist_store.get_member_count", err.Error(), nil)
	}
	return res, nil
}

func (s *sqlPlaylistStore) ChangeMemberPosition(playlistId int64, episodeId int64, from int, to int) *model.AppError {
	// Change positions of members between to and from
	if to < from {
		sql := fmt.Sprintf(
			`UPDATE playlist_member
				SET position = position + 1, updated_at = %d
				WHERE playlist_id = %d AND position >= %d AND position < %d`,
			datetime.Unix(), playlistId, to, from,
		)

		if err := s.Exec(sql); err != nil {
			return model.New500Error("sql_store.sql_playlist_store.change_member_position", err.Error(), nil)
		}

	} else {
		sql := fmt.Sprintf(
			`UPDATE playlist_member
				SET position = position - 1, updated_at = %d
				WHERE playlist_id = %d AND position > %d AND position <= %d`,
			datetime.Unix(), playlistId, from, to,
		)

		if err := s.Exec(sql); err != nil {
			return model.New500Error("sql_store.sql_playlist_store.change_member_position", err.Error(), nil)
		}
	}

	// Change Member Position
	sql := fmt.Sprintf(
		`UPDATE playlist_member
			SET position = %d, updated_at = %d
			WHERE playlist_id = %d AND episode_id = %d`,
		to, datetime.Unix(), playlistId, episodeId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("sql_store.sql_playlist_store.change_member_position", err.Error(), nil)
	}
	return nil
}

func (s *sqlPlaylistStore) DeleteMember(playlistId int64, episodeId int64) *model.AppError {
	sql := fmt.Sprintf(
		`DELETE FROM playlist_member WHERE playlist_id = %d AND episode_id = %d`,
		playlistId, episodeId,
	)

	if err := s.Exec(sql); err != nil {
		return model.New500Error("store.sqlstore.sql_playlist_store.delete_member", err.Error(), nil)
	}
	return nil
}
