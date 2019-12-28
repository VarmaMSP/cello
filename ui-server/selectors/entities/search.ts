import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function makeGetSearchPodcastResults() {
  return createSelector<
    AppState,
    string,
    MapById<Podcast>,
    $Id<Podcast>[],
    Podcast[]
  >(
    (state, _) => state.entities.podcasts.byId,
    (state, query) => state.entities.search.podcastResults[query] || [],
    (podcasts, ids) => ids.map((id) => podcasts[id]),
  )
}
