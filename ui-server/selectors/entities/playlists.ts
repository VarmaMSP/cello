import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Playlist, User } from 'models'
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
    (state, userId) => state.entities.playlists.byUserId[userId] || [],
    (all, ids) => ids.map((id) => all[id]).filter((x) => !!x),
  )
}
