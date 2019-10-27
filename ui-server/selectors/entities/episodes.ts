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
    MapById<Episode>,
    $Id<Episode>[],
    Episode[]
  >(
    (state, _) => getAllEpisodes(state),
    (state, id) => state.entities.episodes.episodesInPodcast[id] || [],
    (episodes, ids) => {
      return ids.map((id) => episodes[id])
    },
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
