import * as client from 'client/home'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHomePageData() {
  return requestAction(
    () => client.getHomePageData(),
    (dispatch, _, { podcasts, categories }) => {
      dispatch({
        type: T.PODCAST_ADD,
        podcasts,
      })
      dispatch({
        type: T.CURATION_ADD,
        curations: [
          {
            id: 'recommended',
            title: 'recommended',
            members: [],
          },
        ],
        curationType: 'default',
      })
      dispatch({
        type: T.CURATION_ADD,
        curations: categories,
        curationType: 'category',
      })
      dispatch({
        type: T.CURATION_ADD_PODCASTS,
        curationId: 'recommended',
        podcastIds: podcasts.map((x) => x.id),
      })
    },
    {
      requestId: RequestId.getHomePageData(),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
