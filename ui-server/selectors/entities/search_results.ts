import { AppState } from 'store'

export function getPodcastSearchResultById(
  state: AppState,
  searchQuery: string,
  podcastId: string,
) {
  return (state.entities.searchResults.byPodcastId[searchQuery] || {})[
    podcastId
  ]
}
