import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Curation, CurationMember } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export const getAllCategories = createSelector<
  AppState,
  MapById<Curation>,
  $Id<Curation>[],
  Curation[]
>(
  (state) => state.entities.curations.byId,
  (state) => state.entities.curations.byType['CATEGORY'],
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
    MapById<CurationMember>,
    $Id<CurationMember>[],
    string[]
  >(
    (state) => state.entities.curationMember.byId,
    (state, id) => state.entities.curationMember.byCurationId[id] || [],
    (all, ids) => ids.map((id) => all[id].podcastId),
  )
}
