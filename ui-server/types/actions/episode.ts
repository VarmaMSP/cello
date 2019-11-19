import { EpisodePlayback } from 'types/app'

export const RECEIVED_HISTORY_PLAYBACKS = 'RECEIVED_HISTORY_PLAYBACKS'
export const RECEIVED_EPISODE_PLAYBACKS = 'RECEIVED_EPISODE_PLAYBACKS'

export interface ReceivedHistoryPlaybacksAction {
  type: typeof RECEIVED_HISTORY_PLAYBACKS
  playbacks: EpisodePlayback[]
}

export interface ReceivedEpisodePlaybacksAction {
  type: typeof RECEIVED_EPISODE_PLAYBACKS
  playbacks: EpisodePlayback[]
}

export type EpisodeActionTypes =
  | ReceivedHistoryPlaybacksAction
  | ReceivedEpisodePlaybacksAction
