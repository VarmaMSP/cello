import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const podcasts: Reducer<
  {
    [searchQuery: string]: {
      [sortBy: string]: {
        [page: number]: string[]
      }
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_LOAD_PODCAST_PAGE:
      return {
        ...state,
        [action.searchQuery]: {
          ...(state[action.searchQuery] || {}),
          [action.sortBy]: {
            ...((state[action.searchQuery] || {})[action.sortBy] || {}),
            [action.page]: action.podcastIds,
          },
        },
      }

    default:
      return state
  }
}

const episodes: Reducer<
  {
    [searchQuery: string]: {
      [sortBy: string]: {
        [page: number]: string[]
      }
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_LOAD_EPISODE_PAGE:
      return {
        ...state,
        [action.searchQuery]: {
          ...(state[action.searchQuery] || {}),
          [action.sortBy]: {
            ...((state[action.searchQuery] || {})[action.sortBy] || {}),
            [action.page]: action.episodeIds,
          },
        },
      }

    default:
      return state
  }
}

const receivedAll: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_RECEIVED_ALL:
      return [
        ...state,
        `${action.searchQuery}:${action.resultType}:${action.sortBy}`,
      ]

    default:
      return state
  }
}

export default combineReducers({
  podcasts,
  episodes,
  receivedAll,
})
