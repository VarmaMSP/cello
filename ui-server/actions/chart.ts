import * as T from 'types/actions'
import { Podcast } from 'types/models'
import { doFetch } from 'utils/fetch'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getChartPageData(chartUrlParam: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/charts/${chartUrlParam}`,
      }),
    (dispatch, _, { raw, categories }) => {
      const podcasts: Podcast[] = (raw.podcasts || []).map(
        (o: any) => new Podcast(o),
      )

      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.CATEGORY_ADD, categories })
      dispatch({
        type: T.CATEGORY_ADD_PODCASTS,
        categoryId: categories[0].id,
        podcastIds: podcasts.map((x) => x.id),
      })
    },
    {
      requestId: RequestId.getPodcastsInChart(chartUrlParam),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
