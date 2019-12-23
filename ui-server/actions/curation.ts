import * as client from 'client/chart'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getCurationPageData(chartId: string) {
  return requestAction(
    () => client.getChartPageData(chartId),
    (dispatch, _, { podcasts }) => {
      dispatch({
        type: T.PODCAST_ADD,
        podcasts,
      })
      dispatch({
        type: T.CURATION_ADD_PODCASTS,
        chartId,
        podcastIds: podcasts.map((x) => x.id),
      })
    },
    {
      requestId: RequestId.getPodcastsInChart(chartId),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
