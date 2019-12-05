import * as client from 'client/subscription'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getSubscriptionsFeed(offset: number, limit: number) {
  return requestAction(
    () => client.getSubscriptionsFeed(offset, limit),
    (dispatch, _, { episodes }) => {
      dispatch({ type: T.RECEIVED_SUBSCRIPTION_FEED, offset, episodes })

      if (episodes.length < limit) {
        dispatch({ type: T.RECEIVED_ALL_SUBSCRIPTION_FEED })
      }
    },
    { requestId: RequestId.getSubscriptionsFeed() },
  )
}

export function subscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.subscribeToPodcast(podcastId),
    (dispatch) => {
      dispatch({ type: T.SUBSCRIBED_TO_PODCAST, podcastId })
    },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.unsubscribeToPodcast(podcastId),
    (dispatch) => {
      dispatch({ type: T.UNSUBSCRIBED_TO_PODCAST, podcastId })
    },
  )
}