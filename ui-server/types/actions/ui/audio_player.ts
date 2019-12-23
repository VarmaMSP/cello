import { AudioState } from 'types/app'

export const AUDIO_PLAYER_PLAY_EPISODE = 'audio_player/play_episode'
export const AUDIO_PLAYER_UPDATE_STATE = 'audio_player/update_state'
export const AUDIO_PLAYER_UPDATE_DURATION = 'audio_player/update_duration'
export const AUDIO_PLAYER_UPDATE_CURRENT_TIME =
  'audio_player/update_current_time'
export const AUDIO_PLAYER_UPDATE_VOLUME = 'audio_player/update_volume'
export const AUDIO_PLAYER_UPDATE_PLAYBACK_RATE =
  'audio_player/update_playback_rate'
export const AUDIO_PLAYER_TOGGLE_EXPAND = 'audio_player/update_toggle_expand'

interface PlayEpisodeAction {
  type: typeof AUDIO_PLAYER_PLAY_EPISODE
  episodeId: string
  beginAt: number
}

interface UpdateStateAction {
  type: typeof AUDIO_PLAYER_UPDATE_STATE
  state: AudioState
}

interface UpdateDurationAction {
  type: typeof AUDIO_PLAYER_UPDATE_DURATION
  duration: number
}

interface UpdateCurrentTimeAction {
  type: typeof AUDIO_PLAYER_UPDATE_CURRENT_TIME
  currentTime: number
}

interface UpdateVolumeAction {
  type: typeof AUDIO_PLAYER_UPDATE_VOLUME
  volume: number
}

interface UpdatePlaybackRateAction {
  type: typeof AUDIO_PLAYER_UPDATE_PLAYBACK_RATE
  playbackRate: number
}

interface ToggleExpandAction {
  type: typeof AUDIO_PLAYER_TOGGLE_EXPAND
}

export type AudioPlayerActionTypes =
  | PlayEpisodeAction
  | UpdateStateAction
  | UpdateDurationAction
  | UpdateCurrentTimeAction
  | UpdateVolumeAction
  | UpdatePlaybackRateAction
  | ToggleExpandAction
