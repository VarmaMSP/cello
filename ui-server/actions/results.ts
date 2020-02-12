import Router from 'next/router'
import { Dispatch } from 'redux'
import * as T from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'
import { doFetch } from 'utils/fetch'
import * as gtag from 'utils/gtag'
import * as RequestId from 'utils/request_id'
import { encodeQueryParam, qs } from 'utils/utils'
import { requestAction } from './utils'

export function loadResultsPage(
  query: string,
  resultType: SearchResultType,
  sortBy: SearchSortBy,
) {
  return (dispatch: Dispatch<T.AppActions>) => {
    gtag.search(query)

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
      `/results?query=${encodeQueryParam(
        query,
      )}&type=${resultType}&sort_by=${sortBy}`,
    )
  }
}

export function loadPodcastPage(podcastUrlParam: string) {
  return (dispatch: Dispatch<T.AppActions>) => {
    dispatch({
      type: T.HISTORY_PUSH_ENTRY,
      entry: {
        urlPath: Router.asPath,
        scrollY: window.scrollY,
      },
    })

    Router.push(
      {
        pathname: '/podcasts',
        query: { podcastUrlParam: podcastUrlParam },
      },
      `/podcasts/${podcastUrlParam}`,
    )
  }
}

export function getResultsPageData(
  query: string,
  resultType: 'podcast' | 'episode',
  sortBy: 'relevance' | 'publish_date',
) {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/results?${qs({
          query: query,
          type: resultType,
          sort_by: sortBy,
        })}`,
      }),
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
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/ajax/browse?${qs({
          endpoint: 'search_results',
          query: query,
          type: resultType,
          sort_by: sortBy,
          offset: offset,
          limit: limit,
        })}`,
      }),
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
