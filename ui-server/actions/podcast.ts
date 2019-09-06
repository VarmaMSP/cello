import {
  AppActions,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  RECEIVED_PODCAST,
  RECEIVED_EPISODES,
  SEARCH_PODCASTS_REQUEST,
  SEARCH_PODCASTS_SUCCESS,
  RECEIVED_SEARCH_PODCASTS,
} from '../types/actions'
import { Dispatch } from 'redux'
import client from '../client'

export const getPodcast = (podcastId: string) => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({
      type: GET_PODCAST_REQUEST,
      podcastId,
    })
    const { podcast, episodes } = await client.getPodcastById(podcastId)
    dispatch({
      type: RECEIVED_PODCAST,
      podcast,
    })
    dispatch({
      type: RECEIVED_EPISODES,
      podcastId: podcast.id,
      episodes,
    })
    dispatch({
      type: GET_PODCAST_SUCCESS,
      podcastId,
    })
  }
}

export const searchPodcasts = (searchQuery: string) => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({
      type: SEARCH_PODCASTS_REQUEST,
    })
    const { podcasts } = await client.searchPodcasts(searchQuery)
    dispatch({
      type: RECEIVED_SEARCH_PODCASTS,
      podcasts,
    })
    dispatch({
      type: SEARCH_PODCASTS_SUCCESS,
    })
  }
}
