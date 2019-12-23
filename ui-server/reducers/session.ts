import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { addKeysToArr, delKeysFromArr } from 'utils/immutable'

const userId: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SESSION_INIT:
      return action.userId

    case T.SESSION_DELETE:
      return ''

    default:
      return state
  }
}

const subscriptions: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.SESSION_SUBSCRIBE_PODCASTS:
      return addKeysToArr(action.podcastIds, state)

    case T.SESSION_UNSUBSCRIBE_PODCASTS:
      return delKeysFromArr(action.podcastIds, state)

    default:
      return state
  }
}

export default combineReducers({
  userId,
  subscriptions,
})
