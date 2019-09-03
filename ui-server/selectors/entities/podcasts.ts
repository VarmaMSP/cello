import { AppState } from 'store'

export function getAllPodcasts(state: AppState) {
  return state.entities.podcasts.podcasts
}

export function getPodcastById(state: AppState, podcastId: string) {
  return state.entities.podcasts.podcasts[podcastId]
}
