import { EpisodePlayback } from 'types/app'

export const GET_PLAYBACK_HISTORY_REQUEST = 'GET_PLAYBACK_HISTORY_REQUEST'
export const GET_PLAYBACK_HISTORY_SUCCESS = 'GET_PLAYBACK_HISTORY_SUCCESS'
export const GET_PLAYBACK_HISTORY_FAILURE = 'GET_PLAYBACK_HISTORY_FAILURE'

export const SYNC_PLAYBACK_REQUEST = 'SYNC_PLAYBACK_REQUEST'
export const SYNC_PLAYBACK_SUCCESS = 'SYNC_PLAYBACK_SUCCESS'
export const SYNC_PLAYBACK_FAILURE = 'SYNC_PLAYBACK_FAILURE'

export const SYNC_PLAYBACK_PROGRESS_REQUEST = 'SYNC_PLAYBACK_PROGRESS_REQUEST'

export const RECEIVED_HISTORY_PLAYBACKS = 'RECEIVED_HISTORY_PLAYBACKS'
export const RECEIVED_EPISODE_PLAYBACKS = 'RECEIVED_EPISODE_PLAYBACKS'

export interface GetPlaybackHistoryRequestAction {
  type: typeof GET_PLAYBACK_HISTORY_REQUEST
}

export interface GetPlaybackHistorySuccessAction {
  type: typeof GET_PLAYBACK_HISTORY_SUCCESS
}

export interface GetPlaybackHistoryFailureAction {
  type: typeof GET_PLAYBACK_HISTORY_FAILURE
}

export interface SyncPlaybackRequestAction {
  type: typeof SYNC_PLAYBACK_REQUEST
}

export interface SyncPlaybackSuccessAction {
  type: typeof SYNC_PLAYBACK_SUCCESS
}

export interface SyncPlaybackFailureAction {
  type: typeof SYNC_PLAYBACK_FAILURE
}

export interface SyncPlaybackProgressRequestAction {
  type: typeof SYNC_PLAYBACK_PROGRESS_REQUEST
}

export interface ReceivedHistoryPlaybacksAction {
  type: typeof RECEIVED_HISTORY_PLAYBACKS
  playbacks: EpisodePlayback[]
}

export interface ReceivedEpisodePlaybacksAction {
  type: typeof RECEIVED_EPISODE_PLAYBACKS
  playbacks: EpisodePlayback[]
}

export type EpisodeActionTypes =
  | GetPlaybackHistoryRequestAction
  | GetPlaybackHistorySuccessAction
  | GetPlaybackHistoryFailureAction
  | SyncPlaybackRequestAction
  | SyncPlaybackSuccessAction
  | SyncPlaybackFailureAction
  | SyncPlaybackProgressRequestAction
  | ReceivedHistoryPlaybacksAction
  | ReceivedEpisodePlaybacksAction
