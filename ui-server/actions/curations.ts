import client from 'client'
import {
  GET_PODCAST_CURATIONS_REQUEST,
  GET_PODCAST_CURATIONS_SUCCESS,
  RECEIVED_PODCAST_CURATION,
} from 'types/actions'
import { requestAction } from './utils'

export function getCurations() {
  return requestAction(
    () => client.getPodcastCurations(),
    (dispatch, { podcastCurations }) => {
      podcastCurations
        .filter(({ podcasts }) => !!podcasts)
        .map(({ podcasts, curation }) =>
          dispatch({ type: RECEIVED_PODCAST_CURATION, curation, podcasts }),
        )
    },
    { type: GET_PODCAST_CURATIONS_REQUEST },
    { type: GET_PODCAST_CURATIONS_SUCCESS },
    { type: GET_PODCAST_CURATIONS_SUCCESS },
  )
}
