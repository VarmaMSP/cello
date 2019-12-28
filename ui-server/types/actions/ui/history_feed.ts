
export const HISTORY_FEED_LOAD_PAGE = 'history_feed/load_page'
export const HISTORY_FEED_RECEIVED_ALL = 'history_feed/received_all'

interface LoadPageAction {
  type: typeof HISTORY_FEED_LOAD_PAGE
  page: number
  episodeIds: string[]
}

interface ReceivedAllAction {
  type: typeof HISTORY_FEED_RECEIVED_ALL
}

export type HistoryFeedActionTypes = LoadPageAction | ReceivedAllAction
