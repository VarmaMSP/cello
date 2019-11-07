import client from 'client'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function createPlaylist(
  title: string,
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
) {
  return requestAction(
    () => client.createPlaylist(title, privacy),
    (dispatch, { playlist }) => {
      dispatch({ type: T.RECEIVED_PLAYLIST, playlist })
    },
    { type: T.CREATE_PLAYLIST_REQUEST },
    { type: T.CREATE_PLAYLIST_SUCCESS },
    { type: T.CREATE_PLAYLIST_FAILURE },
  )
}
