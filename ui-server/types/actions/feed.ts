import { Episode } from 'types/app'

export const RECEIVED_SUBSCRIPTION_FEED = 'RECEIVED_SUBSCRIPTION_FEED'
export const RECEIVED_ALL_SUBSCRIPTION_FEED = 'RECEIVED_ALL_SUBSCRIPTION_FEED'

export interface ReceivedSubscriptionFeedAction {
  type: typeof RECEIVED_SUBSCRIPTION_FEED
  offset: number
  episodes: Episode[]
}

export interface ReceivedAllSubscriptionFeedAction {
  type: typeof RECEIVED_ALL_SUBSCRIPTION_FEED
}

export type FeedActionTypes =
  | ReceivedSubscriptionFeedAction
  | ReceivedAllSubscriptionFeedAction
