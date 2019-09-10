import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  GET_PODCAST_FAILURE,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
} from 'types/actions'
import { initialRequestState, RequestState } from './utils'

const getPodcast: Reducer<RequestState, AppActions> = (
  state = initialRequestState(),
  action,
) => {
  switch (action.type) {
    case GET_PODCAST_REQUEST:
      return { status: 'STARTED', error: null }
    case GET_PODCAST_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case GET_PODCAST_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

export default combineReducers({
  getPodcast,
})
