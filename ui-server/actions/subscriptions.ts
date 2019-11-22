import client from 'client'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getSubscriptionsFeed_(offset: number, limit: number) {
  return requestAction(
    () => client.getSubscriptionsFeed(offset, limit),
    (dispatch, _, { episodes, playbacks }) => {
      dispatch({
        type: T.RECEIVED_SUBSCRIPTION_FEED,
        offset,
        episodes,
      })
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
      })

      if (episodes.length < limit) {
        dispatch({
          type: T.RECEIVED_ALL_SUBSCRIPTION_FEED,
        })
      }
    },
    { requestId: RequestId.getSubscriptionsFeed() },
  )
}
