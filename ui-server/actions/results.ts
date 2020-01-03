import * as client from 'client/search'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getResultsPageData(query: string) {
  return requestAction(
    () => client.getResultsPageData(query),
    (dispatch, _, { podcasts, podcastSearchResults }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })

      dispatch({
        type: T.SEARCH_RESULT_ADD_PODCAST,
        podcastSearchResults,
        searchQuery: query,
      })

      dispatch({
        type: T.SEARCH_RESULTS_LIST_LOAD_PAGE,
        searchQuery: query,
        page: 0,
        podcastIds: podcasts.map((x) => x.id),
      })

      if (podcasts.length < 25) {
        dispatch({
          type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
          searchQuery: query,
        })
      }
    },
  )
}

export function getResults(query: string, offset: number, limit: number) {
  return requestAction(
    () => client.getResults(query, offset, limit),
    (dispatch, _, { podcasts, podcastSearchResults }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })

      dispatch({
        type: T.SEARCH_RESULT_ADD_PODCAST,
        podcastSearchResults,
        searchQuery: query,
      })

      dispatch({
        type: T.SEARCH_RESULTS_LIST_LOAD_PAGE,
        searchQuery: query,
        page: Math.floor(offset / 10),
        podcastIds: podcasts.map((x) => x.id),
      })

      if (podcasts.length < limit) {
        dispatch({
          type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
          searchQuery: query,
        })
      }
    },
  )
}
