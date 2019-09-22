import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const searchPodcasts = defaultRequestReducer(
  T.SEARCH_PODCASTS_REQUEST,
  T.SEARCH_PODCASTS_SUCCESS,
  T.SEARCH_PODCASTS_FAILURE,
)

export default combineReducers({
  searchPodcasts,
})
