import { PodcastSearchResult } from 'types/models'

export const SEARCH_SUGGESTIONS_ADD_PODCAST = `search_suggestions/add_podcast`

interface AddPodcastAction {
  type: typeof SEARCH_SUGGESTIONS_ADD_PODCAST
  podcasts: PodcastSearchResult[]
}

export type SearchSuggestionsActionTypes = AddPodcastAction
