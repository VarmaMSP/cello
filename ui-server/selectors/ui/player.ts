import { AppState } from 'store'
import { createSelector } from 'reselect'
import { MapById, $Id } from '../../types/utilities'
import { Episode, Podcast } from '../../types/app'
import { getAllEpisodes } from '../entities/episodes'
import { getAllPodcasts } from '../entities/podcasts'

export function getPlayingEpisodeId(state: AppState) {
  return state.ui.player.episode || ''
}

export function getAudioState(state: AppState) {
  return state.ui.player.audioState || 'LOADING'
}

export function getAudioDuration(state: AppState) {
  return state.ui.player.audioDuration || 0
}

export function getAudioCurrentTime(state: AppState) {
  return state.ui.player.audioCurrentTime || 0
}

export function getExpandOnMobile(state: AppState) {
  return state.ui.player.expandOnMobile || false
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

  return createSelector<AppState, MapById<Podcast>, $Id<Podcast>, Podcast>(
    getAllPodcasts,
    (state) => {
      const episode = getPlayingEpisode(state)
      return episode ? episode.podcastId : ''
    },
    (podcast, id) => podcast[id],
  )
}
