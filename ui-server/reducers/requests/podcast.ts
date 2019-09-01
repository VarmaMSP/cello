import {
  AppActions,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_PODCAST_FAILURE,
} from 'types/actions'
import { combineReducers } from 'redux'

interface RequestState {
  status: string
  error: Error | null
}

const getPodcastReducer = (
  state: RequestState = {
    status: 'not_started',
    error: null,
  },
  action: AppActions,
): RequestState => {
  switch (action.type) {
    case GET_PODCAST_REQUEST:
      return { status: 'PENDING', error: null }
    case GET_PODCAST_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case GET_PODCAST_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

export default combineReducers({
  getPodcast: getPodcastReducer,
})
