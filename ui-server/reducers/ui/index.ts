import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  CLOSE_SIGN_IN_MODAL,
  SEARCH_BAR_TEXT_CHANGE,
  SET_CURRENT_PATH_NAME,
  SET_SCREEN_WIDTH,
  SHOW_SIGN_IN_MODAL,
} from 'types/actions'
import { ScreenWidth } from 'types/app'
import player from './player'

const screenWidth: Reducer<ScreenWidth, AppActions> = (
  state = 'LG',
  action,
) => {
  switch (action.type) {
    case SET_SCREEN_WIDTH:
      return action.width
    default:
      return state
  }
}

const currentPathName: Reducer<string, AppActions> = (state = '', action) => {
  switch (action.type) {
    case SET_CURRENT_PATH_NAME:
      return action.pathName
    default:
      return state
  }
}

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
  screenWidth,
  currentPathName,
  searchText,
  showSignInModal,
  player,
})
