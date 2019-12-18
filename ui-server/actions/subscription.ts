import * as client from 'client/subscription'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getSubscriptionsPageData() {
  return requestAction(
    () => client.getSubscriptionsPageData(),
    (dispatch, _, { episodes }) => {
      dispatch({
        type: T.RECEIVED_SUBSCRIPTION_FEED,
        offset: 0,
        episodes,
      })
      if (episodes.length < 15) {
        dispatch({ type: T.RECEIVED_ALL_SUBSCRIPTION_FEED })
      }
    },
    { requestId: RequestId.getSubscriptionsPageData() },
  )
}

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
    () => client.subscribePodcast(podcastId),
    (dispatch) => {
      dispatch({ type: T.SUBSCRIBED_TO_PODCAST, podcastId })
    },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () => client.unsubscribePodcast(podcastId),
    (dispatch) => {
      dispatch({ type: T.UNSUBSCRIBED_TO_PODCAST, podcastId })
    },
  )
}
