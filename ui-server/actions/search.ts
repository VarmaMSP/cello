import * as client from 'client/search'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function searchPodcasts(query: string) {
  return requestAction(
    () => client.searchPodcasts(query),
    (dispatch, _, { podcasts }) => {
      dispatch({ type: T.RECEIVED_SEARCH_PODCASTS, query, podcasts })
    },
  )
}
