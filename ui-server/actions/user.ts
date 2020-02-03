import * as T from 'types/actions'
import { doFetch } from 'utils/fetch'
import * as gtag from 'utils/gtag'
import { requestAction } from './utils'

export function getCurrentUser() {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=load_session`,
      }),
    (dispatch, _, { users, podcasts }) => {
      if (users.length === 1) {
        gtag.userId(users[0].id)
        dispatch({ type: T.USER_ADD, users })
        dispatch({ type: T.PODCAST_ADD, podcasts })
        dispatch({ type: T.SESSION_INIT, userId: users[0].id })
        dispatch({
          type: T.SESSION_SUBSCRIBE_PODCASTS,
          podcastIds: podcasts.map((x) => x.id),
        })
      }
    },
  )
}

export function signOutUser() {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=end_session`,
      }),
    (dispatch) => {
      dispatch({ type: T.SESSION_DELETE })
    },
  )
}
