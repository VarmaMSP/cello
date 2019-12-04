import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'types/app'
import { MapById } from 'types/utilities'

export function makeGetSubscriptionsFeed() {
  return createSelector<
    AppState,
    MapById<Episode>,
    {
      byPubDateDesc: { [offset: string]: string[] }
      receivedAll: string[]
    },
    {
      episodes: Episode[]
      receivedAll: boolean
    }
  >(
    (state) => state.entities.episodes.episodes,
    (state) => state.entities.feed.subscriptions,
    (episodes, obj) => ({
      episodes: Object.values(obj.byPubDateDesc)
        .reduce<Episode[]>(
          (acc, episodeIds) => [
            ...acc,
            ...episodeIds.map((id) => episodes[id]),
          ],
          [],
        )
        .sort((a, b) => +new Date(b.pubDate) - +new Date(a.pubDate)),
      receivedAll: obj.receivedAll.includes('by_pub_date_desc'),
    }),
  )
}

export function makeGetHistoryFeed() {
  return createSelector<
    AppState,
    MapById<Episode>,
    {
      byPubDateDesc: { [offset: string]: string[] }
      receivedAll: string[]
    },
    {
      episodes: Episode[]
      receivedAll: boolean
    }
  >(
    (state) => state.entities.episodes.episodes,
    (state) => state.entities.feed.history,
    (episodes, obj) => ({
      episodes: Object.values(obj.byPubDateDesc)
        .reduce<Episode[]>(
          (acc, episodeIds) => [
            ...acc,
            ...episodeIds.map((id) => episodes[id]),
          ],
          [],
        )
        .sort((a, b) => +new Date(b.pubDate) - +new Date(a.pubDate)),
      receivedAll: obj.receivedAll.includes('by_pub_date_desc'),
    }),
  )
}
