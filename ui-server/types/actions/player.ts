import { AudioState } from 'types/app'

export const PLAY_EPISODE = 'PLAY_EPISODE'
export const SET_AUDIO_STATE = 'SET_AUDIO_STATE'
export const TOGGLE_EXPAND_ON_MOBILE = 'TOGGLE_MOBILE_PLAYER'

export interface PlayEpisodeAction {
  type: typeof PLAY_EPISODE
  episodeId: string
}

export interface SetAudioStateAction {
  type: typeof SET_AUDIO_STATE
  state: AudioState
}

export interface ToggleExpandOnMobileAction {
  type: typeof TOGGLE_EXPAND_ON_MOBILE
}

export type PlayerActionTypes =
  | PlayEpisodeAction
  | SetAudioStateAction
  | ToggleExpandOnMobileAction
