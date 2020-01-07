import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { SearchSortBy } from 'types/search'
import { $Id } from 'types/utilities'

export function getQuery(state: AppState) {
  return state.ui.resultsList.query
}

export function getResultType(state: AppState) {
  return state.ui.resultsList.resultType
}

export function getSortBy(state: AppState) {
  return state.ui.resultsList.sortBy
}

export function makeGetPodcasts() {
  return createSelector<
    AppState,
    string,
    SearchSortBy,
    {
      [query: string]: {
        [sortBy: string]: {
          [page: string]: $Id<Podcast>[]
        }
      }
    },
    string[],
    [$Id<Podcast>[], boolean]
  >(
    getQuery,
    getSortBy,
    (state) => state.ui.resultsList.podcasts,
    (state) => state.ui.resultsList.receivedAll,
    (query, sortBy, obj, receivedAll) => [
      !!obj[query] && !!obj[query][sortBy]
        ? Object.keys(obj[query][sortBy]).reduce<$Id<Podcast>[]>(
            (acc, key) => [...acc, ...obj[query][sortBy][key]],
            [],
          )
        : [],
      receivedAll.includes(`${query}:podcast:${sortBy}`),
    ],
  )
}

export function makeGetEpisodes() {
  return createSelector<
    AppState,
    string,
    SearchSortBy,
    {
      [query: string]: {
        [sortBy: string]: {
          [page: string]: $Id<Episode>[]
        }
      }
    },
    string[],
    [$Id<Episode>[], boolean]
  >(
    getQuery,
    getSortBy,
    (state) => state.ui.resultsList.episodes,
    (state) => state.ui.resultsList.receivedAll,
    (query, sortBy, obj, receivedAll) => [
      !!obj[query] && !!obj[query][sortBy]
        ? Object.keys(obj[query][sortBy]).reduce<$Id<Episode>[]>(
            (acc, key) => [...acc, ...obj[query][sortBy][key]],
            [],
          )
        : [],
      receivedAll.includes(`${query}:episode:${sortBy}`),
    ],
  )
}
