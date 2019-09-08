import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  SEARCH_BAR_TEXT_CHANGE,
  SET_SCREEN_WIDTH,
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

const searchText: Reducer<string, AppActions> = (state = '', action) => {
  switch (action.type) {
    case SEARCH_BAR_TEXT_CHANGE:
      return action.text
    default:
      return state
  }
}

export default combineReducers({
  screenWidth,
  searchText,
  player,
})
