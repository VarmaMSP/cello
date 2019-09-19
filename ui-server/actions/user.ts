import client from 'client'
import {
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  RECEIVED_SIGNED_IN_USER,
  SIGN_OUT_USER_FAILURE,
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
} from 'types/actions'
import { requestAction } from './utils'

export function getSignedInUser() {
  return requestAction(
    () => client.getSignedInUser(),
    (dispatch, { user }) => dispatch({ type: RECEIVED_SIGNED_IN_USER, user }),
    { type: GET_SIGNED_IN_USER_REQUEST },
    { type: GET_SIGNED_IN_USER_SUCCESS },
    { type: GET_SIGNED_IN_USER_FAILURE },
  )
}

export function signOutUser() {
  return requestAction(
    () => client.signOutUser(),
    () => {},
    { type: SIGN_OUT_USER_REQUEST },
    { type: SIGN_OUT_USER_SUCCESS },
    { type: SIGN_OUT_USER_FAILURE },
  )
}
