import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const getSignedInUser = defaultRequestReducer(
  T.GET_SIGNED_IN_USER_REQUEST,
  T.GET_SIGNED_IN_USER_SUCCESS,
  T.GET_SIGNED_IN_USER_FAILURE,
)

const signOutUser = defaultRequestReducer(
  T.SIGN_OUT_USER_REQUEST,
  T.SIGN_OUT_USER_SUCCESS,
  T.SIGN_OUT_USER_FAILURE,
)

const subscribeToPodcast = defaultRequestReducer(
  T.SUBSCRIBE_TO_PODCAST_REQUEST,
  T.SUBSCRIBE_TO_PODCAST_SUCCESS,
  T.SUBSCRIBE_TO_PODCAST_FAILURE,
)

const getUserFeed = defaultRequestReducer(
  T.GET_USER_FEED_REQUEST,
  T.GET_USER_FEED_SUCCESS,
  T.GET_USER_FEED_FAILURE,
)

export default combineReducers({
  getSignedInUser,
  getUserFeed,
  signOutUser,
  subscribeToPodcast,
})
