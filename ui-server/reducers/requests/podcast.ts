import {
  AppActions,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_PODCAST_FAILURE,
} from '../../types/actions'
import { combineReducers, Reducer } from 'redux'
import { RequestState, initalRequestState } from './utils'

const getPodcast: Reducer<RequestState, AppActions> = (
  state = initalRequestState(),
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
