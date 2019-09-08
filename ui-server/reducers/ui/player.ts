import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  PLAY_EPISODE,
  SET_AUDIO_STATE,
  TOGGLE_EXPAND_ON_MOBILE,
} from 'types/actions'
import { AudioState } from 'types/app'

const episode: Reducer<string, AppActions> = (state = '', action) => {
  switch (action.type) {
    case PLAY_EPISODE:
      return action.episodeId
    default:
      return state
  }
}

const audioState: Reducer<AudioState, AppActions> = (
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

const expandOnMobile: Reducer<boolean, AppActions> = (
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
