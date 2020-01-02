import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const podcasts: Reducer<
  {
    [searchQuery: string]: {
      [page: number]: string[]
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_LOAD_PAGE:
      return {
        ...state,
        [action.searchQuery]: {
          ...(state[action.searchQuery] || {}),
          [action.page]: action.podcastIds,
        },
      }

    default:
      return state
  }
}

const receivedAll: Reducer<
  {
    [searchQuery: string]: string[]
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_RECEIVED_ALL:
      return {
        ...state,
        [action.searchQuery]: [...(state[action.searchQuery] || []), 'default'],
      }

    default:
      return state
  }
}

export default combineReducers({
  podcasts,
  receivedAll,
})
