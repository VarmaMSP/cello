import { AudioState } from 'types/app'

// Modal Actions
export const SHOW_SIGNIN_MODAL = 'SHOW_SIGNIN_MODAL'
export const SHOW_EPISODE_MODAL = 'SHOW_EPISODE_MODAL'
export const SHOW_ADD_TO_PLAYLIST_MODAL = 'SHOW_ADD_TO_PLAYLIST_MODAL'
export const CLOSE_MODAL = 'CLOSE_MODAL'

export interface ShowSigninModalAction {
  type: typeof SHOW_SIGNIN_MODAL
}

export interface ShowEpisodeModalAction {
  type: typeof SHOW_EPISODE_MODAL
  episodeId: string
}

export interface ShowAddToPlaylistModalAction {
  type: typeof SHOW_ADD_TO_PLAYLIST_MODAL
  episodeId: string
}

export interface CloseModalAction {
  type: typeof CLOSE_MODAL
}

// Searchbar Actions
export const SEARCH_BAR_TEXT_CHANGE = 'SEARCH_BAR_TEXT_CHANGE'

export interface SearchBarTextChangeAction {
  type: typeof SEARCH_BAR_TEXT_CHANGE
  text: string
}

// Audio Player Actions
export const PLAY_EPISODE = 'PLAY_EPISODE'
export const SET_DURATION = 'SET_DURATION'
export const SET_AUDIO_STATE = 'SET_AUDIO_STATE'
export const SET_CURRENT_TIME = 'SET_CURRENT_TIME'
export const SET_VOLUME = 'SET_VOLUME'
export const SET_PLAYBACK_RATE = 'SET_PLAYBACK_RATE'
export const TOGGLE_EXPAND_ON_MOBILE = 'TOGGLE_MOBILE_PLAYER' // TODO: The UI change has to be depricated

export interface PlayEpisodeAction {
  type: typeof PLAY_EPISODE
  episodeId: string
  currentTime: number
}

export interface SetDurationAction {
  type: typeof SET_DURATION
  duration: number
}

export interface SetAudioStateAction {
  type: typeof SET_AUDIO_STATE
  state: AudioState
}

export interface SetCurrentTimeAction {
  type: typeof SET_CURRENT_TIME
  currentTime: number
}

export interface SetVolumeAction {
  type: typeof SET_VOLUME
  volume: number
}

export interface SetPlaybackRateAction {
  type: typeof SET_PLAYBACK_RATE
  playbackRate: number
}

export interface ToggleExpandOnMobileAction {
  type: typeof TOGGLE_EXPAND_ON_MOBILE
}

export type UiActionTypes =
  // Modal
  | ShowSigninModalAction
  | ShowEpisodeModalAction
  | ShowAddToPlaylistModalAction
  | CloseModalAction
  // Searchbar
  | SearchBarTextChangeAction
  // Audio Player
  | PlayEpisodeAction
  | SetDurationAction
  | SetAudioStateAction
  | SetCurrentTimeAction
  | SetVolumeAction
  | SetPlaybackRateAction
  | ToggleExpandOnMobileAction
