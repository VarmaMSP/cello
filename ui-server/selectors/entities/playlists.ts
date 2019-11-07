import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Playlist, User } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getAllPlaylists(state: AppState) {
  return state.entities.playlist.playlists
}

export function makeGetUserPlaylists() {
  return createSelector<
    AppState,
    $Id<User>,
    $Id<Playlist>[],
    MapById<Playlist>,
    Playlist[]
  >(
    (state, userId) => state.entities.playlist.byUser[userId] || [],
    (state, _) => getAllPlaylists(state),
    (ids, playlists) => ids.map((id) => playlists[id]),
  )
}
