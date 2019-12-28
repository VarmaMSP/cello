import * as client from 'client/podcast'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPodcastPageData(podcastUrlParam: string) {
  return requestAction(
    () => client.getPodcastPageData(podcastUrlParam),
    (dispatch, _, { podcast, episodes }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts: [podcast] })
      dispatch({ type: T.EPISODE_ADD, episodes })
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
        type: T.PODCAST_EPISODES_LIST_LOAD_PAGE,
        podcastId: podcastId,
        episodeIds: episodes.map((x) => x.id),
        order: order,
        page: Math.floor(offset / 10),
      })

      if (episodes.length < limit) {
        dispatch({
          type: T.PODCAST_EPISODES_LIST_RECEIVED_ALL,
          podcastId,
          order,
        })
      }
    },
    { requestId: RequestId.getPodcastEpisodes(podcastId) },
  )
}
