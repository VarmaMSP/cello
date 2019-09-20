import { combineReducers } from 'redux'
import {
  SEARCH_PODCASTS_FAILURE,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
} from 'types/actions'
import { defaultRequestReducer } from './utils'

const searchPodcasts = defaultRequestReducer(
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
  SEARCH_PODCASTS_FAILURE,
)

export default combineReducers({
  searchPodcasts,
})
