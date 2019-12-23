import { AppState } from 'store'

export function getPlayingEpisodeId(state: AppState) {
  return state.ui.audioPlayer.playingEpisodeId
}

export function getDuration(state: AppState) {
  return state.ui.audioPlayer.duration
}

export function getState(state: AppState) {
  return state.ui.audioPlayer.state
}

export function getCurrentTime(state: AppState) {
  return state.ui.audioPlayer.currentTime
}

export function getVolume(state: AppState) {
  return state.ui.audioPlayer.volume
}

export function getPlaybackRate(state: AppState) {
  return state.ui.audioPlayer.playbackRate
}

export function getExpandOnMobile(state: AppState) {
  return state.ui.audioPlayer.expandOnMobile
}
