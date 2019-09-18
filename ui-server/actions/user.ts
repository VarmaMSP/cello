import client from 'client'
import {
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  RECEIVED_SIGNED_IN_USER,
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
