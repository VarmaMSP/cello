import { Episode, Playlist } from 'types/app'

export const RECEIVED_PLAYLISTS = 'RECEIVED_PLAYLISTS'
export const RECEIVED_PLAYLIST_EPISODES = 'RECEIVED_PLAYLIST_EPISODES'

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
  | ReceivedPlaylistsAction
  | ReceivedPlaylistEpisodesAction
