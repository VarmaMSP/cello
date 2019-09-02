import { AppState } from 'store'

export function getPodcastById(state: AppState, podcastId: string) {
  return state.entities.podcasts.podcasts[podcastId]
}
