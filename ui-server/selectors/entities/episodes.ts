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
