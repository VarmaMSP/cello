import { combineReducers } from 'redux'
import {
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  SIGN_OUT_USER_FAILURE,
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
  SUBSCRIBE_TO_PODCAST_FAILURE,
  SUBSCRIBE_TO_PODCAST_REQUEST,
  SUBSCRIBE_TO_PODCAST_SUCCESS,
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

const subscribeToPodcast = defaultRequestReducer(
  SUBSCRIBE_TO_PODCAST_REQUEST,
  SUBSCRIBE_TO_PODCAST_SUCCESS,
  SUBSCRIBE_TO_PODCAST_FAILURE,
)

export default combineReducers({
  getSignedInUser,
  signOutUser,
  subscribeToPodcast,
})
