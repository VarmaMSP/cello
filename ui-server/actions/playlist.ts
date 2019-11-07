import client from 'client'
import { getCurrentUserId } from 'selectors/entities/users'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getCurrentUserPlaylists() {
  return requestAction(
    () => client.getCurrentUserPlaylists(),
    (dispatch, { playlists }, getState) => {
      dispatch({
        type: T.RECEIVED_PLAYLISTS,
        playlists,
        userId: getCurrentUserId(getState()),
      })
    },
    { type: T.GET_USER_PLAYLISTS_REQUEST },
    { type: T.GET_USER_PLAYLISTS_SUCCESS },
    { type: T.GET_USER_PLAYLISTS_FAILURE },
  )
}

export function createPlaylist(
  title: string,
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
) {
  return requestAction(
    () => client.createPlaylist(title, privacy),
    (dispatch, { playlist }, getState) => {
      dispatch({
        type: T.RECEIVED_PLAYLISTS,
        playlists: [playlist],
        userId: getCurrentUserId(getState()),
      })
    },
    { type: T.CREATE_PLAYLIST_REQUEST },
    { type: T.CREATE_PLAYLIST_SUCCESS },
    { type: T.CREATE_PLAYLIST_FAILURE },
  )
}
