import { Episode, EpisodePlayback } from 'types/app'

export const RECEIVED_EPISODE = 'RECEIVED_EPISODE'
export const RECEIVED_EPISODE_PLAYBACKS = 'RECEIVED_EPISODE_PLAYBACKS'

export interface ReceivedEpisodeAction {
  type: typeof RECEIVED_EPISODE
  episode: Episode
}

export interface ReceivedEpisodePlaybacksAction {
  type: typeof RECEIVED_EPISODE_PLAYBACKS
  playbacks: EpisodePlayback[]
}

export type EpisodeActionTypes =
  | ReceivedEpisodeAction
  | ReceivedEpisodePlaybacksAction
