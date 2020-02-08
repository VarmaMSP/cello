import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Curation, Podcast } from 'types/models'
import { $Id, MapById } from 'types/utilities'

export function getCurationById(state: AppState, curationId: string) {
  return state.entities.curations.byId[curationId]
}

export function makeGetPodcastsInCuration() {
  return createSelector<
    AppState,
    $Id<Curation>,
    MapById<Podcast>,
    $Id<Podcast>[],
    Podcast[]
  >(
    (state) => state.entities.podcasts.byId,
    (state, id) => (state.entities.curations.byId[id] || {}).podcastIds || [],
    (all, ids) => ids.map((id) => all[id]),
  )
}
