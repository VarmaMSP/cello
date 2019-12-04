import { Episode, Playback } from 'types/app'

export const RECEIVED_EPISODE = 'RECEIVED_EPISODE'
export const RECEIVED_PLAYBACKS = 'RECEIVED_PLAYBACKS'

export interface ReceivedEpisodeAction {
  type: typeof RECEIVED_EPISODE
  episode: Episode
}

export interface ReceivedPlaybacksAction {
  type: typeof RECEIVED_PLAYBACKS
  playbacks: Playback[]
}

export type EpisodeActionTypes = ReceivedEpisodeAction | ReceivedPlaybacksAction
