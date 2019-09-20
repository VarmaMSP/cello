import client from 'client'
import {
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  RECEIVED_SIGNED_IN_USER,
  SIGN_OUT_USER_FAILURE,
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
  SUBSCRIBED_TO_PODCAST,
  SUBSCRIBE_TO_PODCAST_FAILURE,
  SUBSCRIBE_TO_PODCAST_REQUEST,
  SUBSCRIBE_TO_PODCAST_SUCCESS,
  UNSUBSCRIBED_TO_PODCAST,
  UNSUBSCRIBE_TO_PODCAST_FAILURE,
  UNSUBSCRIBE_TO_PODCAST_REQUEST,
  UNSUBSCRIBE_TO_PODCAST_SUCCESS,
} from 'types/actions'
import { requestAction } from './utils'

export function getSignedInUser() {
  return requestAction(
    () => client.getSignedInUser(),
    (dispatch, { user }) => dispatch({ type: RECEIVED_SIGNED_IN_USER, user }),
    { type: GET_SIGNED_IN_USER_REQUEST },
    { type: GET_SIGNED_IN_USER_SUCCESS },
    { type: GET_SIGNED_IN_USER_FAILURE },
  )
}

export function signOutUser() {
  return requestAction(
    () => client.signOutUser(),
    () => {},
    { type: SIGN_OUT_USER_REQUEST },
    { type: SIGN_OUT_USER_SUCCESS },
    { type: SIGN_OUT_USER_FAILURE },
  )
}

export function subscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.subscribeToPodcast(podcastId),
    (dispatch, _, getState) => {
      dispatch({
        type: SUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
    { type: SUBSCRIBE_TO_PODCAST_REQUEST },
    { type: SUBSCRIBE_TO_PODCAST_SUCCESS },
    { type: SUBSCRIBE_TO_PODCAST_FAILURE },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.unsubscribeToPodcast(podcastId),
    (dispatch, _, getState) => {
      dispatch({
        type: UNSUBSCRIBED_TO_PODCAST,
        userId: getState().entities.user.currentUserId,
        podcastId,
      })
    },
    { type: UNSUBSCRIBE_TO_PODCAST_REQUEST },
    { type: UNSUBSCRIBE_TO_PODCAST_SUCCESS },
    { type: UNSUBSCRIBE_TO_PODCAST_FAILURE },
  )
}
