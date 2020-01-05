import { EpisodeSearchResult, PodcastSearchResult } from 'types/app'

export const SEARCH_RESULT_ADD_PODCAST = 'search_result/add_podcast'
export const SEARCH_RESULT_ADD_EPISODE = 'search_result/add_epiosde'

interface AddPodcastAction {
  type: typeof SEARCH_RESULT_ADD_PODCAST
  searchQuery: string
  podcastSearchResults: PodcastSearchResult[]
}

interface AddEpisodeAction {
  type: typeof SEARCH_RESULT_ADD_EPISODE
  searchQuery: string
  episodeSearchResults: EpisodeSearchResult[]
}

export type SearchResultActionTypes = AddPodcastAction | AddEpisodeAction
