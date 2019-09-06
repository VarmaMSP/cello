import { AppState } from 'store'
import { createSelector } from 'reselect'
import { MapById, $Id } from '../../types/utilities'
import { Episode, Podcast } from '../../types/app'
import { getAllEpisodes } from '../entities/episodes'
import { getAllPodcasts } from '../entities/podcasts'

export function getPlayingEpisodeId(state: AppState) {
  return state.ui.player.episode
}

export function getAudioState(state: AppState) {
  return state.ui.player.audioState
}

export function getExpandOnMobile(state: AppState) {
  return state.ui.player.expandOnMobile
}

export function makeGetPlayingEpisode() {
  return createSelector<AppState, MapById<Episode>, $Id<Episode>, Episode>(
    getAllEpisodes,
    getPlayingEpisodeId,
    (episodes, id) => episodes[id],
  )
}

export function makeGetPlayingPodcast() {
  const getPlayingEpisode = makeGetPlayingEpisode()

  return createSelector<AppState, MapById<Podcast>, Episode, Podcast>(
    getAllPodcasts,
    getPlayingEpisode,
    (podcasts, episode) => {
      const podcastId = !!episode ? episode.podcastId : ''
      return podcasts[podcastId]
    },
  )
}
