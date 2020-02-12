import { Episode, Playback } from 'models'

export const EPISODE_ADD = 'episode/add'
export const EPISODE_JOIN_PLAYBACK = 'episode/join_playbacks'

interface AddAction {
  type: typeof EPISODE_ADD
  episodes: Episode[]
}

interface JoinPlaybacksAction {
  type: typeof EPISODE_JOIN_PLAYBACK
  playbacks: Playback[]
}

export type EpisodeActionTypes =
  | AddAction
  | JoinPlaybacksAction
