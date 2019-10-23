import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const getPlaybackHistory = defaultRequestReducer(
  T.GET_PLAYBACK_HISTORY_REQUEST,
  T.GET_PLAYBACK_HISTORY_SUCCESS,
  T.GET_PLAYBACK_HISTORY_FAILURE,
)

export default combineReducers({
  getPlaybackHistory,
})
