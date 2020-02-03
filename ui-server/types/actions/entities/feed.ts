import { Episode } from 'types/models'

export const RECEIVED_HISTORY_FEED = 'RECEIVED_HISTORY_FEED'
export const RECEIVED_ALL_HISTORY_FEED = 'RECEIVED_ALL_HISTORY_FEED'
export const RECEIVED_SUBSCRIPTION_FEED = 'RECEIVED_SUBSCRIPTION_FEED'
export const RECEIVED_ALL_SUBSCRIPTION_FEED = 'RECEIVED_ALL_SUBSCRIPTION_FEED'

export interface ReceivedHistoryFeedAction {
  type: typeof RECEIVED_HISTORY_FEED
  offset: number
  episodes: Episode[]
}

export interface ReceivedAllHistoryFeedAction {
  type: typeof RECEIVED_ALL_HISTORY_FEED
}

export interface ReceivedSubscriptionFeedAction {
  type: typeof RECEIVED_SUBSCRIPTION_FEED
  offset: number
  episodes: Episode[]
} 

export interface ReceivedAllSubscriptionFeedAction {
  type: typeof RECEIVED_ALL_SUBSCRIPTION_FEED
}

export type FeedActionTypes =
  | ReceivedHistoryFeedAction
  | ReceivedAllHistoryFeedAction
  | ReceivedSubscriptionFeedAction
  | ReceivedAllSubscriptionFeedAction
