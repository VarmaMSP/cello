package app

import "github.com/varmamsp/cello/model"

func (a *App) CreatePlaylistWithEpisodes(playlist *model.Playlist, episodeIds []int64) (*model.Playlist, *model.AppError) {
	if err := a.Store.Playlist().Save(playlist); err != nil {
		return nil, err
	}

	for i, episodeId := range episodeIds {
		member := &model.PlaylistMember{
			PlaylistId: playlist.Id,
			EpisodeId:  episodeId,
			Position:   i + 1,
		}
		if err := a.Store.Playlist().SaveMember(member); err != nil {
			return nil, err
		}
		playlist.Members = append(playlist.Members, member)
	}

	if err := a.Store.Playlist().UpdateMemberStats(playlist.Id); err != nil {
		return nil, err
	}

	return playlist, nil
}

func (a *App) AddEpisodeToPlaylist(playlistId, episodeId int64) *model.AppError {
	count, err := a.Store.Playlist().GetMemberCount(playlistId)
	if err != nil {
		return err
	}

	if err := a.Store.Playlist().SaveMember(&model.PlaylistMember{
		PlaylistId: playlistId,
		EpisodeId:  episodeId,
		Position:   count + 1,
	}); err != nil {
		return err
	}

	return a.Store.Playlist().UpdateMemberStats(playlistId)
}

func (a *App) RemoveEpisodeFromPlaylist(playlistId, episodeId int64) *model.AppError {
	member, err := a.Store.Playlist().GetMember(playlistId, episodeId)
	if err != nil {
		return err
	}

	if err := a.Store.Playlist().ChangeMemberPosition(playlistId, episodeId, member.Position, 10000); err != nil {
		return err
	}

	if err := a.Store.Playlist().DeleteMember(playlistId, episodeId); err != nil {
		return err
	}

	return a.Store.Playlist().UpdateMemberStats(playlistId)
}

func (a *App) DeletePlaylist(playlistId int64) *model.AppError {
	if err := a.Store.Playlist().Delete(playlistId); err != nil {
		return err
	}

	return nil
}

func (a *App) GetPlaylist(playlistId int64) (*model.Playlist, *model.AppError) {
	playlist, err := a.Store.Playlist().Get(playlistId)
	if err != nil {
		return nil, err
	}

	members, err := a.Store.Playlist().GetMembersByPlaylist(playlistId)
	if err != nil {
		return nil, err
	}
	playlist.Members = members

	return playlist, nil
}

func (a *App) GetPlaylistsByUser(userId int64, includeEpisodes ...int64) ([]*model.Playlist, *model.AppError) {
	playlists, err := a.Store.Playlist().GetByUser(userId)
	if err != nil {
		return nil, err
	}

	if len(includeEpisodes) == 0 {
		return playlists, nil
	}

	playlistMap := map[int64]*model.Playlist{}
	playlistIds := make([]int64, len(playlists))
	for i, playlist := range playlists {
		playlistIds[i] = playlist.Id
		playlistMap[playlist.Id] = playlist
	}

	members, err := a.Store.Playlist().GetMembers(playlistIds, includeEpisodes)
	if err != nil {
		return nil, err
	}

	for _, member := range members {
		if playlist, ok := playlistMap[member.PlaylistId]; ok {
			playlist.Members = append(playlist.Members, member)
		}
	}

	return playlists, nil
}

func (a *App) HasPermissionToPlaylist(userId int64, playlistId int64) (bool, *model.AppError) {
	if playlist, err := a.Store.Playlist().Get(playlistId); err != nil {
		return false, err
	} else {
		return playlist.UserId == userId, nil
	}
}

func (a *App) HasPermissionToViewPlaylist(session *model.Session, playlist *model.Playlist) bool {
	if playlist.Privacy == "PRIVATE" {
		return session != nil && session.UserId == playlist.UserId
	}
	return true
}
