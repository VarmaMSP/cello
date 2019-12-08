import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Playlist, User } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getAllPlaylists(state: AppState) {
  return state.entities.playlists.playlists
}

export function makeGetUserPlaylists() {
  return createSelector<
    AppState,
    $Id<User>,
    MapById<Playlist>,
    $Id<Playlist>[],
    Playlist[]
  >(
    (state, _) => getAllPlaylists(state),
    (state, userId) => state.entities.playlists.byUser[userId] || [],
    (all, ids) => ids.map((id) => all[id]),
  )
}
