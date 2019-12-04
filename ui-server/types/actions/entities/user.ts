import { Podcast, User } from 'types/app'

export const SIGN_OUT_USER = 'SIGN_OUT_USER'
export const SIGN_OUT_USER_FORCEFULLY = 'SIGN_OUT_USER_FORCEFULLY'
export const RECEIVED_SIGNED_IN_USER = 'RECEIVED_SIGNED_IN_USER'

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

export type UserActionTypes =
  | ReceivedSignedInUserAction
  | SignOutUserAction
  | SignOutUserForcefullyAction
