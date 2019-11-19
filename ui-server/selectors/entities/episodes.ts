import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getAllEpisodes(state: AppState) {
  return state.entities.episodes.episodes
}

export function getEpisodeById(state: AppState, id: $Id<Episode>): Episode {
  return getAllEpisodes(state)[id]
}

export function makeGetEpisodesInPodcast() {
  return createSelector<
    AppState,
    $Id<Podcast>,
    $Id<Episode>[],
    MapById<Episode>,
    Episode[]
  >(
    (state, id) =>
      Object.values(
        (state.entities.episodes.episodesInPodcast[id] || {})[
          'byPubDateDesc'
        ] || {},
      ).reduce<string[]>((acc, ids) => [...acc, ...ids], []),
    (state, _) => getAllEpisodes(state),
    (ids, episodes) => {
      return ids.map((id) => episodes[id])
    },
  )
}

export function makeGetReceivedAllEpisodes() {
  return createSelector<AppState, $Id<Podcast>, ("pub_date_desc" | "pub_date_asc")[], boolean>(
    (state, id) => (state.entities.episodes.episodesInPodcast[id] || {}).receivedAll || [],
    (x) => x.includes("pub_date_desc"),
  )
}

/*
 * FEED
 */

export function makeGetCurrentUserFeed() {
  return createSelector<AppState, $Id<Episode>[], MapById<Episode>, Episode[]>(
    (state) =>
      Object.values(
        state.entities.episodes.currentUserFeed.publishedBefore,
      ).reduce<string[]>((acc, x) => [...acc, ...x], []),
    getAllEpisodes,
    (ids, episodes) => {
      return [...new Set(ids)]
        .map((id) => episodes[id])
        .sort((a, b) => +new Date(b.pubDate) - +new Date(a.pubDate))
    },
  )
}

/*
 * HISTORY
 */

export function makeGetCurrentUserHistory() {
  return createSelector<AppState, $Id<Episode>[], MapById<Episode>, Episode[]>(
    (state) => state.entities.episodes.currentUserHistory,
    getAllEpisodes,
    (ids, episodes) => {
      return ids.map((id) => episodes[id])
    },
  )
}

/*
 * PLAYBACK
 */

export function getCurrentUserPlayback(state: AppState, episodeId: string) {
  return state.entities.episodes.currentUserPlayback[episodeId]
}
