import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getAllPodcasts(state: AppState) {
  return state.entities.podcasts.podcasts
}

export function getPodcastById(state: AppState, podcastId: string) {
  return state.entities.podcasts.podcasts[podcastId]
}

/*
 * SUBSCRIPTIONS
 */

export function getIsUserSubscribedToPodcast(
  state: AppState,
  podcastId: string,
) {
  return state.entities.podcasts.subscriptions.some((id) => id === podcastId)
}

export function makeGetSubscriptions() {
  return createSelector<AppState, MapById<Podcast>, $Id<Podcast>[], Podcast[]>(
    (state) => state.entities.podcasts.podcasts,
    (state) => state.entities.podcasts.subscriptions,
    (podcasts, subscriptions) => subscriptions.map((id) => podcasts[id]),
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
