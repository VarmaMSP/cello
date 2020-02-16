import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'

const query: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_QUERY:
      return action.query

    default:
      return state
  }
}

const resultType: Reducer<SearchResultType, T.AppActions> = (
  state = 'episode',
  action,
) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_RESULT_TYPE:
      return action.resultType

    default:
      return state
  }
}

const sortBy: Reducer<SearchSortBy, T.AppActions> = (
  state = 'relevance',
  action,
) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_SORT_BY:
      return action.sortBy

    default:
      return state
  }
}

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

const podcastsBestMatch: Reducer<
  {
    [searchQuery: string]: string[]
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULTS_LIST_LOAD_PODcAST_BEST_MATCH:
      return {
        ...state,
        [action.searchQuery]: action.podcastIds,
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
  query,
  resultType,
  sortBy,
  podcasts,
  podcastsBestMatch,
  episodes,
  receivedAll,
})
