import * as client from 'client/search'
import Router from 'next/router'
import { Dispatch } from 'redux'
import { getText } from 'selectors/ui/search_bar'
import { getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function loadResultsPage() {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    const state = getState()
    const query = getText(state)
    const resultType = getResultType(state)
    const sortBy = getSortBy(state)

    dispatch({
      type: T.HISTORY_PUSH_ENTRY,
      entry: {
        urlPath: Router.asPath,
        scrollY: window.scrollY,
      },
    })

    Router.push(
      {
        pathname: '/results',
        query: { query, resultType, sortBy },
      },
      `/results?query=${query}&type=${resultType}&sort_by=${sortBy}`,
    )
  }
}

export function getResultsPageData(
  query: string,
  resultType: 'podcast' | 'episode',
  sortBy: 'relevance' | 'publish_date',
) {
  return requestAction(
    () => client.getResultsPageData(query, resultType, sortBy),
    (
      dispatch,
      _,
      { podcasts, episodes, podcastSearchResults, episodeSearchResults },
    ) => {
      if (resultType === 'podcast') {
        dispatch({ type: T.PODCAST_ADD, podcasts })

        dispatch({
          type: T.SEARCH_RESULT_ADD_PODCAST,
          podcastSearchResults,
          searchQuery: query,
        })

        dispatch({
          type: T.SEARCH_RESULTS_LIST_LOAD_PODCAST_PAGE,
          searchQuery: query,
          sortBy: sortBy,
          page: 0,
          podcastIds: podcastSearchResults.map((x) => x.id),
        })

        if (podcasts.length < 25) {
          dispatch({
            type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
            searchQuery: query,
            resultType: resultType,
            sortBy: sortBy,
          })
        }
      }

      if (resultType === 'episode') {
        dispatch({ type: T.PODCAST_ADD, podcasts })

        dispatch({ type: T.EPISODE_ADD, episodes })

        dispatch({
          type: T.SEARCH_RESULT_ADD_EPISODE,
          episodeSearchResults,
          searchQuery: query,
        })

        dispatch({
          type: T.SEARCH_RESULTS_LIST_LOAD_EPISODE_PAGE,
          searchQuery: query,
          sortBy: sortBy,
          page: 0,
          episodeIds: episodeSearchResults.map((x) => x.id),
        })

        if (episodes.length < 25) {
          dispatch({
            type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
            searchQuery: query,
            resultType: resultType,
            sortBy: sortBy,
          })
        }
      }
    },
    { requestId: RequestId.getResults() },
  )
}

export function getResults(
  query: string,
  resultType: SearchResultType,
  sortBy: SearchSortBy,
  offset: number,
  limit: number,
) {
  return requestAction(
    () => client.getResults(query, resultType, sortBy, offset, limit),
    (
      dispatch,
      _,
      { podcasts, episodes, podcastSearchResults, episodeSearchResults },
    ) => {
      if (resultType === 'podcast') {
        dispatch({ type: T.PODCAST_ADD, podcasts })

        dispatch({
          type: T.SEARCH_RESULT_ADD_PODCAST,
          podcastSearchResults,
          searchQuery: query,
        })

        dispatch({
          type: T.SEARCH_RESULTS_LIST_LOAD_PODCAST_PAGE,
          searchQuery: query,
          sortBy: sortBy,
          page: Math.floor(offset / limit),
          podcastIds: podcastSearchResults.map((x) => x.id),
        })

        if (podcasts.length < limit) {
          dispatch({
            type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
            searchQuery: query,
            resultType: resultType,
            sortBy: sortBy,
          })
        }
      }

      if (resultType === 'episode') {
        dispatch({ type: T.PODCAST_ADD, podcasts })

        dispatch({ type: T.EPISODE_ADD, episodes })

        dispatch({
          type: T.SEARCH_RESULT_ADD_EPISODE,
          episodeSearchResults,
          searchQuery: query,
        })

        dispatch({
          type: T.SEARCH_RESULTS_LIST_LOAD_EPISODE_PAGE,
          searchQuery: query,
          sortBy: sortBy,
          page: Math.floor(offset / limit),
          episodeIds: episodeSearchResults.map((x) => x.id),
        })

        if (episodes.length < limit) {
          dispatch({
            type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
            searchQuery: query,
            resultType: resultType,
            sortBy: sortBy,
          })
        }
      }
    },
    { requestId: RequestId.getResults() },
  )
}
