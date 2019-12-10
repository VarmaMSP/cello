import { Episode, Playlist } from 'types/app'

export const RECEIVED_PLAYLIST = 'RECEIVED_PLAYLIST'
export const RECEIVED_PLAYLISTS = 'RECEIVED_PLAYLISTS'
export const RECEIVED_PLAYLIST_EPISODES = 'RECEIVED_PLAYLIST_EPISODES'
export const ADD_EPISODE_TO_PLAYLIST = 'ADD_EPISODE_TO_PLAYLISTS'
export const REMOVE_EPISODE_FROM_PLAYLIST = 'REMOVE_EPISODE_FROM_PlAYLIST'

export interface ReceivedPlaylistAction {
  type: typeof RECEIVED_PLAYLIST
  playlist: Playlist
}

export interface ReceivedPlaylistsAction {
  type: typeof RECEIVED_PLAYLISTS
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
  | ReceivedPlaylistsAction
  | ReceivedPlaylistEpisodesAction
  | AddEpisodeToPlaylistAction
  | RemoveEpisodeFromPlaylistAction
