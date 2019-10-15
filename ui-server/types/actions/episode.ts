export const SYNC_PLAYBACK_REQUEST = 'SYNC_PLAYBACK_REQUEST'
export const SYNC_PLAYBACK_SUCCESS = 'SYNC_PLAYBACK_SUCCESS'
export const SYNC_PLAYBACK_FAILURE = 'SYNC_PLAYBACK_FAILURE'

export const SYNC_PLAYBACK_PROGRESS_REQUEST = 'SYNC_PLAYBACK_PROGRESS_REQUEST'

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

export type EpisodeActionTypes =
  | SyncPlaybackRequestAction
  | SyncPlaybackSuccessAction
  | SyncPlaybackFailureAction
  | SyncPlaybackProgressRequestAction
