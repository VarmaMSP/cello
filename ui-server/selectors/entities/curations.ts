import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Curation, Podcast } from 'types/models'
import { $Id, MapById } from 'types/utilities'

export function getCurationById(state: AppState, curationId: string) {
  return state.entities.curations.byId[curationId]
}

export const getAllCategories = createSelector<
  AppState,
  MapById<Curation>,
  $Id<Curation>[],
  Curation[]
>(
  (state) => state.entities.curations.byId,
  (state) => state.entities.curations.byType['category'],
  (all, ids) => ids.map((id) => all[id]),
)

export const getMainCategories = createSelector<
  AppState,
  Curation[],
  Curation[]
>(
  (state) => getAllCategories(state),
  (all) => all.filter((c) => c.parentId === undefined),
)

export function makeGetSubCategories() {
  return createSelector<
    AppState,
    $Id<Curation>,
    Curation[],
    $Id<Curation>,
    Curation[]
  >(
    (state) => getAllCategories(state),
    (_, id) => id,
    (all, parentId) => all.filter((c) => c.parentId === parentId),
  )
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
    (state, id) => (state.entities.curations.byId[id] || {}).members || [],
    (all, ids) => ids.map((id) => all[id]),
  )
}
