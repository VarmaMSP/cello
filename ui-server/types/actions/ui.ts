import { AudioState, ScreenWidth } from 'types/app'

// DOM and Window Actions
export const SET_SCREEN_WIDTH = 'SET_SCREEN_WIDTH'

export interface SetScreenWidthAction {
  type: typeof SET_SCREEN_WIDTH
  width: ScreenWidth
}

// Modal Actions
export const SHOW_SIGN_IN_MODAL = 'SHOW_SIGN_IN_MODAL'
export const CLOSE_SIGN_IN_MODAL = 'CLOSE_SIGN_IN_MODAL'

export interface ShowSignInModalAction {
  type: typeof SHOW_SIGN_IN_MODAL
}

export interface CloseSignInModalAction {
  type: typeof CLOSE_SIGN_IN_MODAL
}

// Searchbar Actions
export const SEARCH_BAR_TEXT_CHANGE = 'SEARCH_BAR_TEXT_CHANGE'

export interface SearchBarTextChangeAction {
  type: typeof SEARCH_BAR_TEXT_CHANGE
  text: string
}

// Audio Player Actions
export const PLAY_EPISODE = 'PLAY_EPISODE'
export const SET_AUDIO_STATE = 'SET_AUDIO_STATE'
export const TOGGLE_EXPAND_ON_MOBILE = 'TOGGLE_MOBILE_PLAYER' // TODO: The UI change has to be depricated

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

export type UiActionTypes =
  // Dom
  | SetScreenWidthAction
  // Modal
  | ShowSignInModalAction
  | CloseSignInModalAction
  // Searchbar
  | SearchBarTextChangeAction
  // Audio Player
  | PlayEpisodeAction
  | SetAudioStateAction
  | ToggleExpandOnMobileAction
