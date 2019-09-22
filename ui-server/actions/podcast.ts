import client from 'client'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getPodcast(podcastId: string) {
  return requestAction(
    () => client.getPodcastById(podcastId),
    (dispatch, { podcast, episodes }) => {
      dispatch({ type: T.RECEIVED_PODCAST, podcast })
      dispatch({ type: T.RECEIVED_EPISODES, podcastId: podcast.id, episodes })
    },
    { type: T.GET_PODCAST_REQUEST },
    { type: T.GET_PODCAST_SUCCESS },
    { type: T.GET_PODCAST_FAILURE },
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

export function searchPodcasts(searchQuery: string) {
  return requestAction(
    () => client.searchPodcasts(searchQuery),
    (dispatch, { podcasts }) => {
      dispatch({ type: T.RECEIVED_SEARCH_PODCASTS, podcasts })
    },
    { type: T.SEARCH_PODCASTS_REQUEST },
    { type: T.SEARCH_PODCASTS_SUCCESS },
    { type: T.SEARCH_PODCASTS_FAILURE },
  )
}
