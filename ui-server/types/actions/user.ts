import { Episode, Podcast, User } from 'types/app'

export const SIGN_OUT_USER = 'SIGN_OUT_USER'
export const SIGN_OUT_USER_FORCEFULLY = 'SIGN_OUT_USER_FORCEFULLY'
export const RECEIVED_SIGNED_IN_USER = 'RECEIVED_SIGNED_IN_USER'
export const SUBSCRIBED_TO_PODCAST = 'SUBSCRIBED_TO_PODCAST'
export const UNSUBSCRIBED_TO_PODCAST = 'UNSUBSCRIBED_TO_PODCAST'
export const RECEIVED_USER_FEED = 'RECEIVED_USER_FEED'
export const RECEIVED_USER_FEED_PUBLISHED_BEFORE =
  'RECEIVED_USER_FEED_PUBLISHED_BEFORE'

export interface ReceivedSignedInUserAction {
  type: typeof RECEIVED_SIGNED_IN_USER
  user: User
  subscriptions: Podcast[]
}
export interface SignOutUserAction {
  type: typeof SIGN_OUT_USER
}

export interface SignOutUserForcefullyAction {
  type: typeof SIGN_OUT_USER_FORCEFULLY
}

export interface SubscribedToPodcastAction {
  type: typeof SUBSCRIBED_TO_PODCAST
  userId: string
  podcastId: string
}

export interface UnsubscribedToPodcastAction {
  type: typeof UNSUBSCRIBED_TO_PODCAST
  userId: string
  podcastId: string
}

export interface ReceivedUserFeedAction {
  type: typeof RECEIVED_USER_FEED
  episodes: Episode[]
}

export interface ReceivedUserFeedPublishedBeforeAction {
  type: typeof RECEIVED_USER_FEED_PUBLISHED_BEFORE
  episodes: Episode[]
  publishedBefore: string
}

export type UserActionTypes =
  | ReceivedSignedInUserAction
  | SignOutUserAction
  | SignOutUserForcefullyAction
  | SubscribedToPodcastAction
  | UnsubscribedToPodcastAction
  | ReceivedUserFeedAction
  | ReceivedUserFeedPublishedBeforeAction
