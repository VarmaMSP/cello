import {
  AppActions,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  RECEIVED_PODCAST,
  RECEIVED_EPISODES,
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
      type: GET_PODCAST_SUCCESS,
      podcastId,
    })
    dispatch({
      type: RECEIVED_PODCAST,
      podcast,
    })
    dispatch({
      type: RECEIVED_EPISODES,
      podcastId: podcast.id,
      episodes,
    })
  }
}
