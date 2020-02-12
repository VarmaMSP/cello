import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { User } from 'models'

const byId: Reducer<{ [userId: string]: User }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.USER_ADD:
      return {
        ...state,
        ...action.users.reduce<{ [userId: string]: User }>(
          (acc, u) => ({ ...acc, [u.id]: { ...(state[u.id] || {}), ...u } }),
          {},
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
 