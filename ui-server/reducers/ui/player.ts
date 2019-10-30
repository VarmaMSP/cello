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

const duration: Reducer<number, T.AppActions> = (state = 0, action) => {
  switch (action.type) {
    case T.SET_DURATION:
      return action.duration
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

const currentTime: Reducer<number, T.AppActions> = (state = 0, action) => {
  switch (action.type) {
    case T.PLAY_EPISODE:
    case T.SET_CURRENT_TIME:
      return action.currentTime
    default:
      return state
  }
}

const volume: Reducer<number, T.AppActions> = (state = 1, action) => {
  switch (action.type) {
    case T.SET_VOLUME:
      return action.volume
    default:
      return state
  }
}

const playbackSpeed: Reducer<number, T.AppActions> = (state = 1, action) {
  switch (action.type) {
    case T.SET_PLAYBACK_SPEED: 
      return action.playbackSpeed
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
  duration,
  audioState,
  currentTime,
  volume,
  playbackSpeed,
  expandOnMobile,
})
