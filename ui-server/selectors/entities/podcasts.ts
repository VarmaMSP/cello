import { AppState } from 'store'


export function getPodcastById(state: AppState, podcastId: string) {
  return state.entities.podcasts.byId[podcastId]
}

export function getPodcastsByIds(state: AppState, podcastIds: string[]) {
  return podcastIds.map(id => getPodcastById(state, id))
}