import { createSelector } from 'reselect'
import { getAllEpisodes } from 'selectors/entities/episodes'
import { getAllPodcasts } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getPlayingEpisodeId(state: AppState) {
  return state.ui.player.episode
}

export function getAudioDuration(state: AppState) {
  return state.ui.player.duration
}

export function getAudioState(state: AppState) {
  return state.ui.player.audioState
}

export function getAudioCurrentTime(state: AppState) {
  return state.ui.player.currentTime
}

export function getAudioVolume(state: AppState) {
  return state.ui.player.volume
}

export function getAudioPlaybackSpeed(state: AppState) {
  return state.ui.player.playbackSpeed
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
