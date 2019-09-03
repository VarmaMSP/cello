import { Reducer } from 'react'
import {
  AppActions,
  SET_AUDIO_STATE,
  SET_AUDIO_DURATION,
  SET_AUDIO_CURRENT_TIME,
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

const podcast: Reducer<string | undefined, AppActions> = (
  state = '',
  action,
) => {
  switch (action.type) {
    case PLAY_EPISODE:
      return action.podcastId
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

const audioDuration: Reducer<number | undefined, AppActions> = (
  state = 0,
  action,
) => {
  switch (action.type) {
    case SET_AUDIO_DURATION:
      return action.duration
    default:
      return state
  }
}

const audioCurrentTime: Reducer<number | undefined, AppActions> = (
  state = 0,
  action,
) => {
  switch (action.type) {
    case SET_AUDIO_CURRENT_TIME:
      return action.time
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
  present: combineReducers({
    podcast,
    episode,
  }),
  audioState,
  audioDuration,
  audioCurrentTime,
  expandOnMobile,
})
