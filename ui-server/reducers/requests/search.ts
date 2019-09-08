import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  SEARCH_PODCASTS_FAILURE,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
} from 'types/actions'
import { initalRequestState, RequestState } from './utils'

const searchPodcasts: Reducer<RequestState, AppActions> = (
  state = initalRequestState(),
  action,
) => {
  switch (action.type) {
    case SEARCH_PODCASTS_REQUEST:
      return { status: 'STARTED', error: null }
    case SEARCH_PODCASTS_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case SEARCH_PODCASTS_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

export default combineReducers({
  searchPodcasts,
})
