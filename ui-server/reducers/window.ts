import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { ViewportSize } from 'types/app'

const viewportSize: Reducer<ViewportSize, T.AppActions> = (
  state = 'LG',
  action,
) => {
  switch (action.type) {
    case T.WINDOW_VIEWPORT_SIZE:
      return action.size

    default:
      return state
  }
}

export default combineReducers({
  viewportSize,
})
