import * as T from 'types/actions'
import { Curation, Podcast } from 'models'
import { doFetch } from 'utils/fetch'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHomePageData() {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: '/',
      }),
    (dispatch, _, { raw, categories }) => {
      const podcasts: Podcast[] = (raw.recommended || []).map(
        (o: any) => new Podcast(o),
      )
      const curations: Curation[] = [
        {
          id: 'recommended',
          title: 'recommended',
          podcastIds: [],
        },
      ]

      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.CATEGORY_ADD, categories })
      dispatch({ type: T.CURATION_ADD, curations })
      dispatch({
        type: T.CURATION_ADD_PODCASTS,
        curationId: curations[0].id,
        podcastIds: podcasts.map((x) => x.id),
      })
    },
    {
      requestId: RequestId.getHomePageData(),
      skip: { cond: 'REQUEST_ALREADY_MADE' },
    },
  )
}
