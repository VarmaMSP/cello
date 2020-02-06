import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PodcastSearchResult } from 'types/models'

const podcasts: Reducer<PodcastSearchResult[], T.AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    case T.SEARCH_SUGGESTIONS_ADD_PODCAST:
      return action.podcasts

    default:
      return state
  }
}

export default combineReducers({
  podcasts,
})
