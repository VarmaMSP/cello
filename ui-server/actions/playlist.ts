import * as client from 'client/playlist'
import { getSignedInUserId } from 'selectors/session'
import * as T from 'types/actions'
import { PlaylistPrivacy } from 'types/app'
import { getIdFromUrlParam } from 'utils/format'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPlaylist(playlistId: string) {
  return requestAction(
    () => client.getPlaylist(playlistId),
    (dispatch, _, { playlist, episodes }) => {
      dispatch({
        type: T.PLAYLIST_ADD,
        playlists: [playlist],
      })
      dispatch({
        type: T.PLAYLIST_ADD_EPISODES,
        playlistId: playlist.id,
        episodeIds: episodes.map((x) => x.id),
      })
    },
  )
}

export function loadAndShowAddToPlaylistModal(episodeId: string) {
  return requestAction(
    () => client.getSignedInUserPlaylists(),
    (dispatch, _, { playlists }) => {
      dispatch({
        type: T.PLAYLIST_ADD,
        playlists,
      })
      dispatch({
        type: T.MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      })
    },
  )
}

export function createPlaylist(
  title: string,
  privacy: PlaylistPrivacy,
  episodeId: string,
) {
  return requestAction(
    () => client.createPlaylist(title, privacy),
    (dispatch, getState, { urlParam }) => {
      const id = getIdFromUrlParam(urlParam)
      const userId = getSignedInUserId(getState())

      dispatch({
        type: T.PLAYLIST_ADD,
        playlists: [{ id, urlParam, title, privacy, userId }],
      })
      dispatch({
        type: T.MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      })
    },
    { requestId: RequestId.createPlaylist() },
  )
}
