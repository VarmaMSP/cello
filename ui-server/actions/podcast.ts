import client from 'client'
import {
  GET_PODCAST_FAILURE,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_TRENDING_PODCASTS_FAILURE,
  GET_TRENDING_PODCASTS_REQUEST,
  GET_TRENDING_PODCASTS_SUCCESS,
  RECEIVED_EPISODES,
  RECEIVED_PODCAST,
  RECEIVED_SEARCH_PODCASTS,
  RECEIVED_TRENDING_PODCASTS,
  SEARCH_PODCASTS_FAILURE,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
} from 'types/actions'
import { requestAction } from './utils'

export function getPodcast(podcastId: string) {
  return requestAction(
    () => client.getPodcastById(podcastId),
    (dispatch, { podcast, episodes }) => {
      dispatch({ type: RECEIVED_PODCAST, podcast })
      dispatch({ type: RECEIVED_EPISODES, podcastId: podcast.id, episodes })
    },
    { type: GET_PODCAST_REQUEST },
    { type: GET_PODCAST_SUCCESS },
    { type: GET_PODCAST_FAILURE },
  )
}

export function getTrendingPodcasts() {
  return requestAction(
    () => client.getTrendingPodcasts(),
    (dispatch, podcasts) => {
      dispatch({ type: RECEIVED_TRENDING_PODCASTS, podcasts })
    },
    { type: GET_TRENDING_PODCASTS_REQUEST },
    { type: GET_TRENDING_PODCASTS_SUCCESS },
    { type: GET_TRENDING_PODCASTS_FAILURE },
  )
}

export function searchPodcasts(searchQuery: string) {
  return requestAction(
    () => client.searchPodcasts(searchQuery),
    (dispatch, { podcasts }) => {
      dispatch({ type: RECEIVED_SEARCH_PODCASTS, podcasts })
    },
    { type: SEARCH_PODCASTS_REQUEST },
    { type: SEARCH_PODCASTS_SUCCESS },
    { type: SEARCH_PODCASTS_FAILURE },
  )
}
