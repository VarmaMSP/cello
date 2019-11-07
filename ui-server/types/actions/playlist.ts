import { Episode, Playlist } from 'types/app'

export const CREATE_PLAYLIST_REQUEST = 'CREATE_PLAYLIST_REQUEST'
export const CREATE_PLAYLIST_SUCCESS = 'CREATE_PLAYLIST_SUCCESS'
export const CREATE_PLAYLIST_FAILURE = 'CREATE_PLAYLIST_FAILURE'

export const RECEIVED_PLAYLIST = 'RECEIVED_PLAYLIST'
export const RECEIVED_PLAYLIST_EPISODES = 'RECEIVED_PLAYLIST_EPISODES'

export interface CreatePlaylistRequestAction {
  type: typeof CREATE_PLAYLIST_REQUEST
}

export interface CreatePlaylistSuccessAction {
  type: typeof CREATE_PLAYLIST_SUCCESS
}

export interface CreatePlaylistFailureAction {
  type: typeof CREATE_PLAYLIST_FAILURE
}

export interface ReceivedPlaylistAction {
  type: typeof RECEIVED_PLAYLIST
  playlist: Playlist
}

export interface ReceivedPlaylistEpisodesAction {
  type: typeof RECEIVED_PLAYLIST_EPISODES
  playlistId: string
  episodes: Episode[]
}

export type PlaylistActionTypes =
  | CreatePlaylistRequestAction
  | CreatePlaylistSuccessAction
  | CreatePlaylistFailureAction
  | ReceivedPlaylistAction
  | ReceivedPlaylistEpisodesAction
