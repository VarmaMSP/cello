import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Podcast } from 'types/app'
import { $Id } from 'types/utilities'

export function makeSelectPodcasts() {
  return createSelector<
    AppState,
    { searchQuery: string },
    { [page: string]: $Id<Podcast>[] },
    boolean,
    [$Id<Podcast>[], boolean]
  >(
    (state, { searchQuery }) =>
      state.ui.resultsList.podcasts[searchQuery] || {},

    (state, { searchQuery }) =>
      (state.ui.resultsList.receivedAll[searchQuery] || []).includes('default'),

    (obj, receivedAll) => [
      Object.keys(obj).reduce<$Id<Podcast>[]>(
        (acc, key) => [...acc, ...obj[key]],
        [],
      ),
      receivedAll,
    ],
  )
}
