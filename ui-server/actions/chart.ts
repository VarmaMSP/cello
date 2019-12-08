import * as client from 'client/chart'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPodcastsInChart(chartId: string) {
  return requestAction(
    () => client.getPodcastsInChart(chartId),
    (dispatch, _, { podcasts }) => {
      dispatch({
        type: T.RECEIVED_CHART_PODCASTS,
        chartId,
        podcasts,
      })
    },
    {
      requestId: RequestId.getPodcastsInChart(chartId),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
