import { User } from 'types/app'

export const GET_SIGNED_IN_USER_REQUEST = 'GET_SIGNED_IN_USER_REQUEST'
export const GET_SIGNED_IN_USER_SUCCESS = 'GET_SIGNED_IN_USER_SUCCESS'
export const GET_SIGNED_IN_USER_FAILURE = 'GET_SIGNED_IN_USER_FAILURE'

export const SIGN_OUT_USER_REQUEST = 'SIGN_OUT_USER_REQUEST'
export const SIGN_OUT_USER_SUCCESS = 'SIGN_OUT_USER_SUCCESS'
export const SIGN_OUT_USER_FAILURE = 'SIGN_OUT_USER_FAILURE'

export const RECEIVED_SIGNED_IN_USER = 'RECEIVED_SIGNED_IN_USER'

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

export type UserActionTypes =
  | GetSignedInUserRequestAction
  | GetSignedInUserSuccessAction
  | GetSignedInUserFailureAction
  | ReceivedSignedInUserAction
  | SignOutUserRequestAction
  | SignOutUserSuccessAction
  | SignOutUserFailureAction
