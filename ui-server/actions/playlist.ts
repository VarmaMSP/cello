import { doFetch } from 'utils/fetch'
import { getPlaylistById } from 'selectors/entities/playlists'
import * as T from 'types/actions'
import { PlaylistPrivacy } from 'types/app'
import * as gtag from 'utils/gtag'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPlaylistLibrary() {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/playlists`,
      }),
    (dispatch, _, { playlists }) => {
      dispatch({ type: T.PLAYLIST_ADD, playlists })
    },
  )
}

export function getPlaylist(playlistId: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/playlists/${playlistId}`,
      }),
    (dispatch, _, { playlists, episodes, podcasts }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.EPISODE_ADD, episodes })
      dispatch({ type: T.PLAYLIST_ADD, playlists })
    },
  )
}

export function loadAndShowAddToPlaylistModal(episodeId: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: '/ajax/service?endpoint=add_to_playlist',
        body: {
          episode_ids: [episodeId],
        },
      }),
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
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=create_playlist`,
        body: { title, privacy, description, episode_ids: [episodeId] },
      }),
    (dispatch, _, { playlists }) => {
      gtag.createPlaylist(title)
      gtag.addEpisodeToPlaylist(title)

      dispatch({ type: T.PLAYLIST_ADD, playlists })
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
    () =>
      doFetch({
        method: 'POST',
        urlPath: '/ajax/service?endpoint=edit_playlist&action=add_episode',
        body: {
          episode_id: episodeId,
          playlist_id: playlistId,
        },
      }),
    (_, getState) => {
      gtag.addEpisodeToPlaylist(
        (getPlaylistById(getState(), playlistId) || {}).title,
      )
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
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=edit_playlist&action=remove_episode`,
        body: {
          episode_id: episodeId,
          playlist_id: playlistId,
        },
      }),
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
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=delete_playlist`,
        body: { playlist_id: playlistId },
      }),
    (dispatch) => {
      dispatch({
        type: T.PLAYLIST_REMOVE,
        playlistIds: [playlistId],
      })
    },
  )
}
