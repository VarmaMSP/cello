import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'types/models'
import { $Id } from 'types/utilities'

export function makeGetEpisodeIds() {
  return createSelector<AppState, { [page: number]: string[] }, $Id<Episode>[]>(
    (state) => state.ui.historyFeed.feed,
    (obj) =>
      Object.keys(obj).reduce<string[]>((acc, id) => [...acc, ...obj[+id]], []),
  )
}

export function getReceivedAll(state: AppState) {
  return state.ui.historyFeed.receivedAll.some((x) => x === 'default')
}
