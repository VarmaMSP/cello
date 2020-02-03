import { doFetch_ } from 'client/fetch'
import * as client from 'client/podcast'
import * as T from 'types/actions'
import { PodcastEpisodeListOrder } from 'types/ui'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPodcastPageData(podcastUrlParam: string) {
  return requestAction(
    () =>
      doFetch_({
        method: 'GET',
        urlPath: `/podcasts/${podcastUrlParam}`,
      }),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.EPISODE_ADD, episodes })

      dispatch({
        type: T.PODCAST_EPISODES_LIST_LOAD_PAGE,
        podcastId: podcasts[0].id,
        episodeIds: episodes.map((x) => x.id),
        order: 'pub_date_desc',
        page: 0,
      })
    },
  )
}

export function getPodcastEpisodes(
  podcastId: string,
  limit: number,
  offset: number,
  order: PodcastEpisodeListOrder,
) {
  return requestAction(
    () => client.getPodcastEpisodes(podcastId, limit, offset, order),
    (dispatch, _, { episodes }) => {
      dispatch({
        type: T.EPISODE_ADD,
        episodes,
      })

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
