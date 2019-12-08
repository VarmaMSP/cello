import * as client from 'client/home'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHomePageData() {
  return requestAction(
    () => client.getHomePageData(),
    (dispatch, _, { podcasts, categories }) => {
      dispatch({
        type: T.RECEIVED_PODCAST_CHARTS,
        charts: [...categories, { id: 'recommended', title: 'recommended', type: 'NORMAL' }],
      })
      dispatch({
        type: T.RECEIVED_CHART_PODCASTS,
        chartId: 'recommended',
        podcasts,
      })
    },
    {
      requestId: RequestId.getHomePageData(),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
