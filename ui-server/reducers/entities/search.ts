import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const podcastResults: Reducer<{ [query: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_SEARCH_PODCASTS:
      return {
        ...state,
        [action.query]: action.podcasts.map((p) => p.id),
      }
    default:
      return state
  }
}

export default combineReducers({
  podcastResults,
})
