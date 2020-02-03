import { doFetch_ } from 'client/fetch'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getChartPageData(chartId: string) {
  return requestAction(
    () =>
      doFetch_({
        method: 'GET',
        urlPath: `/charts/${chartId}`,
      }),
    (dispatch, _, { podcasts }) => {
      dispatch({
        type: T.PODCAST_ADD,
        podcasts,
      })
      dispatch({
        type: T.CURATION_ADD_PODCASTS,
        curationId: chartId,
        podcastIds: podcasts.map((x) => x.id),
      })
    },
    {
      requestId: RequestId.getPodcastsInChart(chartId),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
