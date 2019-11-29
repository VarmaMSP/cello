package sqlstore

import (
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
	sql := "SELECT " + Cols(playlist) + " FROM playlist WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, playlistId).Scan(playlist.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_playlist_store.get", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"playlist_id": playlistId},
		)
	}
	return playlist, nil
}

func (s *SqlPlaylistStore) GetByUser(userId int64) (res []*model.Playlist, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Playlist{}) + ` FROM playlist WHERE user_id = ?`

	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_by_user", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"user_id": userId},
		)
	}
	return
}

func (s *SqlPlaylistStore) AddMember(member *model.PlaylistMember) *model.AppError {
	member.PreSave()

	if _, err := s.Insert("playlist_member", []model.DbModel{member}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.add_member", err.Error(), http.StatusInternalServerError,
			map[string]interface{}{"episode_id": member.EpisodeId},
		)
	}
	return nil
}

func (s *SqlPlaylistStore) DeleteMember(playlistId, episodeId int64) *model.AppError {
	return nil
}
