import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'models'
import { $Id } from 'types/utilities'

export function makeSelectSubscriptionsFeed() {
  return createSelector<
    AppState,
    { [page: string]: $Id<Episode>[] },
    string[],
    [$Id<Episode>[], boolean]
  >(
    (state) => state.ui.subscriptionsFeed.feed,
    (state) => state.ui.subscriptionsFeed.receivedAll,
    (obj, x) => [
      Object.keys(obj).reduce<$Id<Episode>[]>(
        (acc, key) => [...acc, ...obj[key]],
        [],
      ),
      x.includes('default'),
    ],
  )
}
