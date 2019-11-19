import { AppState } from 'store'
import * as RequestId from 'utils/request_id'

export function requestStatus(state: AppState, requestId: string) {
  return state.requests[requestId] || 'COMPLETE'
}

export function getPodcastEpisodesStatus(state: AppState, podcastId: string) {
  return requestStatus(state, RequestId.getPodcastEpisodes(podcastId))
}

export function getSubscriptionsFeedStatus(state: AppState) {
  return requestStatus(state, RequestId.getSubscriptionsFeed())
}
