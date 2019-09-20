import { combineReducers } from 'redux'
import {
  GET_PODCAST_FAILURE,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
} from 'types/actions'
import { defaultRequestReducer } from './utils'

const getPodcast = defaultRequestReducer(
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_PODCAST_FAILURE,
)

export default combineReducers({
  getPodcast,
})
