import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { $Id } from 'types/utilities'

export function makeSelectPodcasts() {
  return createSelector<
    AppState,
    { searchQuery: string; sortBy: string },
    { [page: string]: $Id<Podcast>[] },
    boolean,
    [$Id<Podcast>[], boolean]
  >(
    (state, { searchQuery, sortBy }) =>
      (state.ui.resultsList.podcasts[searchQuery] || {})[sortBy] || {},

    (state, { searchQuery, sortBy }) =>
      state.ui.resultsList.receivedAll.includes(
        `${searchQuery}:podcast:${sortBy}`,
      ),

    (obj, receivedAll) => [
      Object.keys(obj).reduce<$Id<Podcast>[]>(
        (acc, key) => [...acc, ...obj[key]],
        [],
      ),
      receivedAll,
    ],
  )
}

export function makeSelectEpisodes() {
  return createSelector<
    AppState,
    { searchQuery: string; sortBy: string },
    { [page: string]: $Id<Podcast>[] },
    boolean,
    [$Id<Episode>[], boolean]
  >(
    (state, { searchQuery, sortBy }) =>
      (state.ui.resultsList.episodes[searchQuery] || {})[sortBy] || {},

    (state, { searchQuery, sortBy }) =>
      state.ui.resultsList.receivedAll.includes(
        `${searchQuery}:episode:${sortBy}`,
      ),

    (obj, receivedAll) => [
      Object.keys(obj).reduce<$Id<Podcast>[]>(
        (acc, key) => [...acc, ...obj[key]],
        [],
      ),
      receivedAll,
    ],
  )
}
