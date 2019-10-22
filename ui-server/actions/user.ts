import client from 'client'
import * as T from 'types/actions'
import * as gtag from 'utils/gtag'
import { requestAction } from './utils'

export function getCurrentUser() {
  return requestAction(
    () => client.getSignedInUser(),
    (dispatch, { user, subscriptions }) => {
      if (!!user) {
        gtag.userId(user.id)
        dispatch({ type: T.RECEIVED_SIGNED_IN_USER, user, subscriptions })
      }
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

export function getUserFeed() {
  return requestAction(
    () => client.getUserFeed(),
    (dispatch, { episodes, playbacks }, getState) => {
      dispatch({
        type: T.RECEIVED_USER_FEED,
        userId: getState().entities.user.currentUserId,
        episodes,
      })
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
      })
    },
    { type: T.GET_USER_FEED_REQUEST },
    { type: T.GET_USER_FEED_SUCCESS },
    { type: T.GET_USER_FEED_FAILURE },
  )
}
