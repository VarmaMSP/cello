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

export function makeGetPodcastEpisodes() {
  return createSelector<
    AppState,
    $Id<Podcast>,
    MapById<Episode>,
    {
      byPubDateDesc: { [offset: string]: string[] }
      byPubDateAsc: { [offset: string]: string[] }
      receivedAll: ('pub_date_desc' | 'pub_date_asc')[]
    },
    {
      episodes: Episode[]
      receivedAll: boolean
    }
  >(
    (state, _) => state.entities.episodes.episodes,
    (state, podcastId) =>
      state.entities.episodes.episodesInPodcast[podcastId] || {},
    (episodes, x) => {
      return {
        episodes: Object.values(x.byPubDateDesc || {}).reduce<Episode[]>(
          (acc, episodeIds) => [
            ...acc,
            ...episodeIds.map((id) => episodes[id]),
          ],
          [],
        ),
        receivedAll: (x.receivedAll || []).includes('pub_date_desc'),
      }
    },
  )
}
