import { PodcastSearchResult } from 'types/models'

export const SEARCH_SUGGESTIONS_ADD_PODCAST = 'search_suggestions/add_podcast'
export const SEARCH_SUGGESTIONS_RESET = 'search_suggestions/reset'

export interface AddPodcastAction {
  type: typeof SEARCH_SUGGESTIONS_ADD_PODCAST
  podcasts: PodcastSearchResult[]
}

interface ResetPodcastAction {
  type: typeof SEARCH_SUGGESTIONS_RESET
}


export type SearchSuggestionsActionTypes = AddPodcastAction | ResetPodcastAction
