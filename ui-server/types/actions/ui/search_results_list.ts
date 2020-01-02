export const SEARCH_RESULTS_LIST_LOAD_PAGE = 'search_results_list/load_page'
export const SEARCH_RESULTS_LIST_RECEIVED_ALL = 'search_results_list/received_all'

interface LoadPageAction {
  type: typeof SEARCH_RESULTS_LIST_LOAD_PAGE
  searchQuery: string
  podcastIds: string[]
  page: number
}

interface ReceivedAllAction {
  type: typeof SEARCH_RESULTS_LIST_RECEIVED_ALL
  searchQuery: string
}

export type SearchResultsListActionTypes = LoadPageAction | ReceivedAllAction
