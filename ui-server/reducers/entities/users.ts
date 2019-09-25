import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { User } from 'types/app'

const currentUserId: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return action.user.id
    case T.SIGN_OUT_USER_SUCCESS:
      return ''
    default:
      return state
  }
}

const users: Reducer<{ [userId: string]: User }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return { ...state, [action.user.id]: action.user }
    default:
      return state
  }
}

const feed: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.RECEIVED_USER_FEED:
      return action.episodes.map((e) => e.id)
    default:
      return state
  }
}

export default combineReducers({
  currentUserId,
  users,
  feed,
})
