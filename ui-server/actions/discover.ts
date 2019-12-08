import * as client from 'client/discover'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getDiscoverPageData() {
  return requestAction(
    () => client.getDiscoverPageData(),
    (dispatch, _, { podcasts, categories }) => {
      dispatch({
        type: T.RECEIVED_PODCAST_CATEGORY_LISTS,
        categories,
      }),
        dispatch({
          type: T.RECEIVED_RECOMMENDED_PODCASTS,
          podcasts,
        })
    },
    {
      requestId: RequestId.getDiscoverPageData(),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}

export function getPodcastsInList(listId: string) {
  return requestAction(
    () => client.getPodcastsFromList(listId),
    (dispatch, _, { podcasts }) => {
      dispatch({
        type: T.RECEIVED_PODCASTS_IN_LIST,
        listId,
        podcasts,
      })
    },
    {
      requestId: RequestId.getPodcastsInList(listId),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
