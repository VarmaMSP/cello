import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const podcasts: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.RECEIVED_SEARCH_PODCASTS:
      return action.podcasts.map((p) => p.id)
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
})
