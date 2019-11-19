import client from 'client'
import { getCurrentUserId } from 'selectors/entities/users'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getCurrentUserPlaylists() {
  return requestAction(
    () => client.getCurrentUserPlaylists(),
    (dispatch, getState, { playlists }) => {
      dispatch({
        type: T.RECEIVED_PLAYLISTS,
        playlists,
        userId: getCurrentUserId(getState()),
      })
    },
  )
}

export function createPlaylist(
  title: string,
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
) {
  return requestAction(
    () => client.createPlaylist(title, privacy),
    (dispatch, getState, { playlist }) => {
      dispatch({
        type: T.RECEIVED_PLAYLISTS,
        playlists: [playlist],
        userId: getCurrentUserId(getState()),
      })
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId: '',
      })
    },
  )
}
