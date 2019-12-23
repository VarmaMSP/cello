import { Playlist } from 'types/app'

export const PLAYLIST_ADD = `playlist/add`
export const PLAYLIST_ADD_EPISODES = 'playlist/add_episodes'
export const PLAYLIST_REMOVE_EPISODES = 'playlist/remove_episodes'

interface AddAction {
  type: typeof PLAYLIST_ADD
  playlists: Playlist[]
}

interface AddEpisodesAction {
  type: typeof PLAYLIST_ADD_EPISODES
  playlistId: string
  episodeIds: string[]
}

interface RemoveEpisodesAction {
  type: typeof PLAYLIST_REMOVE_EPISODES
  playlistId: string
  episodeIds: string[]
}

export type PlaylistActionTypes =
  | AddAction
  | AddEpisodesAction
  | RemoveEpisodesAction
