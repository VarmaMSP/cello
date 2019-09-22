import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { ViewportSize } from 'types/app'
import previousPage from './previous_page'

const viewportSize: Reducer<ViewportSize, T.AppActions> = (
  state = 'LG',
  action,
) => {
  switch (action.type) {
    case T.SET_VIEWPORT_SIZE:
      return action.size
    default:
      return state
  }
}

const currentUrlPath: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SET_CURRENT_URL_PATH:
      return action.urlPath
    default:
      return state
  }
}

export default combineReducers({
  previousPage,
  viewportSize,
  currentUrlPath,
})
