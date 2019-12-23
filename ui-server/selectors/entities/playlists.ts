import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Playlist, PlaylistMember, User } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getPlaylistById(state: AppState, playlistId: string) {
  return state.entities.playlists.byId[playlistId]
}

export function makeGetPlaylistsByUser() {
  return createSelector<
    AppState,
    $Id<User>,
    MapById<Playlist>,
    $Id<Playlist>[],
    Playlist[]
  >(
    (state) => state.entities.playlists.byId,
    (state, userId) => state.entities.playlists.byUserId[userId],
    (all, ids) => ids.map((id) => all[id]),
  )
}

export function makeGetEpisodesInPlaylist() {
  return createSelector<
    AppState,
    $Id<Playlist>,
    MapById<PlaylistMember>,
    $Id<PlaylistMember>[],
    string[]
  >(
    (state) => state.entities.playlistMember.byId,
    (state, playlistId) =>
      state.entities.playlistMember.byPlaylistId[playlistId],
    (all, ids) => ids.map((id) => all[id].episodeId),
  )
}
