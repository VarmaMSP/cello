import player from './player'
import { combineReducers, Reducer } from 'redux'
import { ScreenWidth } from '../../types/app'
import {
  AppActions,
  SET_SCREEN_WIDTH,
  SEARCH_BAR_TEXT_CHANGE,
} from '../../types/actions'

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
