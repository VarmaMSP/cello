import { AppState } from 'store'
import * as RequestId from 'utils/request_id'

export function requestStatus(state: AppState, requestId: string) {
  return state.requests[requestId] || 'NOT_STARTED'
}

export function getPodcastEpisodesStatus(state: AppState, podcastId: string) {
  return requestStatus(state, RequestId.getPodcastEpisodes(podcastId))
}

export function getSubscriptionsPageStatus(state: AppState) {
  return requestStatus(state, RequestId.getSubscriptionsPageData())
}

export function getSubscriptionsFeedStatus(state: AppState) {
  return requestStatus(state, RequestId.getSubscriptionsFeed())
}

export function getHistoryPageStatus(state: AppState) {
  return requestStatus(state, RequestId.getHistoryPageData())
}

export function getHistoryFeedStatus(state: AppState) {
  return requestStatus(state, RequestId.getHistoryFeed())
}

export function getCreatePlaylistStatus(state: AppState) {
  return requestStatus(state, RequestId.createPlaylist())
}

export function getPlaylistPageStatus(state: AppState) {
  return requestStatus(state, RequestId.getPlaylistPageData())
}

export function getPlaylistFeedStatus(state: AppState) {
  return requestStatus(state, RequestId.getPlaylistFeed())
}

export function getResultsStatus(state: AppState) {
  return requestStatus(state, RequestId.getResults())
}