import client from 'client'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getCurations() {
  return requestAction(
    () => client.getPodcastCurations(),
    (dispatch, { podcastCurations }) => {
      podcastCurations
        .filter(({ podcasts }) => !!podcasts)
        .map(({ podcasts, curation }) =>
          dispatch({ type: T.RECEIVED_PODCAST_CURATION, curation, podcasts }),
        )
    },
    { type: T.GET_PODCAST_CURATIONS_REQUEST },
    { type: T.GET_PODCAST_CURATIONS_SUCCESS },
    { type: T.GET_PODCAST_CURATIONS_SUCCESS },
  )
}
