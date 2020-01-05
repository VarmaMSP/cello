import { SearchResultType, SearchSortBy } from 'types/search'

export const SEARCH_RESULTS_LIST_LOAD_PODCAST_PAGE =
  'search_results_list/load_podcast_page'
export const SEARCH_RESULTS_LIST_LOAD_EPISODE_PAGE =
  'search_results_list/load_episode_page'
export const SEARCH_RESULTS_LIST_RECEIVED_ALL =
  'search_results_list/received_all'

interface LoadPodcastPageAction {
  type: typeof SEARCH_RESULTS_LIST_LOAD_PODCAST_PAGE
  searchQuery: string
  sortBy: SearchSortBy
  page: number
  podcastIds: string[]
}

interface LoadEpisodePageAction {
  type: typeof SEARCH_RESULTS_LIST_LOAD_EPISODE_PAGE
  searchQuery: string
  sortBy: SearchSortBy
  page: number
  episodeIds: string[]
}

interface ReceivedAllAction {
  type: typeof SEARCH_RESULTS_LIST_RECEIVED_ALL
  searchQuery: string
  resultType: SearchResultType
  sortBy: SearchSortBy
}

export type SearchResultsListActionTypes =
  | LoadPodcastPageAction
  | LoadEpisodePageAction
  | ReceivedAllAction
