import client from 'client'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getPodcast(podcastId: string) {
  return requestAction(
    () => client.getPodcastById(podcastId),
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
    (dispatch, _, { episodes, playbacks }) => {
      dispatch({
        type: T.RECEIVED_PODCAST_EPISODES,
        podcastId,
        offset,
        order,
        episodes,
      })
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
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

export function subscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.subscribeToPodcast(podcastId),
    (dispatch, getState) => {
      dispatch({
        type: T.SUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.unsubscribeToPodcast(podcastId),
    (dispatch, getState) => {
      dispatch({
        type: T.UNSUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
  )
}

export function getTrendingPodcasts() {
  return requestAction(
    () => client.getTrendingPodcasts(),
    (dispatch, _, podcasts) => {
      dispatch({ type: T.RECEIVED_TRENDING_PODCASTS, podcasts })
    },
  )
}

export function searchPodcasts(query: string) {
  return requestAction(
    () => client.searchPodcasts(query),
    (dispatch, _, { podcasts }) => {
      dispatch({ type: T.RECEIVED_SEARCH_PODCASTS, query, podcasts })
    },
  )
}
