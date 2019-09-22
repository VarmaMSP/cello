import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import player from './player'

const searchText: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SEARCH_BAR_TEXT_CHANGE:
      return action.text
    default:
      return state
  }
}

const showSignInModal: Reducer<boolean, T.AppActions> = (
  state = false,
  action,
) => {
  switch (action.type) {
    case T.SHOW_SIGN_IN_MODAL:
      return true
    case T.CLOSE_SIGN_IN_MODAL:
      return false
    default:
      return state
  }
}

export default combineReducers({
  searchText,
  showSignInModal,
  player,
})
