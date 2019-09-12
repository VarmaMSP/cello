import client from 'client'
import { Dispatch } from 'redux'
import {
  AppActions,
  GET_PODCAST_CURATIONS_REQUEST,
  GET_PODCAST_CURATIONS_SUCCESS,
  RECEIVED_PODCAST_CURATION,
} from 'types/actions'

export const getCurations = () => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: GET_PODCAST_CURATIONS_REQUEST })

    try {
      const { podcastCurations } = await client.getPodcastCurations()
      podcastCurations
        .filter(({ podcasts }) => !!podcasts)
        .map(({ podcasts, curation }) =>
          dispatch({ type: RECEIVED_PODCAST_CURATION, curation, podcasts }),
        )
      dispatch({ type: GET_PODCAST_CURATIONS_SUCCESS })
    } catch (err) {}
  }
}
