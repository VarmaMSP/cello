import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Podcast, PodcastList } from 'types/app'
import { $Id, MapById, MapOneToMany } from 'types/utilities'

export function getCategoryById(state: AppState, categoryId: string) {
  return state.entities.podcastLists.categories[categoryId]
}

export function makeGetCategories() {
  return createSelector<
    AppState,
    MapById<PodcastList>,
    MapOneToMany<PodcastList, PodcastList>,
    PodcastList[]
  >(
    (state) => state.entities.podcastLists.categories,
    (state) => state.entities.podcastLists.subCategories,
    (all, subCategories) => Object.keys(subCategories).map((id) => all[id]),
  )
}

export function makeGetSubCategories() {
  return createSelector<
    AppState,
    $Id<PodcastList>,
    MapById<PodcastList>,
    $Id<PodcastList>[],
    PodcastList[]
  >(
    (state) => state.entities.podcastLists.categories,
    (state, id) => state.entities.podcastLists.subCategories[id],
    (all, ids) => ids.map((id) => all[id]),
  )
}

export function makeGetRecommendedPodcasts() {
  return createSelector<AppState, MapById<Podcast>, $Id<Podcast>[], Podcast[]>(
    (state) => state.entities.podcasts.podcasts,
    (state) => state.entities.podcastLists.recommended,
    (all, ids) => ids.map((id) => all[id]),
  )
}

export function makeGetPodcastsInList() {
  return createSelector<
    AppState,
    $Id<PodcastList>,
    MapById<Podcast>,
    $Id<Podcast>[],
    Podcast[]
  >(
    (state) => state.entities.podcasts.podcasts,
    (state, listId) => state.entities.podcastLists.podcastsInList[listId] || [],
    (all, ids) => ids.map((id) => all[id]),
  )
}
