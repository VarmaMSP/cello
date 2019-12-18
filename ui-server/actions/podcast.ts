import * as client from 'client/podcast'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPodcastPageData(podcastUrlParam: string) {
  return requestAction(
    () => client.getPodcastPageData(podcastUrlParam),
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

export function getPodcastEpisodes(
  podcastId: string,
  limit: number,
  offset: number,
  order: 'pub_date_desc' | 'pub_date_asc',
) {
  return requestAction(
    () => client.getPodcastEpisodes(podcastId, limit, offset, order),
    (dispatch, _, { episodes }) => {
      dispatch({
        type: T.RECEIVED_PODCAST_EPISODES,
        podcastId,
        order,
        offset,
        episodes,
      })

      if (episodes.length < limit) {
        dispatch({
          type: T.RECEIVED_ALL_PODCAST_EPISODES,
          podcastId,
          order,
        })
      }
    },
    { requestId: RequestId.getPodcastEpisodes(podcastId) },
  )
}
