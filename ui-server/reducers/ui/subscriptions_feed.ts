import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const feed: Reducer<{ [page: number]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.SUBSCRIPTIONS_FEED_LOAD_PAGE:
      return {
        ...state,
        [action.page]: action.episodeIds,
      }

    default:
      return state
  }
}

const receivedAll: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.SUBSCRIPTIONS_FEED_RECEIVED_ALL:
      return ['default']

    default:
      return state
  }
}

export default combineReducers({
  feed,
  receivedAll,
})
