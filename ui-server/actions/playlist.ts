import * as client from 'client/playlist'
import { getPlaylistById } from 'selectors/entities/playlists'
import * as T from 'types/actions'
import { PlaylistPrivacy } from 'types/app'
import * as gtag from 'utils/gtag'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPlaylistLibrary() {
  return requestAction(
    () => client.getPlaylistLibrary(),
    (dispatch, _, { playlists }) => {
      dispatch({
        type: T.PLAYLIST_ADD,
        playlists,
      })
    },
  )
}

export function getPlaylist(playlistId: string) {
  return requestAction(
    () => client.getPlaylist(playlistId),
    (dispatch, _, { playlist, episodes, podcasts }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.EPISODE_ADD, episodes })
      dispatch({
        type: T.PLAYLIST_ADD,
        playlists: [playlist],
      })
    },
  )
}

export function loadAndShowAddToPlaylistModal(episodeId: string) {
  return requestAction(
    () => client.serviceAddToPlaylist(episodeId),
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
  description: string,
  episodeId: string,
) {
  return requestAction(
    () => client.serviceCreatePlaylist(title, privacy, description, episodeId),
    (dispatch, _, { playlist }) => {
      gtag.createPlaylist(title)
      gtag.addEpisodeToPlaylist(title)

      dispatch({
        type: T.PLAYLIST_ADD,
        playlists: [playlist],
      })
      dispatch({
        type: T.MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      })
    },
    { requestId: RequestId.createPlaylist() },
  )
}

export function addEpisodeToPlaylist(playlistId: string, episodeId: string) {
  return requestAction(
    () => client.serviceAddEpisodeToPlaylist(playlistId, episodeId),
    (_, getState) => {
      gtag.addEpisodeToPlaylist((getPlaylistById(getState(), playlistId) || {}).title)
    },
    {
      preAction: {
        type: T.PLAYLIST_ADD_EPISODES,
        playlistId: playlistId,
        episodeIds: [episodeId],
      },
    },
  )
}

export function removeEpisodeFromPlaylist(
  playlistId: string,
  episodeId: string,
) {
  return requestAction(
    () => client.serviceRemoveEpisodeFromPlaylist(playlistId, episodeId),
    () => {},
    {
      preAction: {
        type: T.PLAYLIST_REMOVE_EPISODES,
        playlistId: playlistId,
        episodeIds: [episodeId],
      },
    },
  )
}

export function deletePlaylist(playlistId: string) {
  return requestAction(
    () => client.serviceDeletePlaylist(playlistId),
    (dispatch) => {
      dispatch({
        type: T.PLAYLIST_REMOVE,
        playlistIds: [playlistId],
      })
    },
  )
}
