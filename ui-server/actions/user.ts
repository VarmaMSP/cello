import client from 'client'
import * as T from 'types/actions'
import * as gtag from 'utils/gtag'
import { requestAction } from './utils'

export function getCurrentUser() {
  return requestAction(
    () => client.getSignedInUser(),
    (dispatch, _, { user, subscriptions }) => {
      if (!!user) {
        gtag.userId(user.id)
        dispatch({ type: T.RECEIVED_SIGNED_IN_USER, user, subscriptions })
      }
    },
  )
}

export function signOutUser() {
  return requestAction(
    () => client.signOutUser(),
    (dispatch) => {
      dispatch({ type: T.SIGN_OUT_USER })
    },
  )
}
