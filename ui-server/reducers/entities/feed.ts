import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const subscriptions: Reducer<
  { byPubDateDesc: { [offset: string]: string[] }; receivedAll: string[] },
  T.AppActions
> = (state = { byPubDateDesc: {}, receivedAll: [] }, action) => {
  switch (action.type) {
    case T.RECEIVED_SUBSCRIPTION_FEED:
      return {
        ...state,
        byPubDateDesc: {
          ...state.byPubDateDesc,
          [action.offset.toString()]: action.episodes.map((e) => e.id),
        },
      }
    case T.RECEIVED_ALL_SUBSCRIPTION_FEED:
      return {
        ...state,
        receivedAll: [...state.receivedAll, 'pub_date_asc'],
      }
    default:
      return state
  }
}

export default combineReducers({
  subscriptions,
})
