import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/models'
import { $Id, MapById } from 'types/utilities'

export function getEpisodeById(state: AppState, episodeId: string) {
  return state.entities.episodes.byId[episodeId]
}

export function getEpisodesByIds(state: AppState, episodeIds: string[]) {
  return episodeIds.map((id) => getEpisodeById(state, id))
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
