import { Episode, Podcast, User } from 'types/app'

export const GET_SIGNED_IN_USER_REQUEST = 'GET_SIGNED_IN_USER_REQUEST'
export const GET_SIGNED_IN_USER_SUCCESS = 'GET_SIGNED_IN_USER_SUCCESS'
export const GET_SIGNED_IN_USER_FAILURE = 'GET_SIGNED_IN_USER_FAILURE'

export const SIGN_OUT_USER_REQUEST = 'SIGN_OUT_USER_REQUEST'
export const SIGN_OUT_USER_SUCCESS = 'SIGN_OUT_USER_SUCCESS'
export const SIGN_OUT_USER_FAILURE = 'SIGN_OUT_USER_FAILURE'

export const SUBSCRIBE_TO_PODCAST_REQUEST = 'SUBSCRIBE_TO_PODCAST_REQUEST'
export const SUBSCRIBE_TO_PODCAST_SUCCESS = 'SUBSCRIBE_TO_PODCAST_SUCCESS'
export const SUBSCRIBE_TO_PODCAST_FAILURE = 'SUBSCRIBE_TO_PODCAST_FAILURE'

export const UNSUBSCRIBE_TO_PODCAST_REQUEST = 'UNSUBSCRIBE_TO_PODCAST_REQUEST'
export const UNSUBSCRIBE_TO_PODCAST_SUCCESS = 'UNSUBSCRIBE_TO_PODCAST_SUCCESS'
export const UNSUBSCRIBE_TO_PODCAST_FAILURE = 'UNSUBSCRIBE_TO_PODCAST_FAILURE'

export const GET_USER_FEED_REQUEST = 'GET_USER_FEED_REQUEST'
export const GET_USER_FEED_SUCCESS = 'GET_USER_FEED_SUCCESS'
export const GET_USER_FEED_FAILURE = 'GET_USER_FEED_FAILURE'

export const RECEIVED_SIGNED_IN_USER = 'RECEIVED_SIGNED_IN_USER'
export const SUBSCRIBED_TO_PODCAST = 'SUBSCRIBED_TO_PODCAST'
export const UNSUBSCRIBED_TO_PODCAST = 'UNSUBSCRIBED_TO_PODCAST'
export const RECEIVED_USER_FEED = 'RECEIVED_USER_FEED'

export interface GetSignedInUserRequestAction {
  type: typeof GET_SIGNED_IN_USER_REQUEST
}

export interface GetSignedInUserSuccessAction {
  type: typeof GET_SIGNED_IN_USER_SUCCESS
}

export interface GetSignedInUserFailureAction {
  type: typeof GET_SIGNED_IN_USER_FAILURE
}

export interface ReceivedSignedInUserAction {
  type: typeof RECEIVED_SIGNED_IN_USER
  user: User
  subscriptions: Podcast[]
}

export interface SignOutUserRequestAction {
  type: typeof SIGN_OUT_USER_REQUEST
}

export interface SignOutUserSuccessAction {
  type: typeof SIGN_OUT_USER_SUCCESS
}

export interface SignOutUserFailureAction {
  type: typeof SIGN_OUT_USER_FAILURE
}

export interface SubscribeToPodcastRequestAction {
  type: typeof SUBSCRIBE_TO_PODCAST_REQUEST
}

export interface SubscribeToPodcastSuccessAction {
  type: typeof SUBSCRIBE_TO_PODCAST_SUCCESS
}

export interface SubscribeToPodcastFailureAction {
  type: typeof SUBSCRIBE_TO_PODCAST_FAILURE
}

export interface SubscribedToPodcastAction {
  type: typeof SUBSCRIBED_TO_PODCAST
  userId: string
  podcastId: string
}

export interface UnsubscribeToPodcastRequestAction {
  type: typeof UNSUBSCRIBE_TO_PODCAST_REQUEST
}

export interface UnsubscribeToPodcastSuccessAction {
  type: typeof UNSUBSCRIBE_TO_PODCAST_SUCCESS
}

export interface UnsubscribeToPodcastFailureAction {
  type: typeof UNSUBSCRIBE_TO_PODCAST_FAILURE
}

export interface UnsubscribedToPodcastAction {
  type: typeof UNSUBSCRIBED_TO_PODCAST
  userId: string
  podcastId: string
}

export interface GetUserFeedRequestAction {
  type: typeof GET_USER_FEED_REQUEST
}

export interface GetUserFeedSuccessAction {
  type: typeof GET_USER_FEED_SUCCESS
}

export interface GetUserFeedFailureAction {
  type: typeof GET_USER_FEED_FAILURE
}

export interface ReceivedUserFeedAction {
  type: typeof RECEIVED_USER_FEED
  userId: string
  episodes: Episode[]
}

export type UserActionTypes =
  | GetSignedInUserRequestAction
  | GetSignedInUserSuccessAction
  | GetSignedInUserFailureAction
  | ReceivedSignedInUserAction
  | SignOutUserRequestAction
  | SignOutUserSuccessAction
  | SignOutUserFailureAction
  | SubscribeToPodcastRequestAction
  | SubscribeToPodcastSuccessAction
  | SubscribeToPodcastFailureAction
  | SubscribedToPodcastAction
  | UnsubscribeToPodcastRequestAction
  | UnsubscribeToPodcastSuccessAction
  | UnsubscribeToPodcastFailureAction
  | UnsubscribedToPodcastAction
  | GetUserFeedRequestAction
  | GetUserFeedSuccessAction
  | GetUserFeedFailureAction
  | ReceivedUserFeedAction
