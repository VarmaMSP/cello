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

	if _, err := s.Insert("playlist", []model.DbModel{playlist}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"title": playlist.Title, "user_id": playlist.CreatedBy},
		)
	}
	return nil
}

func (s *SqlPlaylistStore) SaveItem(playlistItem *model.PlaylistItem) *model.AppError {
	playlistItem.PreSave()

	if _, err := s.Insert("playlist_item", []model.DbModel{playlistItem}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_playlist_store.save_item", err.Error(), http.StatusInternalServerError,
			map[string]string{"playlist_id": playlistItem.PlaylistId, "episode_id": playlistItem.EpisodeId},
		)
	}
	return nil
}

func (s *SqlPlaylistStore) Get(playlistId string) (*model.Playlist, *model.AppError) {
	playlist := &model.Playlist{}
	sql := "SELECT " + Cols(playlist) + " FROM playlist WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, playlistId).Scan(playlist.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_playlist_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"playlist_id": playlistId},
		)
	}
	return playlist, nil
}

func (s *SqlPlaylistStore) GetAllByUser(userId string) (res []*model.Playlist, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Playlist{}) + ` FROM playlist WHERE created_by = ?`

	copyTo := func() []interface{} {
		tmp := &model.Playlist{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, userId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_all_by_user", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": userId},
		)
	}
	return
}

func (s *SqlPlaylistStore) GetAllEpisodesInPlaylist(playlistId string) (res []*model.Episode, appE *model.AppError) {
	sql := "SELECT " + Cols(&model.Episode{}, "episode") + `FROM episode
		INNER JOIN playlist_item ON playlist_item.episode_id = episode.id
		WHERE playlist_item.playlist_id = ?
		ORDER BY playlist_item.createdAt DESC`

	copyTo := func() []interface{} {
		tmp := &model.Episode{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql, playlistId); err != nil {
		appE = model.NewAppError(
			"store.sqlstore.sql_playlist_store.get_all_episodes_in_playlist", err.Error(), http.StatusInternalServerError,
			map[string]string{"playlist_id": playlistId},
		)
	}
	return
}
