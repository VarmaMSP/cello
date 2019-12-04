import { Episode, Playlist } from 'types/app'

export const RECEIVED_PLAYLIST = 'RECEIVED_PLAYLIST'
export const RECEIVED_USER_PLAYLISTS = 'RECEIVED_USER_PLAYLIST'
export const RECEIVED_PLAYLIST_EPISODES = 'RECEIVED_PLAYLIST_EPISODES'
export const ADD_EPISODE_TO_PLAYLIST = 'ADD_EPISODE_TO_PLAYLIST'
export const REMOVE_EPISODE_FROM_PLAYLIST = 'REMOVE_EPISODE_FROM_PlAYLIST'

export interface ReceivedPlaylistAction {
  type: typeof RECEIVED_PLAYLIST
  playlist: Playlist
}

export interface ReceivedUserPlaylistsAction {
  type: typeof RECEIVED_USER_PLAYLISTS
  userId: string
  playlists: Playlist[]
}

export interface ReceivedPlaylistEpisodesAction {
  type: typeof RECEIVED_PLAYLIST_EPISODES
  playlistId: string
  episodes: Episode[]
}

export interface AddEpisodeToPlaylistAction {
  type: typeof ADD_EPISODE_TO_PLAYLIST
  episodeId: string
  playlistId: string
}

export interface RemoveEpisodeFromPlaylistAction {
  type: typeof REMOVE_EPISODE_FROM_PLAYLIST
  episodeId: string
  playlistId: string
}

export type PlaylistActionTypes =
  | ReceivedPlaylistAction
  | ReceivedUserPlaylistsAction
  | ReceivedPlaylistEpisodesAction
  | AddEpisodeToPlaylistAction
  | RemoveEpisodeFromPlaylistAction
