import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  CLOSE_SIGN_IN_MODAL,
  SEARCH_BAR_TEXT_CHANGE,
  SHOW_SIGN_IN_MODAL,
} from 'types/actions'
import player from './player'

const searchText: Reducer<string, AppActions> = (state = '', action) => {
  switch (action.type) {
    case SEARCH_BAR_TEXT_CHANGE:
      return action.text
    default:
      return state
  }
}

const showSignInModal: Reducer<boolean, AppActions> = (
  state = false,
  action,
) => {
  switch (action.type) {
    case SHOW_SIGN_IN_MODAL:
      return true
    case CLOSE_SIGN_IN_MODAL:
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
