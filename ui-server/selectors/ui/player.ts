import { AppState } from 'store'
import { createSelector } from 'reselect'
import { MapById, $Id } from '../../types/utilities'
import { Episode, Podcast } from '../../types/app'
import { getAllEpisodes } from '../entities/episodes'
import { getAllPodcasts } from '../entities/podcasts'

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

export function makeGetPresentPodcast() {
  return createSelector<AppState, MapById<Podcast>, $Id<Podcast>, Podcast>(
    getAllPodcasts,
    (state) => state.ui.player.present.podcast || '',
    (podcasts, id) => podcasts[id],
  )
}

export function makeGetPresentEpisode() {
  return createSelector<AppState, MapById<Episode>, $Id<Episode>, Episode>(
    getAllEpisodes,
    (state) => state.ui.player.present.episode || '',
    (episodes, id) => episodes[id],
  )
}
