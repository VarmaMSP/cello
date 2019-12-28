package app

import (
	"net/http"

	"github.com/varmamsp/cello/model"
)

func (app *App) CreatePlaylist(title, privacy, description string, userId int64) (*model.Playlist, *model.AppError) {
	if privacy != "PUBLIC" && privacy != "PRIVATE" {
		return nil, model.NewAppError(
			"app.save_playlist", "invalid privacy", http.StatusBadRequest,
			map[string]interface{}{"privacy": privacy},
		)
	}

	playlist := &model.Playlist{
		UserId:      userId,
		Title:       title,
		Description: description,
		Privacy:     privacy,
	}

	if err := app.Store.Playlist().Save(playlist); err != nil {
		return nil, err
	}
	return playlist, nil
}

func (app *App) AddEpisodeToPlaylist(playlistId, episodeId int64) *model.AppError {
	playlist, err := app.Store.Playlist().Get(playlistId)
	if err != nil {
		return err
	}

	if err := app.Store.Playlist().SaveMember(&model.PlaylistMember{
		PlaylistId: playlistId,
		EpisodeId:  episodeId,
		Position:   playlist.EpisodeCount + 1,
	}); err != nil {
		return err
	}

	playlistU := playlist
	playlistU.EpisodeCount += 1
	playlistU.UpdatedAt = model.Now()

	return app.Store.Playlist().Update(playlist, playlistU)
}

func (app *App) RemoveEpisodeFromPlaylist(playlistId, episodeId int64) *model.AppError {
	members, err := app.GetPlaylistMembers([]int64{playlistId}, []int64{episodeId})
	if err != nil {
		return err
	}
	if len(members) == 0 {
		return model.NewAppError(
			"app.remove_episode_from_playlist", "member not found", http.StatusNotFound,
			map[string]interface{}{"playlist_id": playlistId, "episode_id": episodeId},
		)
	}

	if err := app.Store.Playlist().ChangeMemberPosition(playlistId, episodeId, members[0].Position, 0); err != nil {
		return err
	}
	if err := app.Store.Playlist().DeleteMember(playlistId, episodeId); err != nil {
		return err
	}

	playlist, err := app.Store.Playlist().Get(playlistId)
	if err != nil {
		return err
	}

	playlistU := playlist
	playlistU.EpisodeCount -= 1
	playlistU.UpdatedAt = model.Now()

	return app.Store.Playlist().Update(playlist, playlistU)
}

func (app *App) JoinPlaylistsToEpisodes(playlists []*model.Playlist, episodeIds []int64) *model.AppError {
	playlistMap := map[int64]*model.Playlist{}
	playlistIds := make([]int64, len(playlists))
	for i, playlist := range playlists {
		playlistIds[i] = playlist.Id
		playlistMap[playlist.Id] = playlist
	}

	members, err := app.GetPlaylistMembers(playlistIds, episodeIds)
	if err != nil {
		return err
	}

	for _, member := range members {
		if playlist, ok := playlistMap[member.PlaylistId]; ok {
			playlist.Members = append(playlist.Members, member)
		}
	}

	return nil
}

func (app *App) GetPlaylist(playlistId int64, loadMembers bool) (*model.Playlist, *model.AppError) {
	playlist, err := app.Store.Playlist().Get(playlistId)
	if err != nil {
		return nil, err
	}

	if loadMembers {
		members, err := app.Store.Playlist().GetMembersByPlaylist(playlistId)
		if err != nil {
			return nil, err
		}
		playlist.Members = members
	}

	return playlist, nil
}

func (app *App) GetPlaylistsByUser(userId int64) ([]*model.Playlist, *model.AppError) {
	return app.Store.Playlist().GetByUser(userId)
}

func (app *App) GetPlaylistMembers(playlistIds, episodeIds []int64) ([]*model.PlaylistMember, *model.AppError) {
	if len(playlistIds) == 0 || len(episodeIds) == 0 {
		return []*model.PlaylistMember{}, nil
	}
	return app.Store.Playlist().GetMembers(playlistIds, episodeIds)
}
