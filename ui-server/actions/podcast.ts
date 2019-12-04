import * as client from 'client/podcast'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getPodcast(podcastId: string) {
  return requestAction(
    () => client.getPodcast(podcastId),
    (dispatch, _, { podcast, episodes }) => {
      dispatch({ type: T.RECEIVED_PODCAST, podcast })
      dispatch({
        type: T.RECEIVED_PODCAST_EPISODES,
        podcastId: podcast.id,
        offset: 0,
        order: 'pub_date_desc',
        episodes,
      })
    },
  )
}

export function getTrendingPodcasts() {
  return requestAction(
    () => Promise.resolve(),
    (dispatch) => {
      dispatch({ type: T.RECEIVED_TRENDING_PODCASTS, podcasts: [] })
    },
  )
}
