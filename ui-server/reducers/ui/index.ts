import player from './player'
import { combineReducers, Reducer } from 'redux'
import { ScreenWidth } from '../../types/app'
import { AppActions, SET_SCREEN_WIDTH } from '../../types/actions'

const screenWidth: Reducer<ScreenWidth | undefined, AppActions> = (
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

export default combineReducers({
  screenWidth,
  player,
})
