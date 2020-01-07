import { createSelector } from 'reselect'
import { getQuery } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import {
  Episode,
  EpisodeSearchResult,
  Podcast,
  PodcastSearchResult,
} from 'types/app'
import { $Id } from 'types/utilities'

export function makeGetPodcastSearchResultById() {
  return createSelector<
    AppState,
    $Id<Podcast>,
    { [query: string]: { [id: string]: PodcastSearchResult } },
    string,
    $Id<Podcast>,
    PodcastSearchResult
  >(
    (state) => state.entities.searchResults.byPodcastId,
    (state) => getQuery(state),
    (_, podcastId) => podcastId,
    (obj, query, podcastId) => (obj[query] || {})[podcastId],
  )
}

export function makeGetEpisodeSearchResultById() {
  return createSelector<
    AppState,
    $Id<Episode>,
    { [query: string]: { [id: string]: EpisodeSearchResult } },
    string,
    $Id<Episode>,
    EpisodeSearchResult
  >(
    (state) => state.entities.searchResults.byEpisodeId,
    (state) => getQuery(state),
    (_, episodeId) => episodeId,
    (obj, query, episodeId) => (obj[query] || {})[episodeId],
  )
}
