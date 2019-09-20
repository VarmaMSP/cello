import { combineReducers } from 'redux'
import {
  GET_PODCAST_CURATIONS_FAILURE,
  GET_PODCAST_CURATIONS_REQUEST,
  GET_PODCAST_CURATIONS_SUCCESS,
} from 'types/actions'
import { defaultRequestReducer } from './utils'

const getAllCurations = defaultRequestReducer(
  GET_PODCAST_CURATIONS_REQUEST,
  GET_PODCAST_CURATIONS_SUCCESS,
  GET_PODCAST_CURATIONS_FAILURE,
)

export default combineReducers({
  getAllCurations,
})
