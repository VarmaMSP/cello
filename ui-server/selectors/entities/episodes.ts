import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getEpisodeById(state: AppState, id: string): Episode {
  return state.entities.episodes.byId[id]
}

export function makeGetEpisodesByPodcast() {
  return createSelector<
    AppState,
    $Id<Podcast>,
    MapById<Episode>,
    $Id<Episode>[],
    Episode[]
  >(
    (state) => state.entities.episodes.byId,
    (state, podcastId) => state.entities.episodes.byPodcastId[podcastId],
    (all, ids) => ids.map((id) => all[id]),
  )
}
