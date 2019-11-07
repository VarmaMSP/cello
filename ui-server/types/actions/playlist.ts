import { Episode, Playlist } from 'types/app'

export const GET_USER_PLAYLISTS_REQUEST = 'GET_USER_PLAYLISTS_REQUEST'
export const GET_USER_PLAYLISTS_SUCCESS = 'GET_USER_PLAYLISTS_SUCCESS'
export const GET_USER_PLAYLISTS_FAILURE = 'GET_USER_PLAYLISTS_FAILURE'

export const CREATE_PLAYLIST_REQUEST = 'CREATE_PLAYLIST_REQUEST'
export const CREATE_PLAYLIST_SUCCESS = 'CREATE_PLAYLIST_SUCCESS'
export const CREATE_PLAYLIST_FAILURE = 'CREATE_PLAYLIST_FAILURE'

export const RECEIVED_PLAYLISTS = 'RECEIVED_PLAYLISTS'
export const RECEIVED_PLAYLIST_EPISODES = 'RECEIVED_PLAYLIST_EPISODES'

export interface GetUserPlaylistsRequestAction {
  type: typeof GET_USER_PLAYLISTS_REQUEST
}

export interface GetUserPlaylistsSuccessAction {
  type: typeof GET_USER_PLAYLISTS_SUCCESS
}

export interface GetUserPlaylistsFailureAction {
  type: typeof GET_USER_PLAYLISTS_FAILURE
}

export interface CreatePlaylistRequestAction {
  type: typeof CREATE_PLAYLIST_REQUEST
}

export interface CreatePlaylistSuccessAction {
  type: typeof CREATE_PLAYLIST_SUCCESS
}

export interface CreatePlaylistFailureAction {
  type: typeof CREATE_PLAYLIST_FAILURE
}

export interface ReceivedPlaylistsAction {
  type: typeof RECEIVED_PLAYLISTS
  userId: string
  playlists: Playlist[]
}

export interface ReceivedPlaylistEpisodesAction {
  type: typeof RECEIVED_PLAYLIST_EPISODES
  playlistId: string
  episodes: Episode[]
}

export type PlaylistActionTypes =
  | GetUserPlaylistsRequestAction
  | GetUserPlaylistsSuccessAction
  | GetUserPlaylistsFailureAction
  | CreatePlaylistRequestAction
  | CreatePlaylistSuccessAction
  | CreatePlaylistFailureAction
  | ReceivedPlaylistsAction
  | ReceivedPlaylistEpisodesAction
