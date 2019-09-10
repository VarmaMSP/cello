import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  GET_PODCAST_CURATIONS_FAILURE,
  GET_PODCAST_CURATIONS_REQUEST,
  GET_PODCAST_CURATIONS_SUCCESS,
} from 'types/actions'
import { initialRequestState, RequestState } from './utils'

const getAllCurations: Reducer<RequestState, AppActions> = (
  state = initialRequestState(),
  action,
) => {
  switch (action.type) {
    case GET_PODCAST_CURATIONS_REQUEST:
      return { status: 'STARTED', error: null }
    case GET_PODCAST_CURATIONS_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case GET_PODCAST_CURATIONS_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

export default combineReducers({
  getAllCurations,
})
