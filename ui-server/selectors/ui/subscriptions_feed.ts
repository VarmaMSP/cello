import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'types/app'
import { $Id } from 'types/utilities'

export function makeGetEpisodeIds() {
  return createSelector<AppState, { [page: number]: string[] }, $Id<Episode>[]>(
    (state) => state.ui.subscriptionsFeed.feed,
    (obj) =>
      Object.keys(obj).reduce<string[]>((acc, id) => [...acc, ...obj[+id]], []),
  )
}

export function getReceivedAll(state: AppState) {
  return state.ui.subscriptionsFeed.receivedAll.some((x) => x === 'default')
}
