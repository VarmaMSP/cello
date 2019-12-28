export const SUBSCRIPTIONS_FEED_LOAD_PAGE = 'subscriptions_feed/load_page'
export const SUBSCRIPTIONS_FEED_RECEIVED_ALL = 'subscriptions_feed/received_all'

interface LoadPageAction {
  type: typeof SUBSCRIPTIONS_FEED_LOAD_PAGE
  page: number
  episodeIds: string[]
}

interface ReceivedAllAction {
  type: typeof SUBSCRIPTIONS_FEED_RECEIVED_ALL
}

export type SubscriptionsFeedActionTypes = LoadPageAction | ReceivedAllAction
