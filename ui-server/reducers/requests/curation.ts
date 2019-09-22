import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const getAllCurations = defaultRequestReducer(
  T.GET_PODCAST_CURATIONS_REQUEST,
  T.GET_PODCAST_CURATIONS_SUCCESS,
  T.GET_PODCAST_CURATIONS_FAILURE,
)

export default combineReducers({
  getAllCurations,
})
