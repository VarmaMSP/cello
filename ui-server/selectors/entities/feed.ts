import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'types/app'
import { MapById } from 'types/utilities'

export function makeGetSubscriptionsFeed() {
  return createSelector<
    AppState,
    { byPubDateDesc: { [offset: string]: string[] }; receivedAll: string[] },
    MapById<Episode>,
    Episode[]
  >(
    (state) => state.entities.feed.subscriptions,
    (state) => state.entities.episodes.episodes,
    (obj, episodes) =>
      Object.values(obj.byPubDateDesc)
        .reduce<string[]>((acc, ids) => [...acc, ...ids], [])
        .map((id) => episodes[id])
        .sort((a, b) => +new Date(b.pubDate) - +new Date(a.pubDate)),
  )
}

export function makeGetHistoryFeed() {
  return createSelector<
    AppState,
    { byPubDateDesc: { [offset: string]: string[] }; receivedAll: string[] },
    MapById<Episode>,
    Episode[]
  >(
    (state) => state.entities.feed.history,
    (state) => state.entities.episodes.episodes,
    (obj, episodes) =>
      Object.values(obj.byPubDateDesc)
        .reduce<string[]>((acc, ids) => [...acc, ...ids], [])
        .map((id) => episodes[id])
        .sort((a, b) => +new Date(b.pubDate) - +new Date(a.pubDate)),
  )
}
