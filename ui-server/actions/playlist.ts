import * as client from 'client/playlist'
import { getCurrentUserId } from 'selectors/entities/users'
import * as T from 'types/actions'
import { PlaylistPrivacy } from 'types/app'
import { getIdFromUrlParam } from 'utils/format'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPlaylist(playlistId: string) {
  return requestAction(
    () => client.getPlaylist(playlistId),
    (dispatch, _, { playlist, episodes }) => {
      dispatch({ type: T.RECEIVED_PLAYLIST, playlist })
      dispatch({
        type: T.RECEIVED_PLAYLIST_EPISODES,
        playlistId: playlist.id,
        episodes,
      })
    },
  )
}

export function loadAndShowAddToPlaylistModal(episodeId: string) {
  return requestAction(
    () => client.getUserPlaylists(),
    (dispatch, getState, { playlists }) => {
      dispatch({
        type: T.RECEIVED_USER_PLAYLISTS,
        userId: getCurrentUserId(getState()),
        playlists,
      })
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
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
      const userId = getCurrentUserId(getState())

      dispatch({
        type: T.RECEIVED_PLAYLIST,
        playlist: { id, urlParam, title, privacy, userId },
      })
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      })
    },
    { requestId: RequestId.createPlaylist() },
  )
}

export function addEpisodeToPlaylist(episodeId: string, playlistId: string) {
  return requestAction(
    () => client.addEpisodeToPlaylist(episodeId, playlistId),
    (dispatch) => {
      dispatch({
        type: T.ADD_EPISODE_TO_PLAYLIST,
        episodeId,
        playlistId,
      })
    },
  )
}
