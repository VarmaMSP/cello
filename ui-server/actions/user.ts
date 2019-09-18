import client from 'client'
import { Dispatch } from 'redux'
import {
  AppActions,
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  RECEIVED_SIGNED_IN_USER,
} from 'types/actions'

export const getSignedInUser = () => {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: GET_SIGNED_IN_USER_REQUEST })

    try {
      const { user } = await client.getSignedInUser()
      dispatch({ type: RECEIVED_SIGNED_IN_USER, user: user })
      dispatch({ type: GET_SIGNED_IN_USER_SUCCESS })
    } catch (err) {
      dispatch({ type: GET_SIGNED_IN_USER_FAILURE })
    }
  }
}
