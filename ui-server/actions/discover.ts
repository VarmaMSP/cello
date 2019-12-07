import * as client from 'client/discover'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getDiscoverPageData() {
  return requestAction(
    () => client.getDiscoverPageData(),
    (dispatch, _, { podcasts, categories }) => {
      categories.forEach((category) =>
        dispatch({
          type: T.RECEIVED_PODCAST_CATEGORY_LIST,
          category,
        }),
      )
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
