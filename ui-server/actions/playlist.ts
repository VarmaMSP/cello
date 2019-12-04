import * as client from 'client/playlist'
import { getCurrentUserId } from 'selectors/entities/users'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getUserPlaylists() {
  return requestAction(
    () => client.getUserPlaylists(),
    (dispatch, _, { playlists }) => {
      dispatch({ type: T.RECEIVED_USER_PLAYLISTS, playlists })
    },
  )
}

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

export function createPlaylist(
  title: string,
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
) {
  return requestAction(
    () => client.createPlaylist(title, privacy),
    (dispatch, getState, { urlParam }) => {
      dispatch({
        type: T.RECEIVED_PLAYLIST,
        playlist: {
          id: urlParam,
          title,
          privacy,
          userId: getCurrentUserId(getState()),
        },
      })
    },
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
