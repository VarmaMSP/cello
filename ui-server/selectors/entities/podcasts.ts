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
