import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { AudioState } from 'types/app'

const playingEpisodeId: Reducer<string, T.AppActions> = (
  state = '',
  action,
) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_PLAY_EPISODE:
      return action.episodeId

    default:
      return state
  }
}

const state: Reducer<AudioState, T.AppActions> = (
  state = 'LOADING',
  action,
) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_UPDATE_STATE:
      return action.state

    default:
      return state
  }
}

const duration: Reducer<number, T.AppActions> = (state = 0, action) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_UPDATE_DURATION:
      return action.duration

    default:
      return state
  }
}

const currentTime: Reducer<number, T.AppActions> = (state = 0, action) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_PLAY_EPISODE:
      return action.beginAt

    case T.AUDIO_PLAYER_UPDATE_CURRENT_TIME:
      return action.currentTime

    default:
      return state
  }
}

const volume: Reducer<number, T.AppActions> = (state = 1, action) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_UPDATE_VOLUME:
      return action.volume

    default:
      return state
  }
}

const playbackRate: Reducer<number, T.AppActions> = (state = 1, action) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_UPDATE_PLAYBACK_RATE:
      return action.playbackRate

    default:
      return state
  }
}

const expandOnMobile: Reducer<boolean, T.AppActions> = (
  state = false,
  action,
) => {
  switch (action.type) {
    case T.AUDIO_PLAYER_TOGGLE_EXPAND:
      return !state

    default:
      return state
  }
}

export default combineReducers({
  playingEpisodeId,
  state,
  duration,
  currentTime,
  volume,
  playbackRate,
  expandOnMobile,
})
