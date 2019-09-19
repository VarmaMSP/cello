import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  RECEIVED_SIGNED_IN_USER,
  SIGN_OUT_USER_SUCCESS,
} from 'types/actions'
import { User } from 'types/app'

const currentUserId: Reducer<string, AppActions> = (state = '', action) => {
  switch (action.type) {
    case RECEIVED_SIGNED_IN_USER:
      return action.user.id
    case SIGN_OUT_USER_SUCCESS:
      return ''
    default:
      return state
  }
}

const users: Reducer<{ [userId: string]: User }, AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case RECEIVED_SIGNED_IN_USER:
      return { ...state, [action.user.id]: action.user }
    default:
      return state
  }
}

export default combineReducers({
  currentUserId,
  users,
})
