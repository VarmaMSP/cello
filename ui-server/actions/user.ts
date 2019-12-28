import * as client from 'client/user'
import * as T from 'types/actions'
import * as gtag from 'utils/gtag'
import { requestAction } from './utils'

export function getCurrentUser() {
  return requestAction(
    () => client.init(),
    (dispatch, _, { user, subscriptions }) => {
      if (!!user) {
        gtag.userId(user.id)
        dispatch({ type: T.USER_ADD, users: [user] })
        dispatch({ type: T.PODCAST_ADD, podcasts: subscriptions })
        dispatch({ type: T.SESSION_INIT, userId: user.id })
        dispatch({
          type: T.SESSION_SUBSCRIBE_PODCASTS,
          podcastIds: subscriptions.map((x) => x.id),
        })
      }
    },
  )
}

export function signOutUser() {
  return requestAction(
    () => client.signOut(),
    (dispatch) => {
      dispatch({ type: T.SESSION_DELETE })
    },
  )
}
