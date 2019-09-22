import client from 'client'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getSignedInUser() {
  return requestAction(
    () => client.getSignedInUser(),
    (dispatch, { user, subscriptions }) => {
      dispatch({ type: T.RECEIVED_SIGNED_IN_USER, user, subscriptions })
    },
    { type: T.GET_SIGNED_IN_USER_REQUEST },
    { type: T.GET_SIGNED_IN_USER_SUCCESS },
    { type: T.GET_SIGNED_IN_USER_FAILURE },
  )
}

export function signOutUser() {
  return requestAction(
    () => client.signOutUser(),
    () => {},
    { type: T.SIGN_OUT_USER_REQUEST },
    { type: T.SIGN_OUT_USER_SUCCESS },
    { type: T.SIGN_OUT_USER_FAILURE },
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
