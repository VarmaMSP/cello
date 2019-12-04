import * as client from 'client/episode'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getEpisode(episodeId: string) {
  return requestAction(
    () => client.getEpisode(episodeId),
    (dispatch, _, { podcast, episode }) => {
      dispatch({ type: T.RECEIVED_EPISODE, episode: episode })
      dispatch({
        type: T.RECEIVED_PODCAST,
        podcast,
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
