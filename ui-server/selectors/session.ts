import { AppState } from 'store'

export function getSignedInUserId(state: AppState) {
  return state.session.userId
}

export function getSubscriptions(state: AppState) {
  return state.session.subscriptions
}

export function getIsUserSignedIn(state: AppState) {
  return !!state.session.userId
}

export function getIsSubscribed(state: AppState, podcastId: string) {
  return state.session.subscriptions.some((id) => id === podcastId)
}
