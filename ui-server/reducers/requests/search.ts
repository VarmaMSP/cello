import {
  AppActions,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
  SEARCH_PODCASTS_FAILURE,
} from '../../types/actions'
import { combineReducers, Reducer } from 'redux'
import { RequestState, initalRequestState } from './utils'

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
