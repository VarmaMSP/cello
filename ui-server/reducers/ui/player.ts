import { Reducer } from 'react'
import {
  AppActions,
  SET_AUDIO_STATE,
  TOGGLE_EXPAND_ON_MOBILE,
  PLAY_EPISODE,
} from '../../types/actions'
import { combineReducers } from 'redux'
import { AudioState } from 'types/app'

const episode: Reducer<string | undefined, AppActions> = (
  state = '',
  action,
) => {
  switch (action.type) {
    case PLAY_EPISODE:
      return action.episodeId
    default:
      return state
  }
}

const audioState: Reducer<AudioState | undefined, AppActions> = (
  state = 'LOADING',
  action,
) => {
  switch (action.type) {
    case SET_AUDIO_STATE:
      return action.state
    default:
      return state
  }
}

const expandOnMobile: Reducer<boolean | undefined, AppActions> = (
  state = false,
  action,
) => {
  switch (action.type) {
    case TOGGLE_EXPAND_ON_MOBILE:
      return !state
    default:
      return state
  }
}

export default combineReducers({
  episode,
  audioState,
  expandOnMobile,
})
