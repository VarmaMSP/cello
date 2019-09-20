import { combineReducers } from 'redux'
import {
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  SIGN_OUT_USER_FAILURE,
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
} from 'types/actions'
import { defaultRequestReducer } from './utils'

const getSignedInUser = defaultRequestReducer(
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  GET_SIGNED_IN_USER_FAILURE,
)

const signOutUser = defaultRequestReducer(
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
  SIGN_OUT_USER_FAILURE,
)

export default combineReducers({
  getSignedInUser,
  signOutUser,
})
