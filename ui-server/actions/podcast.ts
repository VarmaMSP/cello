import { AppActions, GET_PODCAST_REQUEST } from '../types/actions'
import { Dispatch } from 'redux'

export const getPodcast = (podcastId: string) => {
  return async (dispatch: Dispatch<AppActions>) => {
    return dispatch({
      type: GET_PODCAST_REQUEST,
      podcastId,
    })
  }
}
