import client from 'client'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getPodcast(podcastId: string) {
  return requestAction(
    () => client.getPodcastById(podcastId),
    (dispatch, { podcast, episodes }) => {
      dispatch({ type: T.RECEIVED_PODCAST, podcast })
      dispatch({
        type: T.RECEIVED_PODCAST_EPISODES,
        podcastId: podcast.id,
        offset: 0,
        order: 'pub_date_desc',
        episodes,
      })
    },
    { type: T.GET_PODCAST_REQUEST },
    { type: T.GET_PODCAST_SUCCESS },
    { type: T.GET_PODCAST_FAILURE },
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
    (dispatch, { episodes, playbacks }) => {
      dispatch({
        type: T.RECEIVED_PODCAST_EPISODES,
        podcastId: podcastId,
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
          podcastId: podcastId,
          order,
        })
      }
    },
    { type: T.GET_PODCAST_EPISODES_REQUEST },
    { type: T.GET_PODCAST_EPISODES_SUCCESS },
    { type: T.GET_PODCAST_EPISODES_FAILURE },
  )
}

export function subscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.subscribeToPodcast(podcastId),
    (dispatch, _, getState) => {
      dispatch({
        type: T.SUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
    { type: T.SUBSCRIBE_TO_PODCAST_REQUEST },
    { type: T.SUBSCRIBE_TO_PODCAST_SUCCESS },
    { type: T.SUBSCRIBE_TO_PODCAST_FAILURE },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.unsubscribeToPodcast(podcastId),
    (dispatch, _, getState) => {
      dispatch({
        type: T.UNSUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
    { type: T.UNSUBSCRIBE_TO_PODCAST_REQUEST },
    { type: T.UNSUBSCRIBE_TO_PODCAST_SUCCESS },
    { type: T.UNSUBSCRIBE_TO_PODCAST_FAILURE },
  )
}

export function getTrendingPodcasts() {
  return requestAction(
    () => client.getTrendingPodcasts(),
    (dispatch, podcasts) => {
      dispatch({ type: T.RECEIVED_TRENDING_PODCASTS, podcasts })
    },
    { type: T.GET_TRENDING_PODCASTS_REQUEST },
    { type: T.GET_TRENDING_PODCASTS_SUCCESS },
    { type: T.GET_TRENDING_PODCASTS_FAILURE },
  )
}

export function searchPodcasts(query: string) {
  return requestAction(
    () => client.searchPodcasts(query),
    (dispatch, { podcasts }) => {
      dispatch({ type: T.RECEIVED_SEARCH_PODCASTS, query, podcasts })
    },
    { type: T.SEARCH_PODCASTS_REQUEST },
    { type: T.SEARCH_PODCASTS_SUCCESS },
    { type: T.SEARCH_PODCASTS_FAILURE },
  )
}
