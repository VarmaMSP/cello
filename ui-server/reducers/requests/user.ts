import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  GET_SIGNED_IN_USER_FAILURE,
  GET_SIGNED_IN_USER_REQUEST,
  GET_SIGNED_IN_USER_SUCCESS,
  SIGN_OUT_USER_FAILURE,
  SIGN_OUT_USER_REQUEST,
  SIGN_OUT_USER_SUCCESS,
} from 'types/actions'
import { initialRequestState, RequestState } from './utils'

const getSignedInUser: Reducer<RequestState, AppActions> = (
  state = initialRequestState(),
  action,
) => {
  switch (action.type) {
    case GET_SIGNED_IN_USER_REQUEST:
      return { status: 'STARTED', error: null }
    case GET_SIGNED_IN_USER_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case GET_SIGNED_IN_USER_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

const signOutUser: Reducer<RequestState, AppActions> = (
  state = initialRequestState(),
  action,
) => {
  switch (action.type) {
    case SIGN_OUT_USER_REQUEST:
      return { status: 'STARTED', error: null }
    case SIGN_OUT_USER_SUCCESS:
      return { status: 'SUCCESS', error: null }
    case SIGN_OUT_USER_FAILURE:
      return { status: 'FAILURE', error: null }
    default:
      return state
  }
}

export default combineReducers({
  getSignedInUser,
  signOutUser,
})
