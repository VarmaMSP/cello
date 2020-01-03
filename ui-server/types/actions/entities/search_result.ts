import { PodcastSearchResult } from 'types/app'

export const SEARCH_RESULT_ADD_PODCAST = 'search_result/add_podcast'
export const SEARCH_RESULT_ADD_EPIOSDE = 'search_result/add_epiosde'

interface AddPodcastAction {
  type: typeof SEARCH_RESULT_ADD_PODCAST
  searchQuery: string
  podcastSearchResults: PodcastSearchResult[]
}

export type SearchResultActionTypes = AddPodcastAction
