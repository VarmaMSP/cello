import client from 'client'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getSubscriptionsFeed(publishedBefore: string) {
  return requestAction(
    () => client.getUserFeed(publishedBefore),
    (dispatch, _, { episodes, playbacks }) => {
      dispatch({
        type: T.RECEIVED_USER_FEED_PUBLISHED_BEFORE,
        publishedBefore,
        episodes,
      })
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
      })
    },
    { requestId: RequestId.getSubscriptionsFeed() },
  )
}
