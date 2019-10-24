import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Curation, Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getAllPodcasts(state: AppState) {
  return state.entities.podcasts.podcasts
}

export function getPodcastById(state: AppState, podcastId: string) {
  return state.entities.podcasts.podcasts[podcastId]
}

export function makeGetPodcastsInCuration() {
  return createSelector<
    AppState,
    $Id<Curation>,
    MapById<Podcast>,
    $Id<Podcast>[],
    Podcast[]
  >(
    (state, _) => getAllPodcasts(state),
    (state, curationId) =>
      state.entities.podcasts.podcastsInCuration[curationId],
    (podcasts, podcastIds) => podcastIds.map((id) => podcasts[id]),
  )
}

/*
 * SUBSCRIPTIONS
 */

export function getIsCurrentUserSubscribedToPodcast(
  state: AppState,
  podcastId: string,
) {
  return state.entities.podcasts.currentUserSubscriptions.some(
    (id) => id === podcastId,
  )
}

export function makeGetCurrentUserSubscriptions() {
  return createSelector<AppState, $Id<Podcast>[], MapById<Podcast>, Podcast[]>(
    (state) => state.entities.podcasts.currentUserSubscriptions,
    (state) => state.entities.podcasts.podcasts,
    (subscriptions, podcasts) => {
      return subscriptions.map((id) => podcasts[id])
    },
  )
}

/*
 * TRENDING
 */

export function makeGetTrendingPodcasts() {
  return createSelector<AppState, $Id<Podcast>[], MapById<Podcast>, Podcast[]>(
    (state) => state.entities.podcasts.podcastsTrending,
    getAllPodcasts,
    (ids, podcasts) => {
      return ids.map((id) => podcasts[id])
    },
  )
}
