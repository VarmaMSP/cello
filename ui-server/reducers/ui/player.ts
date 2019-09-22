import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { AudioState } from 'types/app'

const episode: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.PLAY_EPISODE:
      return action.episodeId
    default:
      return state
  }
}

const audioState: Reducer<AudioState, T.AppActions> = (
  state = 'LOADING',
  action,
) => {
  switch (action.type) {
    case T.SET_AUDIO_STATE:
      return action.state
    default:
      return state
  }
}

const expandOnMobile: Reducer<boolean, T.AppActions> = (
  state = false,
  action,
) => {
  switch (action.type) {
    case T.TOGGLE_EXPAND_ON_MOBILE:
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
