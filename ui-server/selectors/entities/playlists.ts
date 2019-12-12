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
    {
      byCreateDateDesc: { [offset: string]: string[] }
      receivedAll: 'create_date_desc'[]
    },
    {
      playlists: Playlist[]
      receivedAll: boolean
    }
  >(
    (state, _) => getAllPlaylists(state),
    (state, userId) => state.entities.playlists.playlistsByUser[userId] || {},
    (playlists, x) => {
      return {
        playlists: Object.values(x.byCreateDateDesc || {}).reduce<Playlist[]>(
          (acc, playlistIds) => [
            ...acc,
            ...playlistIds.map((id) => playlists[id]),
          ],
          [],
        ),
        receivedAll: (x.receivedAll || []).includes('create_date_desc'),
      }
    }
  )
}
