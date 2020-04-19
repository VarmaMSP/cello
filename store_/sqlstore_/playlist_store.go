package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlPlaylistStore struct {
	sqldb.Broker
}

func (s *sqlPlaylistStore) Save(playlist *model.Playlist) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) Get(playlistId int64) (*model.Playlist, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) GetByUser(userId int64) ([]*model.Playlist, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) GetByUserPaginated(userId int64, offset int, limit int) ([]*model.Playlist, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) Update(old *model.Playlist, new *model.Playlist) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) UpdateMemberStats(playlistId int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) Delete(playlistId int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) SaveMember(member *model.PlaylistMember) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) GetMembers(playlistIds []int64, episodeIds []int64) ([]*model.PlaylistMember, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) GetMembersByPlaylist(playlist int64) ([]*model.PlaylistMember, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) ChangeMemberPosition(playlistId int64, episodeId int64, from int, to int) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaylistStore) DeleteMember(playlistId int64, episodeId int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}
