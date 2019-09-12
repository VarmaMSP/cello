import client from 'client'
import { Dispatch } from 'redux'
import {
  AppActions,
  GET_PODCAST_CURATIONS_FAILURE,
  GET_PODCAST_FAILURE,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  RECEIVED_EPISODES,
  RECEIVED_PODCAST,
  RECEIVED_SEARCH_PODCASTS,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
} from 'types/actions'

export const getPodcast = (podcastId: string) => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: GET_PODCAST_REQUEST, podcastId })
    try {
      const { podcast, episodes } = await client.getPodcastById(podcastId)
      dispatch({ type: RECEIVED_PODCAST, podcast })
      dispatch({ type: RECEIVED_EPISODES, podcastId: podcast.id, episodes })
      dispatch({ type: GET_PODCAST_SUCCESS, podcastId })
    } catch (err) {
      dispatch({ type: GET_PODCAST_FAILURE, error: err.toString() })
    }
  }
}

export const searchPodcasts = (searchQuery: string) => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: SEARCH_PODCASTS_REQUEST })
    try {
      const { podcasts } = await client.searchPodcasts(searchQuery)
      dispatch({ type: RECEIVED_SEARCH_PODCASTS, podcasts })
      dispatch({ type: SEARCH_PODCASTS_SUCCESS })
    } catch (err) {
      dispatch({ type: GET_PODCAST_CURATIONS_FAILURE, error: err.toString() })
    }
  }
}
