import * as client from 'client/search'
import * as T from 'types/actions'
import { requestAction } from './utils'

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
          podcastIds: podcasts.map((x) => x.id),
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
          episodeIds: episodes.map((x) => x.id),
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
  )
}

// export function getResults(query: string, offset: number, limit: number) {
//   return requestAction(
//     () => client.getResults(query, offset, limit),
//     (dispatch, _, { podcasts, podcastSearchResults }) => {
//       dispatch({ type: T.PODCAST_ADD, podcasts })

//       dispatch({
//         type: T.SEARCH_RESULT_ADD_PODCAST,
//         podcastSearchResults,
//         searchQuery: query,
//       })

//       dispatch({
//         type: T.SEARCH_RESULTS_LIST_LOAD_PAGE,
//         searchQuery: query,
//         page: Math.floor(offset / 10),
//         podcastIds: podcasts.map((x) => x.id),
//       })

//       if (podcasts.length < limit) {
//         dispatch({
//           type: T.SEARCH_RESULTS_LIST_RECEIVED_ALL,
//           searchQuery: query,
//         })
//       }
//     },
//   )
// }
