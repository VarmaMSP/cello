import { AppActions, RECEIVED_SEARCH_PODCASTS } from '../../types/actions'
import { Reducer, combineReducers } from 'redux'

const podcasts: Reducer<string[], AppActions> = (state = [], action) => {
  switch (action.type) {
    case RECEIVED_SEARCH_PODCASTS:
      return action.podcasts.map((p) => p.id)
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
})
